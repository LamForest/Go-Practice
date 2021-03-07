//练习：斐波纳契闭包
//让我们用函数做些好玩的事情。

//实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。


package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var lastEle, lastLastEle, cnt = 1, 1, 1
	return func () int{
		var ret int
		if(cnt == 1 || cnt == 2){
			ret = 1
			cnt += 1
		}else{ 
			ret = lastEle + lastLastEle
			lastLastEle = lastEle
			lastEle = ret
			cnt += 1
		}
		return ret
	}
}
//输出:
/*
0-th :   1
 1-th :   1
 2-th :   2
 3-th :   3
 4-th :   5
 5-th :   8
 6-th :  13
 7-th :  21
 8-th :  34
 */
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d-th : %3d\n", i, f())
	}
}
