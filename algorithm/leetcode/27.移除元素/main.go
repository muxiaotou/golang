package main

import "fmt"

//符合题意，不增加新的数组结构，在原数组上进行修改
func removeElement(s []int, t int) (int, []int) {

	index := 0
	for index < len(s) {
		if s[index] == t {
			s = append(s[:index], s[index+1:]...)
			continue //由于index发生了变化，所以此处index不需要+1
		}
		index++
	}
	return len(s), s
}

//引入了额外的数组空间，不符合题意
func removeElement1(s []int, t int) (int, []int) {
	tmp := make(map[int]int)
	for i, v := range s {
		tmp[i] = v
	}

	fmt.Println(tmp)
	for i := 0; i < len(s); i++ {
		if tmp[i] == t {
			delete(tmp, i)
		}
	}

	var res []int
	fmt.Println(tmp)
	for _, v := range tmp {
		res = append(res, v)
	}
	return len(tmp), res
}

func main() {
	s := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement1(s, 2))

	s1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(s1, 2))
}
