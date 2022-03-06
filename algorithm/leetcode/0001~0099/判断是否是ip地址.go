//+build ignore
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//s := make([]int, 5)
	//fmt.Println(cap(s), len(s))
	////s[10] = 10
	////fmt.Println(cap(s), len(s))
	////fmt.Println(s)
	//
	//s1 := "aBC e 1 223 F"
	//fmt.Println(strings.ToLower(s1))
	//
	//s2 := 0b10111111
	//fmt.Printf("%d\n", s2)

	s3 := "10.226.455.8"
	tmp := strings.Split(s3, ".")
	if len(tmp) != 4 {
		fmt.Println("not is a ip addr[1]")
	} else {
		for _, v := range tmp {
			aaa, _ := strconv.Atoi(v)
			if aaa > 255 {
				fmt.Println("not is a ip addr[2]")
				return
			}
		}
		fmt.Println("is a ip addr")
	}
}
