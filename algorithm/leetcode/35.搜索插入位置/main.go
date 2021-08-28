package main

import "fmt"

func searchInsert(s []int, t int) int {
	for i, v := range s {
		if v >= t {
			return i
		}
	}
	return len(s)
}

func main() {
	s := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(s, 2))
	fmt.Println(searchInsert(s, 5))
	fmt.Println(searchInsert(s, 7))
	fmt.Println(searchInsert(s, 0))
	s1 := []int{1}
	fmt.Println(searchInsert(s1, 0))
}
