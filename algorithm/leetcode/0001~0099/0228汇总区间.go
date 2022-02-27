//+build ignore

//结题思路：
//就常规的迭代遍历，比如先定义区间头、尾为同一元素，再判断下一元素是否为当前元素+1，是则尾元素后移，否，则打印，开始下一次判断
package main

import "fmt"

func printrange(num []int) {
	start, end := 0, 0              //先定义头、尾均为0
	for i := 0; i < len(num); i++ { //依次遍历迭代，i<len(num)
		if (i != len(num)-1) && (num[i+1] == num[i]+1) { //此处需要先判定i不为num的最后一个元素
			end = i + 1 //确定了尾
			continue
		}
		fmt.Printf("%d->%d\n", num[start], num[end]) //直接打印首、尾元素
		start, end = i+1, i+1
	}
}

func main() {
	printrange([]int{0, 1, 2, 4, 5, 7})
	printrange([]int{0, 2, 3, 4, 6, 8, 9})
}
