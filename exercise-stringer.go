//来源：https://tour.go-zh.org/methods/18
//练习：Stringer
//通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

//例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。


package main

import "fmt"
import "strconv"
import "strings"

type IPAddr [4]byte

func (addr IPAddr) String() string{
	//拼接字符串，真麻烦。。。
	var strArray = make([]string, len(addr))
	for ind, val := range addr{
		strArray[ind] = strconv.Itoa(int(val))
	}
	
	ret := strings.Join(strArray, ".")
	ret = "IPAddr.String() convert: " + ret
	
	return ret
	
}

func (addr IPAddr) foo() {
	fmt.Println("func foo")	
}

type MyStringer interface{
	String() string
}

type FooI interface{
	foo()
}

type Object interface {}

func MyPrint(obj Object ){
   var str MyStringer
	
	str,ok := obj.(MyStringer)
	fmt.Printf("%v %T\n",str, str)//输出：IPAddr.String() convert: 127.0.0.1 main.IPAddr
	if ok == true{
		fmt.Printf("%v\n",str.String()) //输出：IPAddr.String() convert: 127.0.0.1
	}else{
		fmt.Println("ok == false")
	}
	
}


//这个函数为了验证，结构体X实现了A,B 接口，可以在接口之间任意的转化
func MyFoo(obj Object ){
   
	
	fooI,ok := obj.(FooI)
	
	
	// 输出：Object -> FooI ok, IPAddr.String() convert: 127.0.0.1, main.IPAddr
	if ok == true{
		fooI.foo()
		fmt.Printf("Object -> FooI ok, %v, %T\n",fooI, fooI)
	}else{
		fmt.Println("Object -> FooI fail, ok == false")
	}
	
	
	addr,ok := obj.(IPAddr)
	
	// 输出：Object -> IPAddr ok, IPAddr.String() convert: 127.0.0.1, main.IPAddr
	if ok == true{
		fmt.Printf("Object -> IPAddr ok, %v, %T\n",addr, addr)
	}else{
		fmt.Println("Object -> IPAddr fail, ok == false")
	}

}

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	MyFoo(hosts["loopback"])
	
	fmt.Println("- - - - - - - -")
	
	MyPrint(hosts["loopback"])
	
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
