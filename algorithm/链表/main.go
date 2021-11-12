package main

import "fmt"

//https://www.jianshu.com/p/868a219fe2b5
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247492022&idx=1&sn=35f6cb8ab60794f8f52338fab3e5cda5&scene=21#wechat_redirect

//一个最简单的单向链表，链表的创建、插入、遍历、删除、查询、反转

type LinkNode struct {
	Data     int64     //数据
	NextNode *LinkNode //指向下一链表节点的指针
}

//创建链表，头插法
func CreateNodeOnHead(listData []int64) *LinkNode {
	if len(listData) == 0 {
		return &LinkNode{Data: 0} //给部分key赋值，指定key，若给你全部key赋值，无需指定key
	}
	head := &LinkNode{Data: listData[0]}
	for _, v := range listData[1:] {
		node := &LinkNode{Data: v}
		node.NextNode = head //新节点的next指向原来的头节点head【头部插入】
		head = node          //新的头节点head换成node【头部插入】
	}
	return head
}

//创建链表，尾插法
func CreateNodeOnTail(listData []int64) *LinkNode {
	if len(listData) == 0 {
		return &LinkNode{Data: 0}
	}
	head := &LinkNode{Data: listData[0]} //链表头
	end := head                          //end记录链表的尾，不断向end后插入新节点
	for _, v := range listData[1:] {
		end.NextNode = &LinkNode{Data: v}
		end = end.NextNode //移动end指针，指向新的尾节点
	}
	return head
}

//遍历链表
func (head *LinkNode) Traverse() {
	point := head
	for point != nil {
		fmt.Println(point.Data)
		point = point.NextNode //移动point到链表下一个元素
	}
	fmt.Println("----------------Done--------------")
}

//迭代反转链表
//func (head *LinkNode) ReverseIter() {
func (head *LinkNode) ReverseIter() *LinkNode {
	if head == nil || head.NextNode == nil {
		return head
	}

	var prev *LinkNode = nil          //前一个节点
	var cur *LinkNode = head          //当前节点
	var end *LinkNode = head.NextNode //后一个节点

	for {
		cur.NextNode = prev //当前节点指向前一个节点
		//fmt.Println("a", cur, "prev", prev)
		if end == nil {
			break
		}
		//从prev开始，指针依次后移
		prev = cur
		cur = end
		end = end.NextNode
		//fmt.Println(cur)
	}
	fmt.Println("chenli", head)
	fmt.Println("chenli", cur)
	return cur
}

//网上的一个迭代反转的例子
func ReverseNode1(head *LinkNode) {
	if head == nil || head.NextNode == nil {
		return
	}
	var prev *LinkNode = nil
	var cur *LinkNode = head
	var end *LinkNode = head.NextNode
	for {
		cur.NextNode = prev //当前结点指向前一个节点
		if end == nil {
			break
		}
		prev = cur         //指针后移
		cur = end          //指针后移，处理下一个节点
		end = end.NextNode //temp作为中间节点，记录当前结点的下一个节点的位置
	}
	head = cur
	fmt.Println(head)
}

//递归反转的例子
func ReverseNode2(head *LinkNode) *LinkNode {
	if head == nil || head.NextNode == nil {
		return head
	} else {
		//一直递归，找到链表中最后一个节点
		new_head := ReverseNode2(head.NextNode)

		head.NextNode.NextNode = head
		head.NextNode = nil
		return new_head
	}
}

//头插法反转的例子
func ReverseNode3(head *LinkNode) *LinkNode {
	var new_head *LinkNode = nil
	var temp *LinkNode = nil

	if head == nil || head.NextNode == nil {
		return head
	}
	for head != nil {
		temp = head
		head = head.NextNode
		temp.NextNode = new_head
		new_head = temp
	}
	return new_head
}

//就地逆置反转

//使用快、慢指针的思路来判断链表中是否有环，并返回快慢指针相遇点
func isRingLinkNode(head *LinkNode) (is bool, meet *LinkNode) {
	slow := head
	fast := head
	for fast != nil && fast.NextNode != nil {
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

//使用哈希表判断链表是否有环,并返回环入口
func isRingLinkNodeHash(head *LinkNode) (is bool, meet *LinkNode) {
	seen := map[*LinkNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true, head
		}
		seen[head] = struct{}{}
		head = head.NextNode
	}
	return false, nil
}

//判断环的入口处，https://leetcode-cn.com/problems/c32eOV/solution/lian-biao-zhong-huan-de-ru-kou-jie-dian-vvofe/
//https://leetcode-cn.com/problems/linked-list-cycle-lcci/solution/kuai-man-zhi-zhen-wei-shi-yao-yi-ding-za-s8ts/  快慢指针第一圈相遇
func isRingLinkNode1(head *LinkNode) *LinkNode {
	slow, fast := head, head
	for fast != nil && fast.NextNode != nil {
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
		if fast == slow {
			p := head
			for p != slow {
				p = p.NextNode
				slow = slow.NextNode
			}
			return p
		}
	}
	return nil
}

//插入、查询

func main() {
	//单向链表的例子
	node1 := new(LinkNode)
	// node1 为struct， new初始struct的零值为&{0 <nil>}
	fmt.Println(node1)
	node1.Data = 1

	node2 := new(LinkNode)
	node2.Data = 2
	node1.NextNode = node2 // node2链接到node1上，new返回地址

	node3 := new(LinkNode)
	node3.Data = 3
	node2.NextNode = node3

	//遍历链表
	node1.Traverse()

	//从slice创建链表并遍历一下
	s1 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	link1 := CreateNodeOnHead(s1)
	link2 := CreateNodeOnTail(s1)
	fmt.Println("slice is ", s1)
	fmt.Println("link1 is ")
	link1.Traverse()
	fmt.Println("link2 is ")
	link2.Traverse()

	//此处引入了link3，之前理解ReverseIter可以原地修改link1，测试后发现不能，原因考虑是函数或方法中形参head记录的是链表中
	//第一个元素的地址，后续如果打印head，则head的next反转后已经指向了nil，因此仅仅只能打印出一个元素
	//link3 := link1.ReverseIter()
	fmt.Println("reverse links1 is ", link1)
	//link1.Traverse()
	//link1.ReverseIter()
	//link1.Traverse()
	link3 := link1.ReverseIter()
	link3.Traverse()
	ReverseNode1(link3)
	link3.Traverse()

	fmt.Println("递归反转")
	link4 := ReverseNode2(link2)
	link4.Traverse()

}
