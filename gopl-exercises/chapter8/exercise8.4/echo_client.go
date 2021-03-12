//练习exercise 8.4，对应的echo服务器为../echo_server/echo_server.go
//并行向echo服务器发请求，使用了sync.WaitGroup
package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

const ip = "localhost:9001"

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func worker(msgs <-chan string, echomsgs chan<- string) { //Receive only chan
	defer wc.Done()
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatalf("ERROR : net.Listen return error %v\n", err)
	}

	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatalf("connot convert conn %T to net.TCPconn exit...", conn)
	}
	//关闭conn
	defer func() {
		err := tcpconn.Close()
		if err != nil {
			log.Println("defer error ", err)
		}
	}()
	msg := <-msgs
	tcpconn.Write([]byte(msg + "\n"))
	// var echomsg []byte
	var echomsg = make([]byte, 100)
	// buf := bufio.NewReader(tcpconn)
	//这个read是阻塞的，当前没有数据时，一直等待
	//直到有数据返回，就返回数据，哪怕数据长度小于数组长度
	_, err = tcpconn.Read(echomsg)
	// log.Println(n, err)
	echomsgs <- string(echomsg)

}

func leader(echomsgs chan string) {
	wc.Wait()
	close(echomsgs)
}

var wc sync.WaitGroup

func main() {
	const num = 5
	msgs := make(chan string, num)
	echomsgs := make(chan string, num)

	//go leader(echomsgs)
	//

	for i := 0; i < num; i++ {
		msgs <- fmt.Sprintf("[%2d]-th echo", i)
		wc.Add(1)
		go worker(msgs, echomsgs)

	}

	go leader(echomsgs)

	//这里对信道进行循环
	//所以仅当leader close信道之后，循环才会退出
	//换言之，主goroutine保证在所有goroutine结束之后才结束
	for echomsg := range echomsgs {
		log.Print(echomsg)
	}

	// if _, err := io.Copy(os.Stdout, conn); err != nil {
	// 	log.Printf("error while copy from conn to stdout : %v", err)
	// }
	// tcpconn.CloseRead()
}
