package main

import "fmt"

func plusone(s []int) []int {

	for i := len(s) - 1; i >= 0; i-- {
		s[i]++
		if s[i] == 10 {
			s[i] = 0
			if i == 0 {
				s = append([]int{1}, s...)
				return s
			}
			continue
		} else {
			return s
		}
	}
	return nil
}

func main() {
	s := []int{9, 9, 9}
	fmt.Println(plusone(s))

	s1 := []int{3, 3, 2, 2}
	fmt.Println(plusone(s1))
}
