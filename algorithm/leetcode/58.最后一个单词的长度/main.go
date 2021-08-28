package main

import "fmt"

func returnAfterlenth(s string) int {
	tmp := []byte(s)
	var start int
	var end int
	for i := len(s) - 1; i >= 0; i-- {
		if string(tmp[i]) == " " {
			start = i
		} else if (string(tmp[i]) != " ") && (string(tmp[i-1]) == " ") { //当前不是空格，下一个是
			end = i
			break
		}
	}
	return start - end
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(returnAfterlenth(s))
}
