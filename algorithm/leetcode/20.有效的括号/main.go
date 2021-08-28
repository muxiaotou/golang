package main

import "fmt"

func isValid(s string) bool {
	m := map[string]string{")": "(", "]": "[", "}": "{"} //注意定义的地方key和value

	if len(s) == 0 {
		return true
	}
	var tmp []string
	for i := 0; i < len(s); i++ {
		if len(tmp) == 0 {
			tmp = append(tmp, string(s[i]))
			fmt.Println("1 ", tmp)
		} else {
			if tmp[len(tmp)-1] == m[string(s[i])] { //核心，如果s当中出现后括号，就去检查tmp当中是否有前括号
				tmp = tmp[:len(tmp)-1]
				fmt.Println("2 ", tmp)
			} else {
				tmp = append(tmp, string(s[i]))
				fmt.Println("3 ", tmp)
			}
		}
	}
	fmt.Println(tmp)
	if len(tmp) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
