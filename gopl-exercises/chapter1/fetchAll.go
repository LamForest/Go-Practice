// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func isEmptyLine(line string) bool {
	strings.Trim(line, "\n \t")
	if len(line) == 0 {
		return true
	}
	return false
}

func main() {
	//Read urls from files
	filename := "urls.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error : %v when reading file: %s", err, filename)
	}
	lines := strings.Split(string(data), "\n")
	var validURLCnt int = 0

	//Start Fetching
	start := time.Now()
	ch := make(chan string)
	for _, url := range lines {
		if strings.HasPrefix(url, "#") {
			continue
		}
		validURLCnt++
		go fetch(url, ch) // start a goroutine
	}
	for i := 0; i < validURLCnt; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) { //单向string信道的意思？
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Err %v while http.Get(url = [%s]) ", err, url) // sent to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.4fs %7d %s", secs, nbytes, url)
}
