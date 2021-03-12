//时钟墙，从多个时钟服务器获得时间，练习8.1
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	name    string
	addr    string
	message string
}

func main() {
	myargv := []string{"NY=localhost:8010", "BJ=localhost:8010"}
	servers := parse(myargv)
	for _, ser := range servers {
		conn, err := net.Dial("tcp", ser.addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully connected to server %s\n", ser.addr)
		defer conn.Close()
		go func(ser *server) {
			sc := bufio.NewScanner(conn)

			for sc.Scan() {
				ser.message = sc.Text()
				fmt.Printf("[%s] : %s\n", ser.name, ser.message)
			}
		}(ser)
	}

	// 	fmt.Printf("\n")
	// 	//这没有冲突？
	// 	for _, server := range servers {
	// 		fmt.Printf("%s: %s\n", server.name, server.message)
	// 	}
	// 	fmt.Print("--------")

	time.Sleep(time.Second * 10)

}

func parse(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{s[0], s[1], ""})
	}
	return servers
}
