//+build ingore
package main

import "fmt"

//链表逆序存储每位数字：即最左边是数字的最低位，最右边是数字的最高位
//一种思路
//思路一：同时遍历两个链表，逐位计算他们的和，并与当前位置的进位值相加。如果长度不同，则认为长度短的链表后面有若干个0

//定义单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution1(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode //使用定义的tail来逐步保存链表尾
	n1, n2, carry := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10 //sum表示相加后留在当前位置的熟知，carry表示进位的数值
		fmt.Println("chenli", sum, carry, n1, n2)
		if head == nil {
			head = &ListNode{Val: sum} //使用定义的head来保存链表头，赋值给head
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}

	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return //函数中使用了返回定义的变量head，因此return无序再次指明变量名
}

func Traverse(l *ListNode) (s []int) {
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return
}

func main() {
	//link1 [2, 4, 3]
	l11 := &ListNode{Val: 2}

	l12 := &ListNode{Val: 4}
	l11.Next = l12

	l13 := &ListNode{Val: 3}
	l12.Next = l13

	//link2 [5, 6, 4]
	l21 := &ListNode{Val: 5}

	l22 := &ListNode{Val: 6}
	l21.Next = l22

	l23 := &ListNode{Val: 4}
	l22.Next = l23
	//[7 0 8]
	fmt.Println(Traverse(Solution1(l11, l21)))

	//link3 [0] or [9]
	l31 := &ListNode{Val: 9}
	//link4 [0] or [1]
	l41 := &ListNode{Val: 1}
	//[0] or [0 1]
	fmt.Println(Traverse(Solution1(l31, l41)))

	//	link5 [9, 9, 9, 9, 9, 9, 9]
	l51 := &ListNode{Val: 9}

	l52 := &ListNode{Val: 9}
	l51.Next = l52

	l53 := &ListNode{Val: 9}
	l52.Next = l53

	l54 := &ListNode{Val: 9}
	l53.Next = l54

	l55 := &ListNode{Val: 9}
	l54.Next = l55

	l56 := &ListNode{Val: 9}
	l55.Next = l56

	l57 := &ListNode{Val: 9}
	l56.Next = l57

	//link6 [9, 9, 9, 9]
	l61 := &ListNode{Val: 9}

	l62 := &ListNode{Val: 9}
	l61.Next = l62

	l63 := &ListNode{Val: 9}
	l62.Next = l63

	l64 := &ListNode{Val: 9}
	l63.Next = l64

	//[8 9 9 9 9 9 9 1]
	fmt.Println(Traverse(Solution1(l51, l61)))
}
