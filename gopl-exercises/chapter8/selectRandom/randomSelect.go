//本代码用于验证select在遇到多个满足条件的分支时，是否会随机选择
package main

import "fmt"

func main() {

	const total = 10000
	for i := 0; i < 10; i++ {
		ch1, ch2 := selectTest(total)
		fmt.Printf("Channel 1: %d, Channel 2 %d, total %d\n", ch1, ch2, total)
	}

}

func selectTest(channel_capacity int) (int, int) {
	odd_ch := make(chan struct{}, channel_capacity)
	even_ch := make(chan struct{}, channel_capacity)
	for i := 0; i < channel_capacity; i++ {
		odd_ch <- struct{}{}
		even_ch <- struct{}{}
	}
	var odd_cnt, even_cnt int //初始化为0
	for i := 0; i < channel_capacity; i++ {
		//不是顺序执行的，而是先统计有多少满足的case，再随机选取
		select {
		case <-odd_ch:
			odd_cnt++
		case <-even_ch:
			even_cnt++
		}
	}
	return odd_cnt, even_cnt
}
