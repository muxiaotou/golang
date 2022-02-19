//+build ignore
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]int, 5)
	fmt.Println(cap(s), len(s))
	//s[10] = 10
	//fmt.Println(cap(s), len(s))
	//fmt.Println(s)

	s1 := "aBC e 1 223 F"
	fmt.Println(strings.ToLower(s1))
}
