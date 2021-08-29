package main

import "fmt"

//最大收益来源，必然是每次跌了就买入，涨到顶峰的时候就抛出。只要有涨峰就开始计算赚的钱，连续涨可以用两两相减累加来计算，两两相减累加，相当于涨到波峰的最大值减去谷底的值。这一点看通以后，题目非常简单。
func gupiao(s []int) int {
	res := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i+1] > s[i] {
			res = res + (s[i+1] - s[i])
		}
	}
	return res
}

func main() {
	fmt.Println(gupiao([]int{7, 1, 5, 3, 6, 4}))
}
