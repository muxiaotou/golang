//+build ignore

//解题思路：
//位移进行计算
package main

import "fmt"

func binonecount(n uint32) (res int) {
	for i := 0; i < 32; i++ { //循环遍历32次
		if n&1 > 0 { //n的末位与1进行位与操作，判断末位是否为1
			res++ //是1，则进行计数
		}
		n >>= 1 //n向左移动一位
	}
	return
}

func main() {
	var n uint32 = 0b00000000000000000000000000001011
	fmt.Println(binonecount(n))
}
