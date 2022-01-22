//+build ignore
package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(cap(s), len(s))
	s[10] = 10
	fmt.Println(cap(s), len(s))
	fmt.Println(s)
}
