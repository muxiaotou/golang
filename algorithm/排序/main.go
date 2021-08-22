package main

import (
	"fmt"
)

//冒泡排序，每次相邻的元素按照约定条件确认是否进行调换，按照由小到大进行排列
func Sort1(s []int) []int {
	l := len(s)
	var mid bool                 //加个判断，确认第一次是否有过排序
	for j := l - 1; j > 0; j-- { //从第一次开始，每次比较相邻两个数比较的总次数，倒序，j=l-1确保后续序列最后的那个数字一次排序后已确定，不再参与下一次的排序
		for i := 0; i < j; i++ { //重点在于i<j，因为后续的已经排好了，所以没必要再全部重新比较一遍了，i=0，每次比较从0开始
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				mid = true
			}
		}
		if !mid { //如果第一次没有排序，则说明已经是按照预期排序了，不用排了
			fmt.Println("bubble sorted 1, return")
			return s
		}
	}
	return s
}

//冒泡排序，按照由大到小排列
func Sort2(s []int) []int {
	l := len(s)
	var mid bool
	for i := 0; i < l-1; i++ { //第一次排完后，首个序列后续不再参与排序
		for j := l - 1; j > i; j-- { //每次从最后位置开始比较，直至外层i控制的序列为止
			if s[j] > s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				mid = true
			}
		}
		if !mid {
			fmt.Println("bubble sorted 2, return")
			return s
		}
	}
	return s

}

//冒泡排序，按照由小到大排列，参照Sort2的写法，仅仅只是在比较大小的地方改变一下，即可以自由选择如何排序
func Sort7(s []int) []int {
	l := len(s)
	var mid bool
	for i := 0; i < l-1; i++ { //第一次排完后，首个序列后续不再参与排序
		for j := l - 1; j > i; j-- { //每次从最后位置开始比较，直至外层i控制的序列为止
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				mid = true
			}
		}
		if !mid {
			fmt.Println("bubble sorted 2, return")
			return s
		}
	}
	return s

}

//选择排序，通过两两相互比较，选出预期值，然后进行换位，由小到大排列
func Sort3(s []int) []int {
	l := len(s)
	for i := 0; i < l-1; i++ {
		min := s[i]
		minIndex := i
		for j := i + 1; j < l; j++ {
			if s[j] < min {
				min = s[j]
				minIndex = j
			}
		}
		if i != minIndex {
			s[i], s[minIndex] = s[minIndex], s[i]
		}
	}
	return s
}

//选择排序，通过两两相互比较，选出预期值，然后进行换位，由大到小
func Sort4(s []int) []int {
	l := len(s)
	for i := 0; i < l-1; i++ {
		max := s[i]
		maxIndex := i
		for j := i + 1; j < l; j++ {
			if s[j] > max {
				max = s[j]
				maxIndex = j
			}
		}
		if i != maxIndex {
			s[i], s[maxIndex] = s[maxIndex], s[i]
		}
	}
	return s
}

//插入排序，每次把一个数插入到已经拍好序的梳理里面形成新的排好序的序列，由小到大排序
func Sort5(s []int) []int {
	l := len(s)
	for i := 1; i < l; i++ {
		daipai := s[i] //待排的数，首次是第2个元素
		left := i - 1  //待排的数左边的数，首次是第1个元素
		for ; left >= 0 && daipai < s[left]; left-- {
			s[left+1] = s[left] //如果右边的数小于左边的，则把左边的往右边移动一位
		}
		s[left+1] = daipai //移动完后，需要把待排的数赋值到最左边，此处由于for里面多减了1，所以需要加1
	}
	return s
}

//插入排序，由大到小
func Sort6(s []int) []int {
	l := len(s)
	for i := 1; i < l; i++ {
		daipai := s[i] //待排的数，首次是第2个元素
		left := i - 1  //待排的数左边的数，首次是第1个元素
		for ; left >= 0 && daipai > s[left]; left-- {
			s[left+1] = s[left] //如果右边的数大于左边的，则把左边的往右边移动一位
		}
		s[left+1] = daipai //移动完后，需要把待排的数赋值到最左边，此处由于for里面多减了1，所以需要加1

	}
	return s
}

//快速排序，大体步骤如下：
//1.先从数列中取一个数作为基数，一般取第一个数
//2.f分区过程，将比基数大的放到右边，比基数小的放到左边
//3.再对左右区间重复第二步，直到各区间只有一个数
func partition(arr []int, begin, end int) int {

}

func main() {

	s := []int{6, 4, 5, 2, 3, 1}
	fmt.Println("Sort1")
	fmt.Println(Sort1(s))
	fmt.Println(Sort1(s))
	fmt.Println("Sort2")
	fmt.Println(Sort2(s))
	fmt.Println(Sort2(s))

	s1 := []int{6, 4, 5, 2, 3, 1}
	fmt.Println("Sort3")
	fmt.Println(Sort3(s1))
	fmt.Println("Sort4")
	fmt.Println(Sort4(s1))

	s2 := []int{6, 4, 5, 2, 3, 1}
	fmt.Println("Sort5")
	fmt.Println(Sort5(s2))
	fmt.Println(Sort6(s2))

	s3 := []int{6, 4, 5, 2, 3, 1}
	fmt.Println("Sort7")
	fmt.Println(Sort7(s3))

}
