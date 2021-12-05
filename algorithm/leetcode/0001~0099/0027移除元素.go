//+build ignore
package main

import "fmt"

//解题思路，和0026删除重复项思路一样，继续采用快慢指针的方法

func removeitem(num []int, target int) (res []int, length int) {
	if len(num) == 0 {
		return
	}
	slow, fast := 0, 0
	for fast < len(num) {
		if num[fast] != target { //技巧，当与target不相等时，仅仅slow++即可
			num[slow] = num[fast]
			slow++
		}
		//fast每次for均+1
		fast++
	}
	//上面的slow实际上已经多加了1
	return num[:slow], slow
}

func main() {
	num := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeitem(num, 2))
}
