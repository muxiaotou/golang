package main

import "fmt"

func removeDuplicate(s []int) (int, []int) {
	if len(s) < 2 {
		return len(s), s
	}

	var t []int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			t = append(t, s[i])
		} else {
			if t[len(t)-1] != s[i] { //如果tmp slice最后一位不是s[i]，则存入，否则不做处理
				t = append(t, s[i])
			}
		}
	}
	return len(t), t
}

func main() {
	s := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicate(s))
}
