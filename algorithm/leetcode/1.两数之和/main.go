package main

import "fmt"

//简单的两次循环，第一个从i开始，第二个从i+1开始，确保有相同数时，不至于取的位置一样
func sum1(s []int, t int) []int {
	var w []int
	for i, _ := range s {
		for j := i + 1; j <= len(s)-1; j++ { //注意j的取值
			if s[i]+s[j] == t {
				w = append(w, i, j)
				return w
			}
		}
	}
	return nil
}

//使用一次循环，定义一个map，key：value分别对应slice的值和键，如果查找到求和预期的目标，返回index
func sum2(s []int, t int) []int {
	tmp := make(map[int]int)
	for i, j := range s {
		target := t - j
		_, ok := tmp[target]
		if ok {
			return []int{tmp[target], i} //此处注意一下
		}
		tmp[j] = i //此处注意一下
	}
	return nil
}

func main() {
	s := []int{2, 7, 11, 15}
	t := 9
	fmt.Println(sum1(s, t))
	fmt.Println(sum2(s, t))

	s = []int{3, 3}
	t = 6
	fmt.Println(sum1(s, t))
	fmt.Println(sum2(s, t))

}
