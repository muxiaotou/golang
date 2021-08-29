package main

import "fmt"

func generate(nums int) [][]int {
	result := [][]int{}

	for i := 0; i < nums; i++ { //i是行数
		tmp := []int{}
		for j := 0; j < i+1; j++ { //j是每行的那个列数
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else if i > 1 {
				tmp = append(tmp, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, tmp)
	}
	return result
}

func main() {
	fmt.Println(generate(5))

}
