package main

import "fmt"

func inttoroman(num int) string {
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	r, i := "", 0
	//需要从左边开始进行逐一换算
	for num != 0 {
		//如果v当中的数字大于实际的num，则继续v的下一个数字进行比较
		for v[i] > num {
			i++
		}
		num = num - v[i]
		r = r + s[i]
	}
	return r
}
func main() {
	num := 1994
	fmt.Println(inttoroman(num))
}
