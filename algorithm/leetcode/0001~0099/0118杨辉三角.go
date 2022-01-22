//+build ignore
package main

import "fmt"

//解题思路：核心的办法就是根据题目的说明来找规律：
//1.每行的元素列数等于行号(行号从1开始)
//2.没列的元素值等于上一行同列的元素值+上一行同位置前一列的元素值
//3.每行首、尾列元素均为1

func yanghui1(n int) {
	//res := [][]int{}
	//res := make([][]int, n)  这样定义会index越界？？？？，这样的话append的时候前面元素已经初始化成0 了
	res := make([][]int, 0) //最终的print的结果
	fmt.Println(len(res), cap(res))
	for i := 0; i < n; i++ { //最外层变量行，此处n是函数的传参，指代行数
		//tmp := []int{}
		tmp := make([]int, 0)
		for j := 0; j < i+1; j++ { //此处的列其实跟行有关系，因此第n行，总共n个元素
			if j == 0 || j == i { //首、尾列元素均为1，对应规律3
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, res[i-1][j]+res[i-1][j-1]) //对应规律2
			}
		}
		res = append(res, tmp)

	}
	fmt.Println(res)
}

func main() {
	yanghui1(3)
}
