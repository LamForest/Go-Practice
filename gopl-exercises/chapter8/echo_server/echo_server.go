//echo服务器，并行
package main

import (
	"bufio"
	"log"
	"net"
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
			for scanner.Scan() {
				log.Printf("Receive From Client %s\n", scanner.Text())
				// time.Sleep(time.Second * 10)
				conn.Write([]byte("ECHO : " + scanner.Text() + "\n"))
			}

			// conn.Read()
			// if _, err := io.Copy(conn, conn); err != nil {
			// 	log.Printf("ERROR in go func %v\n", err)
			// }

		}(conn)
	}

}
