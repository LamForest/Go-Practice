//20个长期存活的goroutine
//经测试,并令牌方式慢很多,取得1w个URL需要多几十倍的时间
package main

import (
	"gopl/chapter8/links"
	"log"
)

const MAX_CNT = 10000

func crawl(ch chan<- []string, unseenLinks <-chan string) {

	for URL := range unseenLinks {

		linksInPage, err := links.Extract(URL)
		if err != nil {
			log.Printf("Error when Extract(%s), err is : %v\n", URL, err)
			continue
		}
		// ch <- linksInPage // 等待main取走links,循环等待
		go func() { ch <- linksInPage }()

	}

}

var unseenLinks chan string
var ch chan []string

func init() {
	unseenLinks = make(chan string)
	ch = make(chan []string)

	for i := 0; i < 20; i++ {
		go crawl(ch, unseenLinks)
	}
}

func main() {
	seen := make(map[string]bool)

	go func() {
		// ch <- []string{"https://blog.wolfogre.com/posts/slice-queue-vs-list-queue/"}
		ch <- []string{"https://www.runoob.com/go/go-slice.html"}
		// ch <- []string{"https://stackoverflow.com/questions/29164375/correct-way-to-initialize-empty-slice"}
	}()

	cnt := 1
	for x := range ch {
		log.Printf("x len is %d, current cnt is %d\n", len(x), cnt)
		for _, l := range x {
			if seen[l] == false {
				// go crawl(l, ch, unseenLinks)
				// log.Printf("aa\n") //等待worker取走URL
				seen[l] = true
				unseenLinks <- l
				cnt += 1
				// log.Printf("link %s\n", l)
			}
		}
	}

}
