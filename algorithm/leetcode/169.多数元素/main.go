package main

import "fmt"

//这个写的有问题，因为例子当中1也是出现了3次
//func max(s []int) int {
//	for i := 0; i < len(s)-1; i++ {
//		count := 0
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				count++
//				if count >= len(s)/2 { //多熟元数是指个数大于n/s的元素
//					return s[i]
//				}
//			}
//		}
//	}
//	return 0
//}

func max(s []int) int {
	count := make(map[int]int)
	for _, i := range s {
		count[i] = count[i] + 1
		fmt.Println(count)
		if count[i] > len(s)/2 {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(max([]int{1, 2, 2, 1, 1, 2, 2}))
}
