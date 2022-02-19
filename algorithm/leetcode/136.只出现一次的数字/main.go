package main

import "fmt"

//写的有bug
//func checkOne(s []int) int {
//	var t bool
//	for i := 0; i < len(s)-1; i++ {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				t = true
//				continue
//			}
//		}
//		if !t {
//			return s[i]
//		}
//	}
//	return 0
//}

//异或，同为0，异为1，转为二进制进行异或运算
func checkOne1(s []int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = result ^ s[i] //相同的数字异或就成0，剩下的就是
	}
	return result
}

func main() {
	fmt.Println(2 ^ 1)
	fmt.Println(checkOne1([]int{4, 1, 2, 1, 2}))
	//fmt.Println(checkOne([]int{1, 2, 1, 2, 4}))
}
