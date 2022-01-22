//+build ignore

//此题第一想到就是暴力解法，然后其他文章上面提到动态规划，所以练习时两种方法都学习一下
//此题仅仅买卖一次，因此结果肯定是卖出-买入最大差值

package main

import "fmt"

//暴力解法，比题目要求多输出了买卖的时间点
func gupiao1(s []int) (int, []int) {
	var res int
	ind := make([]int, 2)
	for i := 0; i < len(s); i++ { //从index 0 开始依次遍历
		for j := i + 1; j < len(s); j++ { //从第一次for循环的下一个index开始遍历
			tmp := s[j] - s[i] //计算差值
			if tmp > res {     //如果比之前的差值大，则进行替换
				res = tmp
				ind = []int{i + 1, j + 1}
			}
		}
	}
	return res, ind
}

//动态规划，其实动态规划仅仅一种把复杂问题拆解成若干个子问题的思想，而不是一种特定的算法。
func gupiao2(s []int) int {
	var res int
	min := s[0]
	for i := 1; i < len(s); i++ { //相比暴力解法少了一层for循环
		tmp := s[i] - min //计算当前位置和最小的数的差值
		if tmp > res {    //如果差值比已有结果res大，则替换
			res = tmp
		}
		if s[i] < min { //继续找最小值
			min = s[i]
		}
	}
	return res
}

func main() {
	s := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(gupiao1(s))

	s = []int{7, 6, 4, 3, 1}
	fmt.Println(gupiao1(s))

	s = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(gupiao2(s))
}
