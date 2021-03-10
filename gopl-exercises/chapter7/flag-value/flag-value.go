package main

import (
	"flag"
	"fmt"
	"strings"
)

type distribution struct {
	name string
}

//Set对传入的参数进行解析
//解析后的结果通过指针写入dis
func (dis *distribution) Set(s string) error {

	switch strings.ToLower(s) {
	case "fedora":
		dis.name = "fedora"
		return nil
	case "debian", "ubuntu":
		dis.name = "debian"
		return nil
	}
	return fmt.Errorf("Not supported Linux distribution : %s", s)
}

func (dis *distribution) String() string {
	return fmt.Sprintf("[Arch = %s]", dis.name)
}

var flagDistribution distribution

func init() {
	flagDistribution = distribution{"None"}
	flag.CommandLine.Var(&flagDistribution, //设置默认值
		"dis", //参数名
		"Input a linux distribution; Currently support: fedora, debian based like Ubuntu") //参数提示 usage

}

func main() {
	flag.Parse()
	fmt.Printf("argv distributon = {%s}\n", flagDistribution)
	if flagDistribution.name == "None" {
		fmt.Println("argv distribution cannot be empty")
	}

}
