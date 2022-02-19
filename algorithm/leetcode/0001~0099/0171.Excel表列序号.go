//+build ignore
package main

import "fmt"

//解题思路：
//题目本质上是一个26进制转10进制的题目，因此按照规律进行转换方法的书写

//二十六进制转十进制
func convert1(s string) (ans int) {
	//for _, v := range s {
	//	fmt.Println(i)
	for i := 0; i < len(s); i++ {
		ans = ans*26 + int(s[i]-'A') + 1 //最关键的一行，需要根据s[i]-'A'计算出实际字母表示的数值
	}
	return
}

func main() {
	fmt.Println(convert1("ZY"))
}
