package main

import "fmt"

func merge(a []int, m int, b []int, n int) []int {
	for i := m + n; m > 0 && n > 0; i-- {
		if a[m-1] <= b[n-1] {
			a[i-1] = b[n-1]
			n--
		} else {
			a[i-1] = a[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		a[n-1] = b[n-1]
	}
	return a
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(merge(nums1, 3, nums2, 3))
}
