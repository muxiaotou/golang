//+build ignore

//结题思路
//每次对数字做除法和取余，如果余数不为0，则判定不为2的幂次
package main

import "fmt"

func twopower(num int) bool {
	for num != 1 {
		fmt.Println(num)
		if num%2 != 0 {
			return false
		}
		num /= 2
	}
	return true
}

func main() {
	fmt.Println(twopower(7))
	fmt.Println(twopower(4))
}
