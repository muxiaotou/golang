//+build ignore
package main

import "fmt"

//此题目解法也挺多的，大概记录了一种思路来解题
//原有字符串按照Z字型重新排列后，再按照新排列的行，组成新的字符串
//因此可以定义一个slice，用行号做index，新行的字符逐个组成串做slice的value

func Zsort(s string, row int) (res string) {
	var currentRow int
	var direct bool
	tmp := make([]string, row)
	for _, v := range s {
		tmp[currentRow] += string(v)                //字符串可以直接相加，以达到拼接的效果
		if currentRow == 0 || row-currentRow == 1 { //如果是首行和尾行，则开始改变方向，默认direct从0行开始是false
			//逐步到尾行后变成true，又知道首行才变为false
			direct = !direct
		}
		if direct {
			currentRow++
		} else {
			currentRow--
		}
	}
	for _, v := range tmp {
		res += v //字符串直接相加，以进行拼接
	}
	return
}

func main() {
	s := "abcdefghijklmn"
	fmt.Println(Zsort(s, 3))

	s = "abcdefghijklmn"
	fmt.Println(Zsort(s, 4))
}
