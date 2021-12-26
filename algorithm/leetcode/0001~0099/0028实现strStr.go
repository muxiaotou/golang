//+build ignore

//解题思路：
//字符串匹配此类问题有两种方法，一种是暴力解法，一种叫KMP解法
package main

import "fmt"

//暴力解法
//l表示长字符串，s表示短的子串
func str1(l, s string) int {
	for i := 0; i < len(l); i++ { //遍历长字符串
		//fmt.Println(l[i : len(s)+i])
		if l[i:len(s)+i] == s {
			return i
		}
	}
	return -1
}

//KMP解法，主要应用在字符串匹配上(这个解法不好理解......)
//暴力解法，不匹配时从i+1开始重新进行匹配
//而KMP算法，核心思想是当出现字符串不匹配时，可以记录一部分之前已经匹配的文本内容，利用这些信息避免从头再去做匹配

func main() {
	fmt.Println(str1("hello", "ll"))
}
