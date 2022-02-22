//+build ignore

//解题思路：题目要求会给出32位的二进制数
//最终结果res与二进制数n末位|后，res左移，n右移
package main

import "fmt"

func reversebi(n uint32) (res uint32) {
	for i := 0; i < 32; i++ { //循环32次
		res <<= 1           //可以先将最终结果右移一位，末尾与n的末尾进行位或操作
		res = res | (n & 1) //n 位于 1，判断出n的最后一位是1或者0，然后与res的末位进行位或，判断出事1或者0
		n >>= 1             //n左移，即将最低位剔除，因为最低位已经转换给了res
	}
	return
}

func main() {
	var n uint32 = 0b00000010100101000001111010011100
	fmt.Println(reversebi(n))

	var n1 uint32 = 0b11111111111111111111111111111101
	fmt.Println(reversebi(n1))
}
