package main

import (
	"fmt"
)

//string reverse
func reverse1(s string) string {
	l := len(s)
	t := make([]rune, l)
	s1 := []rune(s)
	for i := l - 1; i >= 0; i-- {
		t = append(t, s1[i])
	}
	return string(t)
}

func reverse2(s string) string {
	l := len(s)
	var t string
	for i := 0; i < l; i++ {
		t = t + fmt.Sprintf("%c", s[l-i-1])
	}
	return t
}

func reverse4(s []string) []string {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

//slice reverse
func reverse3(s []string) []string {
	l := len(s)
	for from, to := 0, l-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
	fmt.Println(s)
	return s
}

func main() {
	s := "abcdefghiglsld"
	fmt.Println(reverse1(s))
	fmt.Println(reverse2(s))
	s1 := []string{"abc", "def"}
	fmt.Println(reverse3(s1))
	s2 := []string{"a", "e", "f"}
	fmt.Println(reverse4(s2))

}
