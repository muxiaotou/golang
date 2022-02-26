//+build ignore

//解题思路：
//一个数字最终考虑有两种情况：一是最终每个位置上面数字的平方和为1，二是数字的平方和出现无限循环，因此结果当中对着两种情况进行判断
package main

import "fmt"

func happynum(m int) bool {
	result := make(map[int]bool) //定义一个map，用来存储每次的的结果

	for result[m] == false { //用来判断，上次计算后的数值是否已经出现过，如果出现过就需要退出了，因为成为死循环了
		var res, n int
		for n = m; n > 0; {
			res += (n % 10) * (n % 10) //计算平方根
			n /= 10                    //高位退低位，继续计算
		}
		fmt.Println(res)
		if res == 1 {
			return true
		} else {
			result[m] = true //如果结果不为1，存储到map当中
			m = res
		}
		fmt.Println(m, res, result[res], !result[m])
		fmt.Println(result)
	}
	return false
}

func main() {
	fmt.Println(happynum(19))
	fmt.Println(happynum(2))
}
