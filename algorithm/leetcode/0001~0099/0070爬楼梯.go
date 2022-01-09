//+build ingore

//此题的核心时推导出规律，要不然觉得不错下手
//假设5层楼梯，由于每次可以爬1或者2层，则分两种情况：
//1）首次爬1层，则还有4层要爬
//2）首次爬2层，则还有5层要爬
//以上1）+2）的方法总和即是总的爬5层的总方法数\
//此类问题递归进行

package main

import "fmt"

func stair(num int) int {
	if num == 0 || num == 1 { //0个台阶时，1种办法，就是不用爬，哈哈哈哈
		return 1
	}
	return stair(num-1) + stair(num-2)
}

//还有一种解法，叫动态规划，所有递归问题动态规划也能解决，我理解动态规划就是固定最基本的base，然后循环根据base(可以理解为基准)逐层计算
func stair1(num int) int {

	if num == 0 || num == 1 {
		return 1
	}

	temp := make([]int, num+1) //+1是为了存储0个台阶的情况
	temp[0] = 1                //base
	temp[1] = 1                //base
	for i := 2; i <= num; i++ {
		temp[i] = temp[i-1] + temp[i-2]
	}
	return temp[num]
}

func main() {
	fmt.Println(stair(5))
	fmt.Println(stair1(5))
}
