package main

import (
	"fmt"
	"gopl/chapter8/links"
	"log"
)

//map,slice默认初始化为nil，必须要显式初始化
type linkWithDepth struct {
	url   string
	depth int
}

var seen = map[string]bool{}
var targetLinks = []linkWithDepth{}

func BFS(URL string, MAX_DEPTH int) {
	// var curDepth int
	targetLinks = append(targetLinks, linkWithDepth{URL, 0})
	seen[URL] = true
	for len(targetLinks) != 0 {
		curlink := targetLinks[0]
		curURL, curDepth := curlink.url, curlink.depth
		targetLinks = targetLinks[1:]

		fmt.Printf("D%d : %s\n", curDepth, curURL)
		if curDepth >= MAX_DEPTH {
			continue
		}
		// targetLinks =
		linksInURL, err := links.Extract(curURL)
		log.Printf("URL %s has %d links\n", curURL, len(linksInURL))
		if err != nil {
			log.Printf("Extract(%s) gets error : %v", URL, err)
			continue
		}
		for _, link := range linksInURL {
			if !seen[link] {
				seen[link] = true
				targetLinks = append(
					targetLinks, linkWithDepth{link, curDepth + 1})
			}
		}
	}

}

func main() {
	const MAX_DEPTH = 2

	// const URL = "https://www.baidu.com/"
	const URL = "https://blog.wolfogre.com/posts/slice-queue-vs-list-queue/"
	BFS(URL, MAX_DEPTH)
}
