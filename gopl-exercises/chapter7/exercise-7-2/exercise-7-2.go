package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type byteCounter struct {
	w       io.Writer
	counter int64
}

func (bc *byteCounter) Write(p []byte) (int, error) {
	bc.counter += int64(len(p))
	return bc.w.Write(p)
}

//CountingWriter : 代码风格检查真傻逼（错误
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := byteCounter{w, 0}
	return &bc, &bc.counter

}

func main() {
	//虽然说CountingWriter返回的io.Writer是指针类型，但是go会自如得在指针和值之间切换
	//对于一个接口来说，我们只会调用它的方法，所以指针和值等价，没有任何影响
	writer, pcount := CountingWriter(ioutil.Discard)
	fmt.Println(*pcount)
	fmt.Fprintf(writer, "111")
	fmt.Println(*pcount)
}
