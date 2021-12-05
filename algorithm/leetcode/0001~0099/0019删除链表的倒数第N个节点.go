//+build ignore
package main

import "fmt"

//链表可以借鉴虚拟头节点的方法，这样就不用再单独考虑删除实际头节点的逻辑了
//这个题目最快的方法是快慢指针，即快的先走N+1步，然后快慢一起走，直到快的为nil,这时慢的在倒数第N+1节点上，然后节点的Next指向N的next
type linkN struct {
	value int
	Next  *linkN
}

//初始化一个struct,四种方法：
//rect1 := new(Rect)
//rect2 := &Rect{}
//rect3 := &Rect{0, 0, 100, 200}
//rect4 := &Rect{width:100, height:200}

func deletelinkN(l *linkN, N int) (res *linkN) {
	res = &linkN{Next: l}
	fast, slow := res, res //快、慢指针，指向虚拟头节点
	i := 0
	for fast != nil {
		fast = fast.Next
		if i > N { //fast走了N+1步，slow才开始走，目的是为了让slow最终停留在倒数第N+1节点上
			slow = slow.Next
		}
		i++
	}
	slow.Next = slow.Next.Next //等号左边是倒数第N+1节点，等号右边是第N-1节点，这样就可以删除倒数第N节点了
	return res.Next
}

func main() {
	link1 := &linkN{value: 1}
	link2 := &linkN{value: 2}
	link1.Next = link2
	fmt.Println(deletelinkN(link1, 2))
}
