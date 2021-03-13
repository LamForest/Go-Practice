package main

import (
	"fmt"
	"log"
	"os"

	"gopl/chapter8/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
//这应该是广度优先遍历吧
func breathFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist //当前循环需要遍历的
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breath-first.
	// starting from the command-line-arguments.
	breathFirst(crawl, os.Args[1:])
}
