//+build ignore
package main

import "fmt"

//解题思路，罗马数
//比较特殊的是4用IV表示，9用IX表示，40用XL表示，90用XC表示，400用CD表示，900用CM表示
//因此在数字转换罗马数字的时候，定义一个struct来存储一下

//题目的数字1<num,罗马数字最大M，比如3000就是MMM

//type r struct {
//	k int
//	v string
//}
//
//var s = []r{{1000, "M"}}

//定义时赋值，比较简洁
//此例子当中的symbols使用slice，因为slice是有序的，代码当中需要有序的减去最大的值
var symbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func inttoraoma(num int) (res string) {
	for num > 0 { //num在内层for里面做减法，最终会变成<=10,然后逐个转换，最终被拆解后被减成0，因此用num>0来判断是否进行循环
		for _, v := range symbols { //逐个循环symbols，使用num去做减法，获取罗马字符
			if num-v.value >= 0 {
				res += v.symbol
				num -= v.value
				break //每次得到一个罗马数字后退出内层循环，从外出再次进入，主要是防止超过1999的时候出错
			}
		}
	}
	return
}

func main() {
	num := 1999
	fmt.Println(inttoraoma(num))

	num = 5999
	fmt.Println(inttoraoma(num))
}
