//练习exercise 8.3，对应的echo服务器为../echo_server/echo_server.go
//主goroutin从tcp中读取，向stdout输出
//次goroutin stdin->tcp
//值得注意的是io.Copy读取到空但非EOF不返回
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("ERROR : net.Listen return error %v\n", err)
	}
	//
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatalf("connot convert conn %T to net.TCPconn exit...", conn)
	}
	defer func() {
		err := tcpconn.Close()
		if err != nil {
			log.Println("defer error ", err)
		}
	}()
	go func(tcpconn *net.TCPConn) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// log.Printf("input : %s\n", scanner.Text())
			if _, err := conn.Write([]byte(scanner.Text() + "\n")); err != nil {
				log.Printf("error while copy from stdin to conn : %v,break", err)
				break
			}

		}
		tcpconn.CloseWrite()

	}(tcpconn)
	//注意，这里不用循环从conn中读取
	//io.Copy内部有个循环，会从reader中一直循环读取
	//tcp中没有数据：读取0个字符，但err为空，继续循环。
	//tcp中有数据：读取0个字符，err为空，继续循环
	//tcp断开：读取0个字符，err为EOF，停止循环，因为EOF不算是错误而是标志，所以返回的err也为空
	//其他异常：返回错误
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Printf("error while copy from conn to stdout : %v", err)
	}
	tcpconn.CloseRead()
}
