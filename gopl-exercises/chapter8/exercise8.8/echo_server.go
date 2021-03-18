//echo服务器，并行
package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {

	// allClose := make(chan struct{})
	log.SetFlags(log.Lshortfile)
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal("ERROR : net.Listen return error %v\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept return ERROR : %v\n", err)
			continue
		}
		log.Printf("Get connection %s\n", conn.RemoteAddr())

		go func(conn net.Conn) {
			defer conn.Close()

			scanner := bufio.NewScanner(conn)

			alive := make(chan struct{})
			//10s后关闭连接
			const countdown = time.Second * 10
			go func() {
				ticker := time.NewTicker(countdown)
				defer ticker.Stop()
				for {

					select {
					case <-alive:
						log.Println("Coming new msg, reset ticker...")
						ticker.Reset(countdown)
					case <-ticker.C:
						conn.Close() //关闭连接应该会使Scan得到EOF
						log.Printf("Last msg over %s secs, close...\n", countdown)
						return
					}
				}

			}()

			for scanner.Scan() {
				log.Printf("Receive From Client %s\n", scanner.Text())
				alive <- struct{}{}
				// time.Sleep(time.Second * 10)
				conn.Write([]byte("ECHO : " + scanner.Text() + "\n"))
			}
			log.Println("exited from : for scanner.Scan()")

			// conn.Read()
			// if _, err := io.Copy(conn, conn); err != nil {
			// 	log.Printf("ERROR in go func %v\n", err)
			// }

		}(conn)
	}

}
