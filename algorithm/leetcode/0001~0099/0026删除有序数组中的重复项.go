//+build ignore
package main

import "fmt"

//解题思路：
//题目要求原地删除重复项，因此无序定义新的存储结构，根据索引来处理，写了两种方法，第二种比较好，采用快慢指针的方法

//这个函数打印了最终的不重复的序列，但是定义了新序列
func removedup(s []int) (res []int, l int) {
	if len(s) == 0 {
		return
	}
	res, l = append(res, s[0]), 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		res = append(res, s[i])
		l++
	}
	return
}

//这个函数跟leetcode上面题目一致了，仅仅打印去重后的个数
func removedup1(s []int) (length int) {
	//if len(s) == 0 {
	//	return
	//}
	//length++
	//for i := 1; i < len(s); i++ {
	//	if s[i] == s[i-1] {
	//		continue
	//	}
	//	length++
	//}
	if len(s) == 0 {
		return
	}
	slow, fast := 0, 1  //快慢指针定义
	for fast < len(s) { //快指针不能超过length
		if s[fast] != s[slow] { //快、慢值不相等时，slow+1
			//s[slow] = s[fast]
			slow++
		}
		//每次都将fast赋值给slow，slow可能在此次for循环时没有+1，但是fast每次for循环时都+1
		s[slow] = s[fast]
		fast++
	}
	fmt.Println(s[:slow+1])
	return slow + 1
}

func main() {
	s := []int{1, 1, 3, 3, 5, 6, 7, 7, 8}
	fmt.Println(removedup(s))

	s1 := []int{1, 1, 3, 3, 5, 6, 7, 7, 8}
	fmt.Println(removedup1(s1))

}
