package main

import "fmt"

//公共前缀是指从第一个字符就开始要相同

func longprefix(s []string) string {

	//先找出slice最短的那个string
	minStr := s[0]

	for i, _ := range s {
		if len(s[i]) < len(minStr) {
			minStr = s[i]
		}
	}

	result := ""

	for i := 0; i < len(minStr); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j][i] != s[j+1][i] {
				return result
			}
		}
		result = result + string(minStr[i])
	}
	return result
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longprefix(s))

}
