package main

import "fmt"

func panlouti(n int) int {
	sum := make([]int, n+1)
	sum[0] = 1
	sum[1] = 1

	for i := 2; i <= n; i++ {
		sum[i] = sum[i-1] + sum[i-2]
	}
	return sum[n]
}

func main() {
	fmt.Println(panlouti(5))
}
