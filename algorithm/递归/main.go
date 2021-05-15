//斐波那契数列  阶乘
package main

import "fmt"

//递归函数必须设置一个锚点来触发递归函数结束并返回结果，否则递归函数会进入无穷无尽的循环当中，最终引发内存溢出

func main() {
	//有一对兔子，从出生后第3个月起(第2个月后)每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问5个月的兔子总对数为多少？
	//手动试算每个月的兔子对数：1,1,2,3,5,8......斐波那契数列呀，F(1)=1, F(2)=1, F(3)=2, F(n)=F(n-1)+F(n-2)当n>2
	var n int = 5
	fmt.Println("rabbit1 count is ", rabbit_born1(n))
	fmt.Println("rabbit2 count is ", rabbit_born2(n, 1, 1))

	//阶乘：N!=N*(N-1)*(N-2)......
	fmt.Println("factorial1 value is ", factorial1(n))
	fmt.Println("factorial2 value is ", factorial2(n, 1))

}

func rabbit_born1(n int) int {
	if n < 3 { //n<3, 触发递归结束的锚点
		return 1
	}
	return rabbit_born1(n-1) + rabbit_born1(n-2)
}

func rabbit_born2(n int, acc1 int, acc2 int) int { //相比普通递归，尾递归它增加两个累加器acc1和acc2，并给出初始的值，即n为1，2 时的结果
	if n == 1 {
		return acc1
	}
	return rabbit_born2(n-1, acc2, acc1+acc2)
}

func factorial1(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial1(n-1)
}

func factorial2(n int, acc int) int {
	if n == 1 {
		return acc
	}
	return factorial2(n-1, acc*n)
}
