//+build ignore
package main

import "fmt"

//解题思路，使用栈

//特别注意map当中的key和value定义，key为右括号，value为左括号
var pairs = map[int32]int32{ //后面用range了，所以是int32，如果用for i:=0的话就可以定义为byte或者uint8
	')': '(',
	'}': '{',
	']': '[',
}

func isVaild(s string) bool {
	if len(s) < 1 || len(s)%2 == 1 { //如果给定的string为空或者个数为基数，则直接返回false
		return false
	}
	stack := []int32{}
	for _, v := range s {
		if pairs[v] > 0 { //v比较关键，如果v是右括号，则进入，知识点：map当中不存在的key的value为对应value的零值，本题即0
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				//stack长度为0，则表示没有左括号，必定是无效的
				//stack[len(stack)-1]表示栈当中的最后一个元素，必定是左括号
				//pairs[v]则表示这个有括号v对应的左括号
				return false
			}
			stack = stack[:len(stack)-1]
		} else { //如果v是左括号，则入栈
			stack = append(stack, v)
		}
	}
	return len(stack) == 0 //最终判断栈中是否还有元素残留
}

func main() {
	s := "()[]{}"
	fmt.Println(isVaild(s))

	s = "([)]"
	fmt.Println(isVaild(s))
}
