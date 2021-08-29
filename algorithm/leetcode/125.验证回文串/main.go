package main

import (
	"fmt"
	"strings"
)

func huiwen(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) { //判断左边，过滤掉特殊字符和空格等
			i++
		}
		for i < j && !isChar(s[j]) { //判断右边，过滤掉特殊字符和空格等
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func main() {
	s := "A man, a plan, a canal: Panama"
	s1 := "race a car"
	fmt.Println(huiwen(s))
	fmt.Println(huiwen(s1))
}
