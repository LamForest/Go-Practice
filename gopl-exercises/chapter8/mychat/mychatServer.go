//聊天服务器
//我觉得很有意思的一点是，GO的同步方式为：
//多个线程将需要同步的数据通过信道传输给一个单例对象（在本文件中是broadcaster)，由该单例对象进行同步
//
//GO对于多线程同步非常友好，不需要上锁操作，因为单例对象不会同时运行
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type msg struct {
	clientname string
	data       string
}

type newClient struct {
	clientname string
	toClient   chan string
}

var (
	mainMsgChan     = make(chan msg)       //所有handleConn向broadcaster传消息的信道
	newClientChan   = make(chan newClient) //handleConn向broadcaster注册自己，并给出向客户端传回信息的信道toClient
	closeClientChan = make(chan string)    //hanleConn处conn断开了连接，向broadcaster发送结束标志
)

func broadcaster() {
	//client channel map
	//用于与clientWriter通信，broadcast不直接与client通信，而是发给clientWriter
	clientChans := map[string]chan string{}

	//广播函数
	broadcast := func(s string) {
		for _, ch := range clientChans {
			ch <- s
		}
	}

	for {
		select {
		//case : when new Client connecting
		case newClient := <-newClientChan:
			broadcastMsg := fmt.Sprintf("New client [%s] Connected...\n", newClient.clientname)
			broadcast(broadcastMsg)
			//add new client channel to channel List
			clientChans[newClient.clientname] = newClient.toClient

			//case: when a client send a msg, which then will be broadcasted to all clients
		case inMsg := <-mainMsgChan:
			broadcastMsg := time.Now().Format("At 00:00:00 by ") + inMsg.clientname + " : " + inMsg.data + "\n"
			broadcast(broadcastMsg)

			//case: when a client disconnect, handleconn will notify broadcaster to close its channel through closeClientChan.
			// handleconn does not  close its channel. Close a channel twice will cause panic
		case clientName := <-closeClientChan:
			close(clientChans[clientName])
			delete(clientChans, clientName)
			broadcastMsg := fmt.Sprintf("Client [%s] Disonnected...\n", clientName)
			broadcast(broadcastMsg)
		}
	}
}

//keep receiving msg from broadcaster and send to corresponding client
func clientWrite(conn net.Conn, ch <-chan string) {
	for out := range ch {
		conn.Write([]byte(out))
	}
	log.Printf("clientWrite exited\n")
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	myChan := make(chan string)
	// defer close(myChan) //close channel should be in sender. so broadcaster will close channel， close a channel twice will cause panic

	clientName := conn.RemoteAddr().String()
	newClientChan <- newClient{clientName, myChan}
	// log.Println(clientName)
	go clientWrite(conn, myChan)

	scanner := bufio.NewScanner(conn)
	// log.Println(clientName + "123")

	for scanner.Scan() {
		log.Printf("Receive From Client [%s] : %s\n", clientName, scanner.Text())
		mainMsgChan <- msg{clientName, scanner.Text()}
		// time.Sleep(time.Second * 10)
		// conn.Write([]byte("ECHO : " + scanner.Text() + "\n"))
	}
	closeClientChan <- clientName
	log.Println("exited from : for scanner.Scan()")

}

func main() {

	// log.SetFlags(log.Lshortfile)
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal("ERROR : net.Listen return error %v\n", err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept return ERROR : %v\n", err)
			continue
		}
		log.Printf("Get connection %s\n", conn.RemoteAddr())
		go handleConn(conn)
	}

}
