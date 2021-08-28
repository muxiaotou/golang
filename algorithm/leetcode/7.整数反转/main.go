package main

import "fmt"

//int, 32位系统是4个byte，32位，64位系统是8个byte，64位
func fanzhuan(m int) int {
	var j int
	for m != 0 {
		j = j*10 + m%10
		m /= 10
	}

	if j > 1<<31-1 || j < -(1<<31) {
		return 0
	}
	return j
}

func main() {
	m := 123
	fmt.Println(fanzhuan(m))

	n := -123
	fmt.Println(fanzhuan(n))

	s := 2147483649
	fmt.Println(fanzhuan(s))
}
