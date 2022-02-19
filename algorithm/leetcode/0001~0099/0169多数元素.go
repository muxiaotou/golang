//+build ignore

//解题思路：
//方法一：定义一个map，遍历数组，数组的值作为map的key，当有相同元素时，map的value+1，判断map的value是否大于数组长度/2
//方法二：对数组进行排序，长度大于数组长度/2的元素肯定就在index为数组长度/2的索引位置

package main

import (
	"fmt"
	"sort"
)

func multielement1(s []int) int {
	sort.Ints(s) //默认升序排序
	return s[len(s)/2]
}

func multielement2(s []int) int {
	count := make(map[int]int)
	for _, v := range s {
		count[v]++
		if count[v] > len(s)/2 {
			return v
		}
	}
	return 0
}

func main() {

	s := []int{1, 1, 1, 2, 6, 6, 6, 6, 6}
	fmt.Println(multielement1(s))
	fmt.Println(multielement2(s))
}
