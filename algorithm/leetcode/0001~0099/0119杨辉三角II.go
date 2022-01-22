//+build ignore

//解题思路：继续利用杨辉三角的特点
//1）每行首、尾列元素值为1
//2）每行的列值等于前一行同列+前一行同列前一列 的值总和
//3）每行的列数等于他的行数值(从1开始计数)

//本题只需要打出最后一行的所有值，则可以仅仅只存储第n-1行的值，用于输出n行的值
//与0118题目不同的是，0118需要将从1~n行的结果都保存下来，0119题目则不需要，仅仅保存n-1行即可，方便计算出第n行
package main

import "fmt"

func yanghui2(n int) {
	res := make([]int, 0)
	for i := 1; i <= n; i++ { //i作为行遍历，遍历1~n，即依次需要去计算没行的所有值
		tmp := make([]int, 0)
		for j := 1; j <= i; j++ { //j作为列遍历，遍历1~i，即依次去计算每行的各个列值
			if j == 1 || j == i {
				tmp = append(tmp, 1) //根据特点，每行首、尾列的值为1
			} else {
				tmp = append(tmp, res[j-1]+res[j-2]) //根据特点2计算，此处j-1和j-2分别代表存储在res当中对应位置的index，但是原题中index是从0开始的！！！
			}
		}
		res = tmp
	}
	fmt.Println(res)
}

func main() {
	yanghui2(5)
}
