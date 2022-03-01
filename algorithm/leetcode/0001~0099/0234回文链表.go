//+build ignore
//想到一个解题思路，就是迭代遍历链表，将元素存储到slice当中，然后再判断slice是否是回文即刻，这种办法看着是引入了额外的存储空间

package main

import "fmt"

type link234 struct {
	Value int
	Next  *link234
}

func hwlink(l *link234) bool {
	var tmplist = []int{}
	for l != nil {
		tmplist = append(tmplist, l.Value)
		l = l.Next
	}
	lenth := len(tmplist)
	for i := 0; i <= lenth/2; i++ { //n/2，偶数是感觉多比较了一次，可以(n-1)/2 ?
		if tmplist[i] != tmplist[lenth-i-1] {
			return false
		}
	}
	return true
}

func main() {
	link11 := &link234{Value: 1}

	link12 := &link234{Value: 2}
	link11.Next = link12

	link13 := &link234{Value: 3}
	link12.Next = link13

	link14 := &link234{Value: 3}
	link13.Next = link14

	link15 := &link234{Value: 2}
	link14.Next = link15

	link16 := &link234{Value: 1}
	link15.Next = link16

	fmt.Println(hwlink(link11))

}
