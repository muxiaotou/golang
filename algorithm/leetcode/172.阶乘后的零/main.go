package main

import "fmt"

//观察乘数因子，比如6！= 6 * 5 * 4 * 3 * 2 * 1，只有2*5会导致阶乘计算结果的结尾为0，
//而因子2的数量要远比因子5的数量多。于是，阶乘结果0的个数取决于因子中5的个数。

func zerocount(n int) int {
	count := 0
	for n > 0 {
		count = count + n/5
		n /= 5
	}
	return count
}

func main() {
	fmt.Println(zerocount(5))
}
