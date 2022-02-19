package main

import (
	"fmt"
	"strconv"
)

//strconv实现了十进制转成其他进制的函数
//分别演示二进制、八进制、十六进制、二十六进制、三十六进制与十进制之间的互转

//const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//
//func strtodec(s string, base int) int {
//
//}

//二进制转十进制
func bintodec(s string) (res int) {
	for i := 0; i < len(s); i++ {
		res = res*2 + int(s[i]-'0') //核心语句，由于for循环遍历的uint8类型，实际对应的
		//fmt.Printf("%T", s[i])
	}
	return
}

//十进制转二进制
func dectobin(s int) (res string) {
	for ; s > 0; s /= 2 {
		l := s % 2
		res = strconv.Itoa(l) + res
	}
	return
}
func main() {
	fmt.Println(bintodec("1101"))
	fmt.Println(dectobin(13))
	fmt.Println('9')
	fmt.Printf("%T", '9')
}
