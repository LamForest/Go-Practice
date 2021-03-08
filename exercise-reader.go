/*
https://tour.go-zh.org/methods/22
这个比较简单，没什么好说的
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func (reader MyReader) Read(b []byte)(int,error){
	for i := range b{
		b[i] = 'A'	
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
