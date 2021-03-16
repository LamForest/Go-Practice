//通过信道,控制最大并行度
package main

import (
	"gopl/chapter8/links"
	"log"
)

const MAX_CNT = 10000

type token struct{}

func crawl(URL string, ch chan<- []string, tokens chan token) {
	tokens <- token{}
	//  <- tokens ,这是我的获取令牌的写法,需要事先向里面填充20个令牌,比较麻烦
	linksInPage, err := links.Extract(URL)
	//信号量操作应该离它所约束的操作越近越好
	<-tokens
	if err != nil {
		log.Printf("Error when Extract(%s), err is : %v\n", URL, err)
		return
	}

	//先还令牌还是传送links
	//感觉应该先传送
	ch <- linksInPage

}

func main() {
	ch := make(chan []string)
	tokens := make(chan token, 20)
	// 不需要往里面填20个令牌
	// 空的位置当作一个令牌即可,往tokens里面写,相当于获得令牌
	// 从tokens中读,相当于释放令牌
	// for i := 0; i < 20; i++ {
	// 	tokens <- token{}
	// }
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
				go crawl(l, ch, tokens)
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
