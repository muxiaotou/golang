//+build ignore

//解题思路：
//1）遍历数组，然后声明一个新的存储空间来存储已遍历过的元素，value可以使用false或者true来控制，最终true的即为仅仅出现一次的元素，此方法是需要额外声明存储空间的
//2）排序后，比对
//3）使用异或运算
//
package main

import (
	"fmt"
	"sort"
)

func checkone1(s []int) int {
	var res = make(map[int]int) //定义一个map
	for _, i := range s {
		res[i]++ //以s的元素未key，出现一次则map的value+1
	}
	for k, v := range res { //遍历map，当value为1的key即为仅仅出现1次的数字
		if v == 1 {
			return k
		}
	}
	return 0
}

func checkone2(s []int) int {
	sort.Ints(s) //先排序一下
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ { //此处要防止i越界，所以i小于len(s)-1
		if s[i] != s[i+1] { //判断一下是否相邻的元素想到，不相等则直接return
			return s[i]
		}
		i++ //此处index+1，for循环里面也index+1，实际加了2，则是for里面一次遍历的不长为2
	}
	return s[len(s)-1]
}

func checkone3(s []int) int {
	var res int
	for _, i := range s {
		res = res ^ i //一切数字与0异或，还为此数字，相同的数字异或为0
	}
	return res
}

//4）暴力解法，就是两个for循环，挨个对比，当两次for循环index相等时，skip
//func checkone4(s []int) int {
//	var flag bool
//	for i := 0; i < len(s); i++ {
//		for j := 0; j < len(s); j++ {
//			if i == j {
//				break
//			}
//			if s[i] == s[j] {
//				flag = true
//			}
//		}
//		if !flag {
//			return s[i]
//		}
//
//	}
//	return 0
//}

func main() {

	s := []int{4, 1, 2, 2, 1}
	fmt.Println(checkone1(s))
	fmt.Println(checkone2(s))
	fmt.Println(checkone3(s))
}
