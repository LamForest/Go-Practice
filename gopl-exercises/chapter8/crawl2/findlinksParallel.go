package main

import (
	"gopl/chapter8/links"
	"log"
)

const MAX_CNT = 10000

func crawl(URL string, ch chan<- []string) {
	linksInPage, err := links.Extract(URL)
	if err != nil {
		log.Printf("Error when Extract(%s), err is : %v\n", URL, err)
		return
	}

	ch <- linksInPage

}

func main() {
	ch := make(chan []string)
	seen := make(map[string]bool)

	go func() {
		// ch <- []string{"https://blog.wolfogre.com/posts/slice-queue-vs-list-queue/"}
		ch <- []string{"https://www.runoob.com/go/go-slice.html"}
		// ch <- []string{"https://stackoverflow.com/questions/29164375/correct-way-to-initialize-empty-slice"}
	}()

	cnt := 0
	for x := range ch {
		log.Printf("x len is %d, current cnt is %d\n", len(x), cnt)
		for _, l := range x {
			if seen[l] == false {
				go crawl(l, ch)
				seen[l] = true
			}
			cnt += 1
		}
		if cnt >= MAX_CNT {
			log.Printf("Total %d links\n", cnt)
			break
		}
	}

}
