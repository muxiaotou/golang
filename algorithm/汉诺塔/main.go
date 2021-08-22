package main

import (
	"fmt"
	"math"
)

var total = 0
var n = 5
var a = "A"
var b = "B"
var c = "C"

func main() {
	fmt.Printf("total count should is %d \n", int(math.Pow(float64(2), float64(n)))-1) //2^n -1
	tower(n, a, b, c)
}

//可以理解printf打印就是一次手动去操作盘子的行为，最终跟手动验算的步骤对应
func tower(n int, a, b, c string) {
	if n == 1 {
		total = total + 1
		fmt.Printf("count %d, from %s to %s \n", total, a, c)
		return
	}

	tower(n-1, a, c, b)
	total = total + 1
	fmt.Printf("count %d, from %s to %s \n", total, a, c)
	tower(n-1, b, a, c)
}
