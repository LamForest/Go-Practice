/*
https://tour.go-zh.org/flowcontrol/8

练习：循环与函数
为了练习函数与循环，我们来实现一个平方根函数：用牛顿法实现平方根函数。

计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：

z -= (z*z - x) / (2*z)
重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。

在提供的 func Sqrt 中实现它。无论输入是什么，对 z 的一个恰当的猜测为 1。 要开始，请重复计算 10 次并随之打印每次的 z 值。观察对于不同的值 x（1、2、3 ...）， 你得到的答案是如何逼近结果的，猜测提升的速度有多快。

提示：用类型转换或浮点数语法来声明并初始化一个浮点数值：

z := 1.0
z := float64(1)
然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数大于还是小于 10。 尝试改变 z 的初始猜测，如 x 或 x/2。你的函数结果与标准库中的 math.Sqrt 接近吗？

（*注：* 如果你对该算法的细节感兴趣，上面的 z² − x 是 z² 到它所要到达的值（即 x）的距离， 除以的 2z 为 z² 的导数，我们通过 z² 的变化速度来改变 z 的调整量。 这种通用方法叫做牛顿法。 它对很多函数，特别是平方根而言非常有效。）

输出：
 0 : z = 1.0000000000, z^2 = 1.0000000000, diff = 1.0000000000, nextZ = 1.5000000000
 1 : z = 1.5000000000, z^2 = 2.2500000000, diff = 0.2500000000, nextZ = 1.4166666667
 2 : z = 1.4166666667, z^2 = 2.0069444444, diff = 0.0069444444, nextZ = 1.4142156863
 3 : z = 1.4142156863, z^2 = 2.0000060073, diff = 0.0000060073, nextZ = 1.4142135624
 4 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
 5 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
 6 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
 7 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
 8 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
 9 : z = 1.4142135624, z^2 = 2.0000000000, diff = 0.0000000000, nextZ = 1.4142135624
1.414213562373095
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	
	for i := 0; i < 10 ; i += 1{
		guess := math.Pow(z, 2)
		diff := math.Abs(guess - x)
		nextZ := z - ((guess - x)/(2*z))
		
		fmt.Printf("%2d : z = %.10f, z^2 = %.10f, diff = %.10f, nextZ = %.10f\n", i, z, guess, diff, nextZ)
		z = nextZ
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}