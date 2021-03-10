package main

import (
	"fmt"
	"io"
)

type myStruct struct {
	i int
}

func (s *myStruct) Write(p []byte) (int, error) {
	return 0, nil
}

func main() {
	var pa *myStruct //空接口
	fmt.Printf("pa == nil : %v\n", pa == nil)
	fmt.Printf("pa = (%v, %T)\n", pa, pa) //true，确实是空接口

	var w io.Writer = pa                    //包含空指针的非空接口
	fmt.Printf("w == nil : %v\n", w == nil) //false，w不为nil，而是包含空指针的非空接口
	fmt.Printf("w = (%v, %T)\n", w, w)

	fmt.Printf("w == pa : %v\n", w == pa) // true，因为二者的动态类型/动态值完全相等
}
