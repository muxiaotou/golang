//+build ignore
package main

import "fmt"

//结题思路：给定链表是排好序的，因此遍历去重即可。当linklist.Next ！= nil 时，判断linklist.value与linklist.Next.value是否相等，相等则linklist.Next=linklist.Next.Next
type link83 struct {
	Value int
	Next  *link83
}

func dumplicatelink(l *link83) *link83 {
	cur := l
	for cur.Next != nil {
		if cur.Value == cur.Next.Value { //仅当相邻节点value相等时去重
			cur.Next = cur.Next.Next
		}
		cur = cur.Next //for循环时每次下移一个元素
	}
	return l //返回头节点
}

func printlink(l *link83) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}

func main() {
	link1 := &link83{Value: 1}

	link2 := &link83{Value: 2}
	link1.Next = link2

	link3 := &link83{Value: 2}
	link2.Next = link3

	link4 := &link83{Value: 4}
	link3.Next = link4

	tmp := dumplicatelink(link1)
	printlink(tmp)

}
