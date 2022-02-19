//+build ignore
package main

import "fmt"

//解题思路：简单粗暴的办法就是两两遍历，然后找到相交点，此处不用两层for循环依次遍历，可以先讲链表一存入map，然后for遍历第二个链表即可

type link160 struct {
	Val  int
	Next *link160
}

func createlink160(s []int) *link160 {
	if len(s) == 0 {
		return nil
	}
	head := &link160{Val: s[0]}
	tmp := head
	for _, v := range s[1:] {
		tmp.Next = &link160{Val: v}
		tmp = tmp.Next
	}
	return head
}

func printlink160(l *link160) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func intersection(l1, l2 *link160) int {
	if l1 == nil || l2 == nil {
		return 0
	}
	flag := make(map[*link160]bool)
	for l1 != nil {
		flag[l1] = true
		l1 = l1.Next
	}
	fmt.Println(flag)

	for l2 != nil {
		fmt.Println(flag[l2])
		fmt.Println(l2)
		if flag[l2] {
			return l2.Val
		}
		l2 = l2.Next
	}
	return 0
}

func main() {
	//s1 := []int{4, 2, 1, 3, 5}
	//s2 := []int{8, 3, 5, 9}
	//l1 := createlink160(s1)
	//l2 := createlink160(s2)

	//上述方法定义的linklist不会有相交点，因为相交点的内存地址是一样的，上述可以改造一下，判断两个链表是否有相同的元素值
	s1 := &link160{Val: 4}

	s2 := &link160{Val: 2}
	s1.Next = s2

	s3 := &link160{Val: 1}
	s2.Next = s3

	s4 := &link160{Val: 3}
	s3.Next = s4

	l1 := &link160{Val: 8}
	l1.Next = s3 //相交链表的特点就在于这句，两个链表有元素的地址会指向同一内存地址

	printlink160(s1)
	printlink160(l1)
	fmt.Println(intersection(l1, s1))
}
