package main

import "fmt"

//计算1....100之间的数字总和

//暴力写法，总需要加n次
func sum1(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

//暴力写法，需要用到除法
func sum2(n int) int {
	num := (1 + n) * n / 2
	return num
}

//递归，比较耗资源
func sum3(n int) int {
	if n == 1 {
		return 1
	}
	return sum3(n-1) + n
}

//尾递归
func sum4(n, m int) int {
	if n == 1 {
		return m
	}
	return sum4(n-1, n+m)
}

func main() {
	fmt.Println(sum1(100))
	fmt.Println(sum2(98))
	fmt.Println(sum3(98))
	fmt.Println(sum4(98, 1))
}
