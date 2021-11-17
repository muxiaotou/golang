//+build ignore
package main

import "fmt"

//解题方法有很多，我暂时理解了其中一种解法，叫 中心扩散法
//中心扩散法的基本思想是：遍历每一个下标，以这个下标为中心，利用回文串的中心对称的特点，往两边扩散，看最多能扩散多远

//细节：由于目标回文子串可能是偶数，也可能是奇数，因此对应的中心不一样，解题时兼容两种情况：
//1.传入重合下标，进行中心扩散
//2.传入相邻下标，进行中心扩散

func maxpalindrome(s, res string, i, j int) string {
	//s 是传入的源字符串，res是传入进来已得到的最长子串，i和j传入的下标(相同或者相邻)
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] { //如果传入两个下标字符相等，则是回文
		sub = s[i : j+1]
		//已经判断传入下标是回文，则当前下标左、右各移动一步，继续判断是否扩散后的字符串是否是回文
		i--
		j++
	}
	//如果本次扩散的子串比已得到到的回文子串长度长，则返回本次新得到的回文子串
	if len(sub) > len(res) {
		return sub
	}
	return res
}

func longestpalingrome(s string) (res string) {
	for i := 0; i < len(s); i++ {
		res = maxpalindrome(s, res, i, i)
		res = maxpalindrome(s, res, i, i+1)
	}
	return
}

func main() {
	s := "babad"
	fmt.Println(longestpalingrome(s))

	s = "bbbccccc"
	fmt.Println(longestpalingrome(s))

	s = "b"
	fmt.Println(longestpalingrome(s))

	s = "ac"
	fmt.Println(longestpalingrome(s))

	s = ""
	fmt.Println(longestpalingrome(s))
}
