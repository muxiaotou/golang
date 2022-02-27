//+build ignore

//结题思路：判断是否有重复元素，即出现非一次
//方法一：引入map，当map当中存在当前元素即判定为重复出现
//方法二：排序后，迭代遍历，如果下一个元素与当前元素相同，则判定为出现重复元素
package main

import (
	"fmt"
	"sort"
)

func repeat1(n []int) bool {
	nmap := make(map[int]struct{}) //value是struct{}
	for _, v := range n {
		if _, ok := nmap[v]; ok { //核心语句，存在于map当中即表示重复
			return false
		} else {
			nmap[v] = struct{}{} //将元素存到map当中，value使用空struct，这样节省内存
		}
	}
	return true
}

func repeat2(n []int) bool {
	sort.Ints(n) //先排序一下
	//fmt.Println(n)
	for i := 0; i < len(n)-1; i++ { //每次下标向后移动一位
		//fmt.Println(n[i], n[i+1])
		if n[i] == n[i+1] { //如果下一个元素与当前元素相等，则判定为重复
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(repeat1([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(repeat2([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	fmt.Println(repeat1([]int{1, 2, 3, 4}))
	fmt.Println(repeat2([]int{1, 2, 3, 4}))
}
