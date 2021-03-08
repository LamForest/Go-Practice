/*
//来源：https://tour.go-zh.org/methods/20

重点理解%v的打印机制
*/

package main

import (
	"fmt"
	"math"
)


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("ErrNegativeSqrt.Error(): %f < 0, cannot be sqrted \n",e)	
	
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		/*
		这里发生了：
			将float64类型的x，强制转换成了ErrNegativeSqrt
 			我们为ErrNegativeSqrt定义了Error()方法
			不过，请注意：Error()只是ErrNegativeSqrt的方法，与float64无关，就算ErrNegativeSqrt实际上是float64类型的
		*/
		return 0, ErrNegativeSqrt(x) 
	}
	
	z := 1.0
	
	for i := 0; i < 5 ; i += 1{
		guess := math.Pow(z, 2)
		diff := math.Abs(guess - x)
		nextZ := z - ((guess - x)/(2*z))
		
		fmt.Printf("%2d : z = %.10f, z^2 = %.10f, diff = %.10f, nextZ = %.10f\n", i, z, guess, diff, nextZ)
		z = nextZ
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	
	/*
	这里发生了：
Sqrt(-2)返回值为：0, ErrNegativeSqrt(-2) 
fmt.Println会为每个参数x调用Printf(%v, x)
当Printf遇到%v时，若x实现了某几个接口，那么%v不会打印x的实际对象，而是：
4. If an operand implements the error interface, the Error method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).

5. If an operand implements method String() string, that method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
	*/
	fmt.Println(Sqrt(-2))
	fmt.Println("to be exited")
}
