//+build ignore

//解题思路：已知是一个有序的序列，因此想到两种办法，暴力破解和二分查找法

package main

import "fmt"

//暴力破解法，叫BF
func searchindex1(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums) //如果上面for循环都比较完了，还没有找到比target大的，那就target排到nums最后的位置
}

//二分查找法，即一分为二的查找，核心是找到mid，然后mid与target进行比较
//核心找到nums[left] < nums[mid] <= nums[right]
func searchindex2(nums []int, target int) (index int) {
	left, right := 0, len(nums)-1
	index = len(nums)
	for left <= right { //此处的left==right是非常关键的！！！
		mid := (right-left)>>1 + left
		if nums[mid] >= target { //如果mid大于或等于target目标值，则right移到mid-1位置
			index = mid
			right = mid - 1
		} else if nums[mid] < target { //如果mid小于target目标值，则left移到mid+1位置
			left = mid + 1
		}
	}
	return
}

func main() {
	nums := []int{1, 3, 5, 7}
	fmt.Println(searchindex1(nums, 0))

	fmt.Println(searchindex2(nums, 6))
}
