//+build ignore

package main

import "fmt"

//本题解题思路：加1即可，但是由于可能原有数字时9，需要进位，依次倒序遍历

func addone(num []int) []int {
	length := len(num)
	for i := length - 1; i >= 0; i-- { //倒序开始遍历
		num[i]++            //当前位+1
		if num[i]/10 == 0 { //+1后如果小于10，则直接返回结果即可
			return num
		}
		num[i] = num[i] % 10 //+1后由于结果不小于10，因此除10后，个位数赋予当前位置，并且进一位也就是+1，不会出现+2、+3等等情况
	}
	num = append([]int{1}, num...) //走到此处还没结束，那就是所有位全是9了，因此append 1进去即可
	return num
}

func main() {
	num := []int{9, 9, 9, 9}
	fmt.Println(addone(num))

	num = []int{1, 5, 7, 4}
	fmt.Println(addone(num))
}
