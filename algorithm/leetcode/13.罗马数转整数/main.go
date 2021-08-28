package main

import (
	"fmt"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romantoint(s string) int {
	l := len(s)
	num, lastnum, total := 0, 0, 0
	//需要从右边开始倒着进行换算
	for i := l - 1; i >= 0; i-- {
		//char := s[i]  //byte
		char := s[i : i+1] //string
		num = roman[char]
		//如果小于右边的数，那就是减，如果大于右边的数，那就加
		if num < lastnum {
			total = total - num
		} else {
			total = total + num
		}
		lastnum = num
	}
	return total
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romantoint(s))
}
