package main

import "fmt"

//暴力写法
func sqrt1(x int) int {
	res := 0
	for !((res*res <= x) && ((res+1)*(res+1) > x)) {
		res++
	}
	return res
}

//二分法
func sqrt2(x int) int {
	i, j := 0, (x/2)+1
	for i < j {
		mid := (i + j) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			i = mid + 1
		} else {
			j = mid - 1
		}

	}
	return j
}

func main() {
	fmt.Println(sqrt1(5))
	fmt.Println(sqrt2(15))
}
