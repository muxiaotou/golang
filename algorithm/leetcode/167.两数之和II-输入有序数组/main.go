package main

import "fmt"

//前提是一个已经排好序的数组
func sum1(s []int, t int) []int {
	i, j := 0, len(s)-1
	for i < j {
		if s[i]+s[j] == t {
			return []int{i + 1, j + 1}
		}
		if s[i]+s[j] < t { //小于目标值，则左边index+1
			i++
		} else {
			j--
		}
	}
	return nil
}

func main() {
	s := []int{2, 7, 11, 15}
	fmt.Println(sum1(s, 13))
}
