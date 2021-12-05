//+build ignore
package main

import "fmt"

//解题思路：此题比较简单，但是需要考虑一个地方就是链表长度不一
//链表问题继续采用虚拟链表头的方法，以避免单独去处理头节点
//终止条件，有一个链表的next为nil，即可以终止排序了

type linkN1 struct {
	value int
	next  *linkN1
}

func mergelink(l1, l2 *linkN1) *linkN1 {
	dummylink := &linkN1{}
	cur := dummylink
	for l1 != nil && l2 != nil { //两个链表的next均不为nil时比较大小排序
		fmt.Println(l1, l2, cur.value)
		if l1.value > l2.value {
			cur.next = l2
			l2 = l2.next
		} else {
			cur.next = l1
			l1 = l1.next
		}
		cur = cur.next //这一句很关键，cur要不停的挪动
	}
	if l1 == nil { //如果l1为nil是，则cur的next为l2，不用再排序了
		cur.next = l2

	}
	if l2 == nil {
		cur.next = l1
	}
	return dummylink.next
}

func diskplink(l *linkN1) {
	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

func main() {
	link1 := &linkN1{value: 1}
	link2 := &linkN1{value: 2}
	link1.next = link2
	link3 := &linkN1{value: 4}
	link2.next = link3

	link11 := &linkN1{value: 1}
	link12 := &linkN1{value: 3}
	link11.next = link12
	link13 := &linkN1{value: 4}
	link12.next = link13

	fmt.Println(mergelink(link1, link11))
	//diskplink(mergelink(link1, link11))  这个和上一句println不能同时存在，因为println已经排序了，再次打印就陷入死循环了
}
