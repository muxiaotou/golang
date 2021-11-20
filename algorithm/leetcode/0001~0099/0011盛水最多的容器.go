//+build ignore
package main

import (
	"fmt"
	"time"
)

//两种解题思路吧，一种是暴力破解，一个是动态规划
//暴力破解：
//两层for循环，分别从i和i+1开始依次计算取最大值
//动态规划(不太好理解)：
//假定左右是i和j，则maxwater = min(i,j)*(j-i)，那就i从0开始，j从len(s)-1开始，i和j小的向里异动
//这样就可能出现更大的maxwater,相比暴力破解，可以减少多次的遍历

func container1(s []int) (l, r, res int) {
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ { //两层for循环
			if min1(s[i], s[j])*(j-i) > res {
				res = min1(s[i], s[j]) * (j - i)
				l, r = i, j
			}
		}
	}
	return
}

func min1(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func container2(s []int) (l, r, res int) {
	for i, j := 0, len(s)-1; i < j; { //限定i从0开始，j从len(s)-1开始
		if res < min1(s[i], s[j])*(j-i) {
			res = min1(s[i], s[j]) * (j - i)
			l, r = i, j
		}
		if s[i] < s[j] {
			i++
		} else {
			j--
		}
	}
	return
}

func main() {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	bT := time.Now()
	fmt.Println("begin", bT)
	fmt.Println(container1(s))
	eT := time.Now()
	fmt.Println("end", eT)
	fmt.Printf("%v\n", time.Since(bT))
	bT = time.Now()
	fmt.Println(container2(s))
	fmt.Printf("%v\n", time.Since(bT))

	s = []int{1, 1}
	fmt.Println(container1(s))
	fmt.Println(container2(s))
}
