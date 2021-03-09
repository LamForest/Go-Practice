package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://1.15.149.213:8080/ping"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v/n", err)
	} else {
		body := resp.Body
		b, err := ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "read body: %v/n", err)
			os.Exit(1)
		}

		fmt.Printf("Fetch %s ok, Result is : %s, %T\n", url, b, b)

	}

}
