//+build ignore

//结题思路： 质因数 ,就是说约数，就是说一个数字，只能被这几个数字整除
//题目给定质因数是2、3、5，因此判断除以这几个数字取余，如果余数为0，则表示可以被整除，进行除法运算
//比如先除以2，还是先除以3，还是先除以5，其实都无所谓，因为只是先除和后除而已，最终就那几个质因数
package main

import "fmt"

func chounum(num int) bool {
	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(chounum(9))
	fmt.Println(chounum(8))
	fmt.Println(chounum(14))
	fmt.Println(chounum(1))
}
