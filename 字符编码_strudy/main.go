//https://joyohub.com/go/go-character/ 好的文章
package main

import "fmt"

//golang语言没有字符类型，字符只是整数的特殊用例，使用rune和byte作为别名，int32和uint8
//go的字符串使用了UTF-8的编码来表示，所以要明确好Unicode和ASCII码的区别
//Unicode仅仅只是一个符号集，并未规定二进制代码应该如何存储
//UTF-8规定了一种统一的编码方式，即Unicode码的一种实现方式，规定了：
//1）单字节符号，字节的第一位为0，后面7为为这个符号的Unicode码，和ASCII码是相同的
//2）n>1字节的符号，第一个字节的前n位为1，第n+1位为0，后面字节的前两位一律设为10，剩下的没有提及的二进制位，全部为这个字符的Unicode码

func main() {
	fmt.Println('A') //输出ascii码
	fmt.Printf("%c\n", 'A')

	var b rune = '陈'
	fmt.Println(b)        //输出了多个byte的ascii码
	fmt.Printf("%c\n", b) //输出的ascii码，汉字出现乱码
	fmt.Printf("%U\n", b) //输出unicode码
}
