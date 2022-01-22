//+build ignore

//结题思路：
//1）从低位开始依次比较排序，但是此方法需要额外声明内存空间来存储，要不然腾挪原有数组就不方便了
//2）从高位开始依次比较，此方法可以不用新声明内存空间，只需要利用最终的有序数组长度为m+n的技巧来实现,就跟题目当中要求一样，第一个slice的空间必须大于或者等于m+n，要不然index溢出
package main

import "fmt"

//低位开始比较，但是需要声明新的内存空间来存储
func mergelist1(a, b []int, m, n int) []int {
	//res := make([]int, m+n)
	var res []int
	var i, j = 0, 0
	for l := 0; l < m+n; l = l + 1 { //通过m+n来控制新的slice的遍历，这个比较方便且准确
		//刚开始使用了双遍历遍历m和n，不好控制最后一个元素
		if i < m && j < n && a[i] < b[j] { //m和n还未超边界，且a元素有小于b中元素时，追加到新的slice
			res = append(res, a[i])
			i++
		} else if i < m && j < n { //m和n还未超边界，且a元素大于或等于b中元素时，追加到新的slice
			res = append(res, b[j])
			j++
		} else { //有m或n超边界了，判断一下，是哪个超了，则把另外一个剩余的元素追加到新的slice当中
			if i == m {
				res = append(res, b[j:]...)
			} else if j == n {
				res = append(res, a[i:]...)
			}
			break
		}
	}
	return res
}

//高位开始比较，不需要声明新的内存空间，但是需要确保a的容量足够可以放下b
func mergelist2(a, b []int, m, n int) {
	for l := m + n - 1; l >= 0; l-- { //m+n 最终slice的长度，依次进行尾遍历
		if m-1 >= 0 && n-1 >= 0 && a[m-1] > b[n-1] { //m和n还未超边界，且a的元素大于b，则a的元素追加到最高位的index位置
			a[l] = a[m-1]
			m--
		} else if m-1 >= 0 && n-1 >= 0 { //m和n还未超边界，且a的元素不大于b，则b的元素追加到最高位的index位置
			a[l] = b[n-1]
			n--
		} else { //如果m和n有超过边界
			if m < 0 { //仅当m超边界了，意味这a当中元素已经移动玩，b当中还有n个元素未移动，因此开始仅仅移动b
				//不用考虑n超边界，n超边界意味着n已经移动完了，那m保持原样即可，不用再移动了
				for ; n >= 0; n-- {
					a[n-1] = b[n-1] //n-1是理解的关键呀，a的长度m+n，n全用0补足了
				}
			}
			break //都走到这里了，直接break即可了
		}
	}
}

func main() {

	s1 := []int{1, 2, 4, 6, 7, 7, 8, 9}
	s2 := []int{2, 2, 5, 6, 8, 8, 10}
	fmt.Println(mergelist1(s1, s2, 8, 7))
	fmt.Println("s1 len and cap is :", len(s1), cap(s1))

	s1 = []int{5, 5, 6, 6, 7, 7, 9, 9}
	s2 = []int{1, 1, 2, 2, 3, 3, 4, 4}
	fmt.Println(mergelist1(s1, s2, 8, 7))

	s1 = []int{1, 2, 3, 3, 5, 0, 0, 0}
	s2 = []int{8, 8, 9}
	mergelist2(s1, s2, 5, 3)
	fmt.Println(s1)
}
