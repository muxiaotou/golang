//+build ingore
package main

import "fmt"

//此题目很简单，由于最后一个单词前后可能都会出现空格，因此想到了两种办法：
//1）倒序遍历，记录非空格的字符start index，直到再次出现空格的end index，计算差值
//2）倒序记录非空格index，然后从此index开始累加非空格字符的个数

func wordcount1(s string) int {
	lenth := len(s)
	var start, end int
	for i := lenth - 1; i >= 0; i-- {
		//从组后一个字符倒序开始判断是否为空，直到不为空即未start对应的index
		if s[i] == ' ' { //s[i]是byte类型，因此用‘ ’来表示空字符
			fmt.Println(s[i], i)
			start = i
		} else if s[i] != ' ' && s[i-1] == ' ' { //判断当前字符非空，下一字符为空时，记录end对应的index，并且需要退出for循环
			end = i
			break
		}
	}
	fmt.Println(start, end)
	return start - end
}

func wordcount2(s string) (count int) {
	lenth := len(s)
	var index int
	for i := lenth - 1; i >= 0; i-- {
		if s[i] != ' ' {
			index = i
			break
		}
	}

	for i := index; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return
}

func main() {
	s := " fly me to the moon   "
	fmt.Println(wordcount1(s))
	fmt.Println(wordcount2(s))
}
