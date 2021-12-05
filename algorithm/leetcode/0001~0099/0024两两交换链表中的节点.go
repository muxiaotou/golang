//+build ignore
package main

import "fmt"

//解题思路：
//继续采用虚拟链表头，dummy head，以避免单独处理头结点

type linkex struct {
	value int
	next  *linkex
}

func exchangelink(l *linkex) *linkex {
	dummy := &linkex{next: l}
	cur := dummy
	for cur.next != nil && cur.next.next != nil {
		n1, n2 := cur.next, cur.next.next

		n1.next = n2.next //第一个元素的next指向第二个元素的next，要把交换后和没有交换的链表连起来
		cur.next = n2     //虚拟头节点的next指向第二个元素
		n2.next = n1      //第二个元素的next指向第一个元素
		cur = n1          //cur指向交换后的第二个元素(原来是第一个元素)，充当后续两两交换元素是的dummy元素
	}
	return dummy.next
}

func diskplink1(l *linkex) {

	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

func main() {
	list1 := &linkex{value: 1}
	list2 := &linkex{value: 2}
	list1.next = list2
	list3 := &linkex{value: 3}
	list2.next = list3
	list4 := &linkex{value: 4}
	list3.next = list4

	//fmt.Println(exchangelink(list1))
	diskplink1(exchangelink(list1))
}
