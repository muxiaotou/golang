//+build ignore

//解题思路：常规迭代遍历的思路即可，删除时记得采用虚拟头节点，这样可以很方便的处理头节点
package main

import "fmt"

type link203 struct {
	Value int
	Next  *link203
}

func deletelink203(l *link203, n int) (res *link203) {
	h := &link203{Next: l}
	//思路是站在当前位置head，去判断当前位置的下一个位置head.Next是否需要删除掉
	for head := h; head.Next != nil; { //当head.Next不为nil时，进入for循环
		fmt.Println(head.Next)
		if head.Next.Value == n { //判断当前head.Next的Value是否和目标值相等，相等则删除
			head.Next = head.Next.Next
		} else { //当head.Next的不需要被删除时，则异动head到head.Next
			head = head.Next
		}
	}
	return h.Next
}

func printlink203(l *link203) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}

func main() {
	head1 := &link203{Value: 1}
	head2 := &link203{Value: 2}
	head1.Next = head2

	head3 := &link203{Value: 6}
	head2.Next = head3

	head4 := &link203{Value: 2}
	head3.Next = head4

	fmt.Println(deletelink203(head1, 2))

	printlink203(head1)
}
