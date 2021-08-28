package main

import (
	"fmt"
)

//字符串普通位置的比较
func huiwen(s string) string {
	l := len(s)
	var isnoH bool
	if l == 1 {
		isnoH = false
	}

	for i := 0; i <= (l-1)/2; i++ { //核心在这里，注意有等号
		if s[i] != s[l-i-1] {
			isnoH = true
			break
		}
	}
	if isnoH {
		fmt.Println("bu shi hui wen")
		return "N"

	}
	fmt.Println("shi hui wen")
	return "Y"

}

//字符串反转之后比较
func huiwen1(s string) string {
	tmp := []rune(s)
	l := len(s)
	for from, to := 0, l-1; from < to; from, to = from+1, to-1 {
		tmp[from], tmp[to] = tmp[to], tmp[from]
	}
	if string(tmp) == s {
		fmt.Println("shi hui wen")
		return "Y"
	}

	fmt.Println("bu shi huiwen")
	return "N"
}

//数字回文判断，不转换成字符串
func huiwen2(t int) string {

	var tmp int
	t1 := t

	for t1 != 0 { //核心在这里
		tmp = tmp*10 + t1%10
		t1 /= 10
	}
	if tmp == t {
		fmt.Println("shi hui wen shu zi")
		return "Y"
	}
	fmt.Println("bu shi hui wen shu zi")
	return "N"
}

func main() {
	s := "abccba"
	s1 := "abc"
	fmt.Println(huiwen1(s))
	fmt.Println(huiwen1(s1))

	t := 1232
	t1 := 12321
	fmt.Println(huiwen2(t))
	fmt.Println(huiwen2(t1))
}
