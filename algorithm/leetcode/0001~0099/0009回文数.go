//+build ignore
package main

import (
	"fmt"
	"strconv"
)

//结题思路
//负数一定不是回文数
//可以转化成string，反转后进行比较
//可以整数反转，进行比较

func palindnum1(num int) bool {

	if num < 0 {
		return false
	}

	s := strconv.Itoa(num)          //数字转成字符串
	for i := 0; i < len(s)/2; i++ { //遍历一半即可
		if s[i] != s[len(s)-i-1] { //注意收、尾索引，有一次不满足回文要求，就不是回文，返回false
			return false
		}
	}
	return true
}

func palindnum2(num int) bool {
	if num < 0 {
		return false
	}
	var res int
	tmp := num //定义一个临时遍历带入for循环当中计算，因为后面需要转换后的数字和源数字比较
	for tmp > 0 {
		res = tmp%10 + res*10
		tmp = tmp / 10
	}
	return num == res //可以直接返回一个表达式，不用再写if else判断了
}

func main() {
	s := 123456
	fmt.Println(palindnum1(s))

	s = 123321
	fmt.Println(palindnum1(s))
	fmt.Println(palindnum2(s))

	s = -11
	fmt.Println(palindnum1(s))
	fmt.Println(palindnum2(s))

	s = 0
	fmt.Println(palindnum1(s))
	fmt.Println(palindnum2(s))
}
