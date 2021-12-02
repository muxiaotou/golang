//+build ignore
package main

import "fmt"

//思路：
//方法一：此类问题是n个for循环，通常通过回溯法来确定，本质也就是递归函数，但是这个题的递归有点难懂
//方法二：那就多层for循环来逐个遍历拼接吧，核心是保存已经拼接好的的字符，然后和新的继续进行拼接

//定义的技巧
//map的value为slice，方便在给定一个数字时，直接可以返回
//map的key为byte，方便后续代码单证range后，对string的value不用做类似转换
var phonesymbols = map[byte][]string{'2': {"a", "b", "c"}, '3': {"d", "e", "f"}, '4': {"g", "h", "i"},
	'5': {"j", "k", "l"}, '6': {"m", "n", "o"}, '7': {"p", "q", "r", "s"}, '8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"}}

func phoneletter(num string) (res []string) {
	fmt.Println(num)
	if len(num) == 0 { //如果数字为空，则直接返回空的slice
		return []string{}
	}

	res = phonesymbols[num[0]] //{a, b, c}
	for i := 1; i < len(num); i++ {
		tmp := make([]string, 0)
		for i1 := 0; i1 < len(phonesymbols[num[i]]); i1++ { //此处注意len里面的元素的i取上层循环
			for _, v2 := range res {
				tmp = append(tmp, v2+phonesymbols[num[i]][i1]) //此处的冷里面的元素的i取上层循环
			}
		}
		res = tmp
	}
	return
}

func main() {
	fmt.Println(phoneletter("234"))
}
