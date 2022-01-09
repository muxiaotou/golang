//+build ignore

//结题思路：
//1）从低位开始依次比较排序，但是此方法需要额外声明内存空间来存储，要不然腾挪原有数组就不方便了
//2）从高位开始依次比较，此方法可以不用新声明内存空间，只需要利用最终的有序数组长度为m+n的技巧来实现
package main

import "fmt"

//低位开始比较，但是需要声明新的内存空间来存储
func mergelist1(a, b []int, m, n int) []int {
	//res := make([]int, m+n)
	var res []int
	for i, j := 0, 0; i < m && j < n; {
		if a[i] < b[i] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func main() {

	s1 := []int{1, 2, 4, 6, 7, 7, 8, 9}
	s2 := []int{2, 2, 5, 6, 8, 8, 10}
	fmt.Println(mergelist1(s1, s2, 8, 7))
}
