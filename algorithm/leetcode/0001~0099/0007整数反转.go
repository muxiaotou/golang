//+build ignore
package main

import "fmt"

//整数反转，跟字符串反转不太一样，因为要考虑负数，所以考虑数字运算的方法进行，核心思想是整除(/)和求余(%)

func reverseInt(s int) (res int) {
	for s != 0 {
		res = s%10 + res*10 //获取s最低位，加到res的最低位，同时res向左整体移动一位(乘以10)
		s = s / 10          //整除，去掉最低位
	}
	return
}

func main() {
	fmt.Println(reverseInt(-123444))
	fmt.Println(reverseInt(9888098))
}
