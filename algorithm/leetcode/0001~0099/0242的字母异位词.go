//build +ignore

//解题思路：考虑到两种方法，第一种是先排序，之后逐位比较；第二种是额外定义一个map，第一个slice变量后存入map，第二个slice遇到相同的就减1

package main

import (
	"fmt"
	"sort"
)

func sorts1(s, t string) bool {
	slist := []byte(s) //string和[]byte可以进行强制类型转换
	tlist := []byte(t)

	sort.Slice(slist, func(i, j int) bool { //string转换成[]byte，后排序
		return slist[i] < slist[j]
	})

	sort.Slice(tlist, func(i, j int) bool {
		return tlist[i] < tlist[j]
	})

	for i, v := range slist { //对排序后的slice进行比较
		if tlist[i] != v {
			return false
		}
	}
	return true
}

func sorts2(s, t string) bool {
	res := make(map[rune]int)
	for _, v := range s {
		res[v]++
	}

	for _, v := range t {
		res[v]--
		if res[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "abcdefg"
	t := "abcdefg"

	fmt.Println(sorts1(s, t))
	fmt.Println(sorts2(s, t))
}
