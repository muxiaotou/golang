//+build ignore
package main

import "fmt"

//本题想到两种解法，一种是暴力解法，一种是二分法
//暴力解法：即从1开始遍历到(num/2)+1,当满足x*x <= num <（x+1）*（x+1）时，x即为平方根
//二分法：初始左标是0，右标是(num/2)+1,mid是(右标-左标)/2,判断mid*mid与num的大小关系，从而挪动左标 ，右标，不用像暴力解法一样挨个遍历，理论可以减少近一半的耗时
//上面(num/2)+1是为了防止当num=1时，导致mid计算成0而出现错误

//以上方法考虑了num为0、1这两种特殊的情况，其中为了防止溢出，可以将i*i<=num写成i<num/i
func sqrtx1(num int) int {
	mid := (num / 2) + 1        //为了防止num=1时计算结果为0
	for i := 1; i <= mid; i++ { //从1开始判断，即使num=0，后面的if判断是不满足的，最后会走return 0
		//if i*i <= num && (i+1)*(i+1) > num { //这句是最核心的，依次遍历判断是否满足平方的条件
		if i <= num/i && (i+1) > num/(i+1) { //这句是最核心的，依次遍历判断是否满足平方的条件
			return i
		}
	}
	return 0 //当上面的num=0时，会走此语句返回
}

//此方法可以降低一定的计算量，但是有几处需要仔细
func sqrtx2(num int) int {
	l := 1
	r := (num / 2) + 1
	for l <= r {
		mid := (r-l)/2 + l
		if mid*mid == num {
			return mid
		} else if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

//考虑num=0、1、8时，推理

func main() {
	fmt.Println(sqrtx1(0))
	fmt.Println(sqrtx1(1))
	fmt.Println(sqrtx1(8))

	fmt.Println(sqrtx2(0))
	fmt.Println(sqrtx2(1))
	fmt.Println(sqrtx2(8))
}
