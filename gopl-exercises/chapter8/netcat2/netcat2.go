// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	dur, _ := time.ParseDuration("5s")
	conn, err := net.DialTimeout("tcp", "localhost:9001", dur)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	// mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	//不用for，一直有数据，直到tcp另一端关闭
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
