//+build ignore
package main

import "fmt"

//两种思路：
//思路一：两个for循环迭代，第一个for从元素1开始，第二个for从元素1+1位置开始迭代，直到找到和为target的两个数字
//思路二：一个for循环，借助定义一个map来存不匹配的数字

func solution1(s []int, t int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == t {
				//var r []int
				//r = append(r, i, j)
				//return r
				return []int{i, j}
			}
		}
	}
	return nil
}

func solution2(s []int, t int) []int {
	//m := make(map[int]struct{})
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		j := t - s[i]
		if v, ok := m[j]; ok {
			//var r []int
			//r = append(r, v, i)
			//return r
			return []int{v, i}
		}
		m[s[i]] = i
		//m[s[i]]= struct {}{}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(solution1(nums, target))
	fmt.Println(solution2(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(solution1(nums, target))
	fmt.Println(solution2(nums, target))

	nums = []int{3, 3, 3}
	target = 6
	fmt.Println(solution1(nums, target))
	fmt.Println(solution2(nums, target))
}
