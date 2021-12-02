//+build ignore

package main

import (
	"fmt"
	"sort"
)

//解题思路和0015类似，也是先排序，然后left从index 0开始遍历，不同的是，此题目当中恰好仅存在一个解

func threesum1(num []int, target int) (res int) {
	sort.Ints(num)
	res = num[0] + num[1] + num[2]
	for i := 0; i < len(num)-2; i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		mid := i + 1
		right := len(num) - 1
		for mid < right {
			sum := num[i] + num[mid] + num[right]
			if sum < target {
				mid++
			} else if sum > target {
				right--
			} else {
				return sum
			}
			if abs(res-target) > abs(sum-target) {
				res = sum
			}
		}
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	num := []int{-1, 2, 1, -4}
	fmt.Println(threesum1(num, 1))

	num = []int{0, 0, 0}
	fmt.Println(threesum1(num, 1))

	fmt.Println("hello")
}
