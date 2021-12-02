//+build ignore

package main

import (
	"fmt"
	"sort"
)

//整体思路就是排序+双指针
//开始看网上的题解，没有理解，误以为：有一种情况就直接给忽略了，比如源列表是[1，-1,1,1,2，-2]，按理应该有两组[1,-1,1]和[1,2,-2],但是题解都
//做了类似去重的操作，会导致仅有一组符合结果的列表输出
//但是实际不是，整体思想是：固定左边界left，然后取left+1和len(s)-1分别为mid和right,left+mid+right三数求和
//如果和小于0，则mid+1，如果大于0，则right-1，获得结果后，left不变，判断一下当前的mid和右移一位的mid是否相同，相同则继续+1
//right一样，判断一下是否和左移一位的right是否一样,之后mid++,right--,退到left的循环里面继续下一个循环

func threeadd(num []int) (res [][]int) {
	sort.Ints(num)                             //升序排序
	for left := 0; left < len(num)-2; left++ { //此处left小于len(num)-2,因为内层的mid和right会取值到len(num)-2和len(num)-1
		if left > 0 && num[left] == num[left-1] { //循环left时，判断left的相邻值是否相同
			continue
		}
		mid := left + 1
		right := len(num) - 1
		for mid < right {
			if num[left]+num[mid]+num[right] > 0 { //三数之和大于0，则right左移
				right--
			} else if num[left]+num[mid]+num[right] < 0 { //三数之和小于0，则left右移
				mid++
			} else { //等于0时，append到结果，并且对left和right继续移动
				res = append(res, []int{num[left], num[mid], num[right]})
				for mid < right && num[mid] == num[mid+1] {
					mid++
				}
				for mid < right && num[right] == num[right-1] {
					right--
				}
				mid++
				right--
			}
		} //当完成此for循环，即固定了left位置，遍历了left+1至len(s)-1位置里面所有符合的列表项
		//	因为整个slice已经排序，外层的left会去重，因此每个符合的slice的首元素都不会重复
	}
	return
}

func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeadd(num))
}
