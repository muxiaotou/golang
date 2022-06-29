package main

import (
	"strings"
)

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

//加这个为了验证代码覆盖率时，预期覆盖率未达100%
func Add(a, b int) (result int) {
	return a + b
}

//func Add1(x, c int) (result int) {
//	for i := 0; i < c; i++ {
//		if x%2 != 0 {
//			x += 5
//		} else {
//			x /= 2
//		}
//		fmt.Println(x)
//	}
//	return x
//}
//
//func main() {
//	fmt.Println(Add1(6, 20))
//}
