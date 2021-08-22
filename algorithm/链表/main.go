package main

import "fmt"

// 一个最简单的单向链表
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	//单向链表的例子
	node1 := new(LinkNode)
	node1.Data = 1

	node2 := new(LinkNode)
	node2.Data = 2
	node1.NextNode = node2 // node2链接到node1上，new返回地址

	node3 := new(LinkNode)
	node3.Data = 3
	node2.NextNode = node3

	//顺序打印数据
	nowNode := node1
	// 如果下一个节点为空，表示链表已经结束了
	for nowNode != nil {
		fmt.Println(nowNode.Data)
		nowNode = nowNode.NextNode
	}
}
