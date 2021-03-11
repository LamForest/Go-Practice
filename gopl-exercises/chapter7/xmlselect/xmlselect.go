// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const filename = "test.xml"

func main() {
	// filename
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Reading XML failed... exit...")
	}

	dec := xml.NewDecoder(strings.NewReader(string(data)))
	var stack []string // stack of element names
	for {
		//tok类型为Token, 是一个万能接口：
		//type Token interface{}
		//
		//A Token is an interface holding one of the token types:
		// StartElement(起始标签),
		// EndElement（终止标签),
		// CharData（有意义的文本）,
		// Comment（评论<!-- -->）,
		// ProcInst(长这样<?target inst?>，不太清楚作用),
		// or Directive.（导语？形式如<!text>）
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		case xml.Directive:
			fmt.Printf("Directive  = %s\n", tok)
		default:
			fmt.Printf("Default get [%s, %T]\n", tok, tok)
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
