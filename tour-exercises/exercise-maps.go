/*
https://tour.go-zh.org/moretypes/23

练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。

输出：
PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}


*/

package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

// 验证程序把符号也算进了单词
func isAlpha(ch rune) bool {
	
	/*
	if((ch <= 'z' && ch >= 'a') || (ch <= 'Z' && ch >= 'A')){
		return true
	}*/
	if(ch != ' ' && ch != '\n'){
		return true	
	}
	
	return false
}

//不使用库函数
/*
func WordCount(s string) map[string]int {
	//fmt.Println(s)
	var startOfWord int
	var inWord bool
	
	ret := make(map[string]int)
	
	for i, ch := range s{
		if inWord == false{
			if( isAlpha(ch) ){
				inWord = true;	
				startOfWord = i
			}else{
					
			}
			
		}else{
			
			if( isAlpha(ch) ){
					
			}else{
				
				inWord = false
				word := s[startOfWord : i]
				//fmt.Println(word)
				ret[word] += 1
			}
		}
		
		
	}
	if inWord {
		word := s[startOfWord:]
		//fmt.Println(word)
		ret[word] += 1
	}
	
	
	return ret
}*/


//使用库函数
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s){ //strings.Fields按照空白字符切分
		m[word] += 1
	}
	return m
	
}


func main() {
	wc.Test(WordCount)
}
