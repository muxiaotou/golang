//+build ignore
package main

import "fmt"

//罗马字符转阿拉伯数字
//解题思路：
//如果遇到比如4、9这种有两个字符表示的数，就减法

//var symbols1 = []struct {
//	symbol string
//	value  int
//}{
//	{"M", 1000},
//	{"D", 500},
//	{"C", 100},
//	{"L", 50},
//	{"X", 10},
//	{"V", 5},
//	{"I", 1},
//}

//此例子相比0012，是因为不要求有序，且map最适合
//由于for遍历string的value为byte类型，因此定义map的时候使用byte
var symbols1 = map[byte]int{
	'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1,
}

func romatoint(s string) (res int) {
	for i := 0; i < len(s); {
		if i < len(s)-1 && symbols1[s[i+1]] > symbols1[s[i]] { //加入i< len(s)-1 ,是为了防止最后一位roma字符串越界
			res += symbols1[s[i+1]] - symbols1[s[i]]
			i += 2
		} else {
			res += symbols1[s[i]]
			i++
		}
	}
	return
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romatoint(s))

	s = "MMMCMXCIX"
	fmt.Println(romatoint(s))
}
