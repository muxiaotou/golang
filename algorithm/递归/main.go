//斐波那契数列  阶乘  二分查找
package main

import "fmt"

//递归函数必须设置一个锚点来触发递归函数结束并返回结果，否则递归函数会进入无穷无尽的循环当中，最终引发内存溢出

func main() {
	//有一对兔子，从出生后第3个月起(第2个月后)每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问5个月的兔子总对数为多少？
	//手动试算每个月的兔子对数：1,1,2,3,5,8......斐波那契数列呀，F(1)=1, F(2)=1, F(3)=2, F(n)=F(n-1)+F(n-2)当n>2
	var n int = 5
	fmt.Println("rabbit1 count is ", rabbit_born1(n))
	fmt.Println("rabbit2 count is ", rabbit_born2(n, 1, 1))

	//阶乘：N!=N*(N-1)*(N-2)......
	fmt.Println("factorial1 value is ", factorial1(n))
	fmt.Println("factorial2 value is ", factorial2(n, 1))

	array := []int{1, 3, 7, 9, 16, 21, 110, 243, 999}
	target := 21
	result := binarysearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
	target = 1
	result = binarysearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target1 := 22
	result1 := binarysearch(array, target1, 0, len(array)-1)
	fmt.Println(target1, result1)
	target1 = 999
	result1 = binarysearch(array, target1, 0, len(array)-1)
	fmt.Println(target1, result1)

}

func rabbit_born1(n int) int {
	if n < 3 { //n<3, 触发递归结束的锚点
		return 1
	}
	return rabbit_born1(n-1) + rabbit_born1(n-2)
}

func rabbit_born2(n int, acc1 int, acc2 int) int { //相比普通递归，尾递归它增加两个累加器acc1和acc2，并给出初始的值，即n为1，2 时的结果
	if n == 1 {
		return acc1
	}
	return rabbit_born2(n-1, acc2, acc1+acc2)
}

func factorial1(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial1(n-1)
}

func factorial2(n int, acc int) int {
	if n == 1 {
		return acc
	}
	return factorial2(n-1, acc*n)
}

//二分查找递归方法
func binarysearch(array []int, target int, l, r int) int {
	if l > r {
		//l -> left, r -> rigth, 如果l > r, 则越界了
		return -1
	}

	// 从中间开始找
	mid := (l + r) / 2 //数字常量除法，除了移位运算，如果二元运算的操作数是不同种类的无类型常量，运算结果使用以下种类靠后的一个:整数、unicode字符、浮点数、复数，示例中除法运算两个常量都是整数常量，运算结果因此也是一个整数常量

	middleNum := array[mid]
	if middleNum == target {
		//找到了
		return mid
	} else if middleNum < target {
		//中间数字比目标数字小，从右边找
		return binarysearch(array, target, mid+1, r)
	} else {
		//中间数字比目标数字大，从左边找
		return binarysearch(array, target, l, mid-1)
	}
}

//二分查找非递归方法
func binarysearch2(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r

	for {
		if ltemp > rtemp {
			return -1 //越界了
		}

		//从中间开始找
		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]

		if middleNum == target {
			return mid //找到了
		} else if middleNum > target {
			rtemp = mid - 1 //中间的数比目标数还大，从左边找
		} else {
			ltemp = mid + 1 //中间的数比目标数还小，从右边找
		}
	}

}
