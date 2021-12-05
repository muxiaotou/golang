//+build ignore
package main

import (
	"fmt"
	"sort"
)

//思路，同0015三数之和思路一样，四个数的话就先固定两个数，多一层for循环而已

//foursum1在每次对四个数的取值时，for循环已经进行了去重处理，就是逻辑判断多了一点
func foursum1(num []int, target int) (res [][]int) {
	if len(num) < 3 {
		return
	}

	sort.Ints(num) //使用双指针，因此排序是必须要做的
	for i := 0; i < len(num)-3; i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		for j := len(num) - 1; j > i+2; j-- {
			if j < len(num)-1 && num[j] == num[j+1] {
				continue
			}

			left := i + 1
			right := j - 1
			for left < right {
				if num[i]+num[left]+num[right]+num[j] > target {
					right--
				} else if num[i]+num[left]+num[right]+num[j] < target {
					left++
				} else {
					res = append(res, []int{num[i], num[left], num[right], num[j]})
					for left < right && num[left] == num[left+1] {
						left++
					}
					for left < right && num[right] == num[right-1] {
						right--
					}
					left++ //两头固定，当有满足条件子串，left和right同时移动，只移动一个没有意义，因为移动了一个，即使出现满足的子串，也跟已有的是重复的
					right--
				}
			}
		}
	}
	return
}

//foursum2 在对四个数取值时没有做去重，但是在append到结果的slice时进行了遍历去重处理
func foursum2(num []int, target int) (res [][]int) {
	if len(num) < 3 {
		return
	}
	sort.Ints(num)
	for i := 0; i < len(num)-3; i++ {
		for j := len(num) - 1; j > i+2; j-- {
			left := i + 1
			right := j - 1
			for left < right {
				tmp := num[i] + num[left] + num[right] + num[j]
				if tmp > target {
					right--
				} else if tmp < target {
					left++
				} else {
					flag := true
					for _, v := range res { //去重的关键在此，如果有相同的则flag标记为false，下面的append则不会发生
						if v[0] == num[i] && v[1] == num[left] && v[2] == num[right] && v[3] == num[j] {
							flag = false
						}
					}
					if flag {
						res = append(res, []int{num[i], num[left], num[right], num[j]})
					}
					left++
					right--
				}
			}
		}
	}
	return
}

func main() {
	num := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(foursum1(num, 0))
	fmt.Println(foursum2(num, 0))

	num = []int{2, 2, 2, 2, 2, 2}
	fmt.Println(foursum1(num, 8))
	fmt.Println(foursum2(num, 8))
}
