package main

import (
	"fmt"
)

type bytecounter float64

func (counter *bytecounter) Write(p []byte) (int, error) {
	*counter += bytecounter(len(p))
	return len(p), nil
}

func init() {
	fmt.Println("init()...")
}

func main() {
	var counter bytecounter = 0.0
	fmt.Println(counter)
	fmt.Fprintf(&counter, "12345%d", 999)
	fmt.Println(counter)
}
