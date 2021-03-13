// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		//不同类型的节点的定义见https://developer.mozilla.org/zh-CN/docs/Web/API/Node/nodeType
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				//将外部变量引入闭包
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//对以n为根节点的DOM树进行前序/后序遍历
//为什么没有后序遍历呢？因为多叉树没有中序遍历这个概念
//forEachNode只是对树进行遍历，对每个结点的处理逻辑由pre post定义
//这种解耦很牛逼
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
