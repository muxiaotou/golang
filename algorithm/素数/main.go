package main

import (
	"fmt"
)

//素数，一个大于1的整数，如果出了1和它本身外，不能被其他正整数整除的数，就叫素数，也叫质数，能被整除的叫合数

//func IsPrime1(a int) bool {
//	if a == 1 {
//		return false
//	}
//
//	//生成2~a-1的整数，存储到slice当中
//	num := make([]int, a-2)
//	for i := range num {
//		num[i] = i + 2
//	}
//
//	//使用被判断的数a除以2~a-1，看能否被整除，如果不能被整除，则返回true
//	for i := range num {
//		if a%num[i] == 0 {
//			return false
//		}
//	}
//	return true
//}

//最简单的素数查找办法，每个数通过for循环去除以2~a-1进行计算
func IsPrime3(a int) bool {
	if a < 2 {
		return false
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	////生成一个slice，存储100以内的整数
	//num := make([]int, 100)
	//for i := range num {
	//	num[i] = i + 1
	//}
	//
	//for i := range num {
	//	if IsPrime3(num[i]) == true {
	//		fmt.Println(num[i])
	//	}
	//}
	for i := 0; i <= 100; i++ {
		if IsPrime3(i) == true {
			fmt.Println(i)
		}
	}
}
