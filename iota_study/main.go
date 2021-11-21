package main

import "fmt"

//iota代表了const声明块的行索引(下标从0开始)
//每增加一行，iota的取值增1，即便声明块中没有使用iota也是如此
//单行语句中，即便出现多个iota，iota的取值也保持不变

const (
	bit0 = 1 << iota
	bit1
	bit2
	bit3 = iota //iota变成行号3
)

const (
	mask0, sub0 = 1 << iota, 1<<iota - 1 //多个iota时，多个iota还是代表行号
	mask1, sub1
	mask2, sub2 = iota, iota
)

func main() {
	//预期   1,2,4,3
	fmt.Println(bit0, bit1, bit2, bit3)
	//预期 1,0,2,1,2,2
	fmt.Println(mask0, sub0, mask1, sub1, mask2, sub2)
}
