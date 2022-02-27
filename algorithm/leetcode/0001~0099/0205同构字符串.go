//+build ignore

//开始读这个题目的时候一直无法明白是要做什么的，最后才看懂了，就是左边每个字符唯一对应右边一个字符，右边同样也是，右边的一个字符，唯一对应左边的一个字符
//开始想到使用一个map，左边字符为key，右边字符为value，循环左边(左右字符串长度是相等的)，如果key不在map当中则保存到map，如果存在，则判断value是否和右边的字符相等
//后面看其他人的说明，比如会存在右边1个字符对应左边不同字符的情况，因此使用2个map来进行存储、比较

package main

import (
	"fmt"
)

func isisolog(m, n string) bool {
	mmap := make(map[byte]byte)
	nmap := make(map[byte]byte)
	for k, _ := range m {
		if _, ok := mmap[m[k]]; !ok {
			mmap[m[k]] = n[k]
		}
		if _, ok := nmap[n[k]]; !ok {
			nmap[n[k]] = m[k]
		}

		if (mmap[m[k]] != n[k]) || (nmap[n[k]] != m[k]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isisolog("egg", "add"))
	fmt.Println(isisolog("abb", "bar"))
}
