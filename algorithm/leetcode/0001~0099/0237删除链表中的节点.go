//+build ignore

//解题思路：
//删除链表的常用办法是定位到待删除的节点的上一节点，修改上一节点的next成next.next，使其指向待删除节点的下一节点即可完成删除操作
//此道题目入参非链表头节点，而是待删除的节点，且此节点非尾节点，因此可以将下一节点赋值给当前节点，丢掉下一节点，待删除的节点充当下一节点来实现

package main

import "fmt"

func deletel(l *link237) {
	l.Value = l.Next.Value
	l.Next = l.Next.Next
}

type link237 struct {
	Value int
	Next  *link237
}

func main() {
	link1 := &link237{Value: 4}

	link2 := &link237{Value: 3}
	link1.Next = link2

	link3 := &link237{Value: 5}
	link2.Next = link3

	link4 := &link237{Value: 6}
	link3.Next = link4

	deletel(link3)

	for link1 != nil {
		fmt.Println(link1.Value)
		link1 = link1.Next
	}
}
