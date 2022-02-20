//+build ignore
package main

import "fmt"

//fmt.Printf()  提供了多种格式化打印，详细的罗列一下
//fmt.Println()  将十进制、二进制、八进制、十六进制的十进制打印出来

func main() {
	//十进制数转换其他进制
	s := 11111 //自动推导，数字的进制是十进制，与生活习惯保持一致

	//%b  二进制形式显示
	fmt.Printf("%b\n", s)
	tmp := fmt.Sprintf("%b", s)
	fmt.Println(tmp)
	fmt.Printf("%T\n", tmp)

	//%o  八进制形式显示
	fmt.Printf("%o\n", s)

	//%x  十六进制形式显示
	fmt.Printf("%x  %T\n", s, s)

	fmt.Println("----------------------")

	//八进制数转换其他进制
	s1 := 025547 //0开头的数字，自动推导为八进制

	//%d  %b  %x  分别转换成十进制、二进制、十六进制
	fmt.Printf("%d\n", s1)
	fmt.Printf("%b\n", s1)
	fmt.Printf("%x\n", s1)

	fmt.Println("----------------------")
	//十六进制转换其他进制
	s2 := 0x2b67 //0x开头的数字，自动推导为十六进制

	//%d %b %o 分别转换成十进制、二进制、八进制
	fmt.Printf("%d\n", s2)
	fmt.Printf("%b\n", s2)
	fmt.Printf("%o\n", s2)

	fmt.Println("----------------------")
	//二进制转换其他进制
	s3 := 0b11111 //0b开头的数字，自动推导为二进制
	fmt.Println(s3)
	fmt.Printf("%T,%#b", s3, s3)
}
