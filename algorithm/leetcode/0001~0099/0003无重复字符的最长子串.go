//+build ignore
package main

import "fmt"

//在golang中，没有专门的字符类型，一般单个字符(字母或汉子)声明为byte或者rune类型，并用单引号括起来
//字符的本质是整数，字符是可以和数字混合在一起直接进行运算

//网上例子比较多，眼花缭乱的，我就先记录一种方法吧，核心思想都是“滑动窗口”，即
//滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现重复数字，就缩小左边界，直到重复的字符
//移出了左边界，然后继续异动滑动窗口的右边界。

//本题返回的是最长子串的长度和首先出现的最长子串，不会打印出后面出现的相同长度的子串

func substring1(s string) (int, string) {
	var bit [256]bool       //临时标记每个字符是否出现过
	var s1 string           //首次出现的最长子串
	result, l, r := 0, 0, 0 //result表示子串长度，l表示左窗口位置，r表示右窗口位置

	for l < len(s) {
		if bit[s[r]] {
			bit[s[l]] = false
			l++
		} else {
			bit[s[r]] = true
			r++
		}
		if result < r-l {
			result = r - l
			s1 = s[l:r] //此处的r最大是len(s)
		}
		if l+result >= len(s) || r >= len(s) {
			break
		}
	}
	return result, s1
}

func main() {
	s := "abcabcbb"
	fmt.Println(substring1(s))

	s = "bbbbb"
	fmt.Println(substring1(s))

	s = "fgacbcbae"
	fmt.Println(substring1(s))

	s = "abcdefg"
	fmt.Println(substring1(s))

	s = "aaaaa"
	fmt.Println(s[:5])
}
