//+build ignore
package main

import (
	"fmt"
	"math"
)

//根据题目要求，使用基本的逻辑判断即可实现，题目要求：
//1.丢弃前置无用的空格
//2.起始可能存在符号‘-’，‘+’，若无符号认为是正数
//3.依次读入下一个字符，非数字或者到达行尾则停止
//4.如果没有读入则为0
//5.数字超范围后截取

//golang  int 类型至少是一个32位的带符号整数，而不是int32的别称，可以使用strconv.IntSize打印int的位数，cpu不同位数不同  32位或者64位

func stringtoint(s string) (res int) {
	//s = strings.TrimSpace(s) //去掉字符串两边空格
	var blank bool
	base := 1 //标记是正数或者负数

	for i, v := range s {
		if v == ' ' && !blank { //跳过字符串前面空格
			continue
		}
		blank = true //空格字符flag，标记前面空格已经跳过，再次出现则不认为是字符串前面的空格了
		s = s[i:]
		break //非空格后退出for
	}

	//判断第一个字符是否是符号位
	if s[0] == '-' {
		base = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, v := range s {
		if v < '0' || v > '9' { //非数字字符则退出for循环
			break
		}
		//fmt.Println(reflect.TypeOf(v))  range的是int32，rune是int32的别名
		//此处v-‘0’很关键，因为v实际是一个byte类型，代表一个ascii码的一个字符
		//'9'不是数字9，而是57，因此需要v和'0'求差，v和‘0’差刚好是v表示的那个数字
		res = res*10 + int(v-'0')
		if res*base > math.MaxInt32 { //const MaxInt32，表示int32位最大值
			return math.MaxInt32
		} else if res*base < math.MinInt32 {
			return math.MinInt32
		}
	}
	return res * base
}

func main() {

	//fmt.Println(int('9' - '0'))
	s := "42"
	fmt.Println(stringtoint(s))

	s = "         -42"
	fmt.Println(stringtoint(s))

	s = "4193 with words"
	fmt.Println(stringtoint(s))

	s = "444  777"
	fmt.Println(stringtoint(s))

	s = "words and 987"
	fmt.Println(stringtoint(s))

	s = "-91283472332"
	fmt.Println(stringtoint(s))

}
