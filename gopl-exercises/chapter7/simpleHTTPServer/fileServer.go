//一个简单的http 文件服务器，类似python simpleHTTPServer
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("p", 8080, "listen to port..")

func main() {
	flag.Parse()
	ip := fmt.Sprintf("0.0.0.0:%d", *port)
	http.Handle("/", http.FileServer(http.Dir("../")))
	fmt.Printf("File Server ListenAndServe on port %d", *port)
	log.Fatal(http.ListenAndServe(ip, nil))
}
