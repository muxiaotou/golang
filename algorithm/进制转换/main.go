package main

import (
	"fmt"
	"strconv"
)

//strconv实现了可以将各种字符串转换成二进制、八进制、十六进制、三十六进制等的数字，也可以将数字转换成对应进制的字符串 ParseInt   FormatInt
//fmt包的Printf格式化，可以使用%d\%b\%o\%x可以将数字以十进制、二进制、八进制、十六进制形式打印(仍是int类型)
//fmt包的Sprintf可以使用%d\%b\%o\%x可以将数字转换成以十进制、二进制、八进制、十六进制形式表示的字符串(结果是string类型)

//分别演示二进制、八进制、十六进制、二十六进制、三十六进制与十进制之间的互转

//https://blog.csdn.net/flyfreelyit/article/details/80097035   进制转换  （搜索关键字：golang 八进制转十进制 算法）

//可以看看fmt包，简单的快速实现了进制转换

//二进制转十进制
func bintodec(s string) (res int) {
	for i := 0; i < len(s); i++ {
		res = res*2 + int(s[i]-'0') //核心语句，由于for循环遍历的uint8类型，实际对应的
		//res*2 表示原先结果逢2进位，因此乘以2
		//fmt.Printf("%T", s[i])
	}
	return
}

//十进制转二进制
func dectobin(s int) (res string) {
	for ; s > 0; s /= 2 {
		l := s % 2                  //核心语句，余数就是最终的二进制结果
		res = strconv.Itoa(l) + res //核心语句，十进制除二，直到商为0，每次除法时的余数倒数排列，因此约除到最后的余数放置在最终结果的最前面
	}
	return
}

//八进制转十进制
func octtodec(s string) (res int) {
	for i := 0; i < len(s); i++ {
		res = res*8 + int(s[i]-'0') //核心语句，从最高位开始计算，每次res的结果都乘以8，即代表左移进位一次
	}
	return
}

//十进制转八进制
func dectooct(s int) (res string) {
	for ; s > 0; s /= 8 {
		l := s % 8
		res = strconv.Itoa(l) + res //核心语句，与二进制原理类似
	}
	return
}

//十六进制转十进制
func hextodec(s string) (res int) {
	digist := map[byte]int{'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15}
	for i := 0; i < len(s); i++ {
		if _, ok := digist[s[i]]; ok { //核心语句，判断是否在定义的map当中
			res = res*16 + digist[s[i]]
		} else {
			res = res*16 + int(s[i]-'0')
		}
	}
	return
}

//十进制转十六进制
func dectohex(s int) (res string) {
	digist := map[int]string{10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F"}
	for ; s > 0; s /= 16 {
		l := s % 16
		if l >= 10 { //核心语句，判断数字是否超过了9，超过后需要使用定义的map当中的字符来表示
			res = digist[l] + res
		} else {
			res = strconv.Itoa(l) + res
		}
	}
	return
}
func main() {
	//十进制、二进制
	fmt.Println(bintodec("1101"))
	fmt.Println(dectobin(13))
	fmt.Println('9')
	fmt.Printf("%T\n", '9')

	//十进制、八进制
	fmt.Println(octtodec("77"))
	fmt.Println(dectooct(63))

	//十六进制、十进制
	fmt.Println(hextodec("FF"))
	fmt.Println(dectohex(255))

}
