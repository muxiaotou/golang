//+build ignore

//结题思路：就最简单的，按照数学方法进行相加，直到sum<10
package main

import "fmt"

func add258(num int) (res int) {
	for {
		for num > 0 { //当前num的各位之和
			res += num % 10
			num /= 10
		}
		if res >= 10 { //判断res是否小于10，小于的话return返回即可，不小于的话赋值给num，继续下一轮循环
			num = res
			res = 0
			continue
		} else {
			return res
		}
	}
	return 0
}

func main() {
	fmt.Println(add258(38))

	fmt.Println(add258(1119))
}
