package main

import "fmt"

func searchStr(m, n string) int {
	if len(n) == 0 {
		return 0
	}

	l := len(n)
	for i := 0; i < len(m)-l; i++ { //注意i的大小限制
		if m[i:i+l] == n { //直接比较两个slice
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(searchStr("hello", "ll"))
	fmt.Println(searchStr("hello", ""))
	fmt.Println(searchStr("hello", "dd"))
}
