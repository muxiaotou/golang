//+build ignore

//解题思路：题目要求只考虑字母和数字，但是给定字符串当中还有其他非二者的字符，所以需要考虑去掉的，且要忽略大小写......
//1）遍历+转换+判断：先依次遍历，讲字母和数字放入临时的字符串，然后转换大小写，再逐个的比对是否是回文串
//2）直接双指针进行判断，即首、尾指针开始判断

package main

import (
	"fmt"
	"strings"
)

//最简单的转换、遍历、判断的写法
func huiwen1(s string) bool {
	s = strings.ToLower(s) //先统一转换成小写字母，其中非字母的字符保持不变
	var tmp string
	for _, i := range s { //遍历转换小写后的s
		//fmt.Println(i)
		//if (i >= '0' && i <= '9') || (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
		if (i >= '0' && i <= '9') || (i >= 'a' && i <= 'z') { //判断是数字或者字母，则加入临时的tmp当中
			//fmt.Println("aaa")
			tmp = tmp + string(i)
		}
	}
	fmt.Println(s, tmp)
	for i := 0; i < len(tmp)/2+1; i++ { //遍历到中间的index，然后首、尾判断
		if tmp[i] != tmp[len(tmp)-i-1] {
			return false
		}
	}
	return true
}

//双指针方法
func huiwen2(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isStrNum(s[l]) { //此处用for循环，将左边非字母和数字的字符统统skip掉
			l++
		}
		for l < r && !isStrNum(s[r]) { //此处用for循环，将右边非字符和数字的字符统统skip掉
			r--
		}

		for l < r { //此时，左、右元素是字母或者数字，判断是否相等，判断完后如果相等，则左、右index分别加、减1
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isStrNum(c uint8) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func main() {
	s := "A b C def12 45 54 21FEd C B a"
	fmt.Println(huiwen1(s))

	fmt.Println(huiwen2(s))
}
