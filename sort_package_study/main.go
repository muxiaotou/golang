//+build ingore
package main

//https://learnku.com/articles/38269

//sort包提供了三个可以直接使用的函数Ints(a []int)\Float64s(a []float64)\Strings(a []string),对三种类型的slice进行升序排序
//如果要降序排序，则需要sort.Sort(sort.Reverse(sort.IntSlice(a)))
import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 1, 2, 7, 6, 5}
	sort.Ints(s)
	fmt.Println(s)

	s1 := []int{3, 1, 2, 7, 7, 9}
	sort.Sort(sort.IntSlice(s1)) //与sort.Ints(s)等价，sort.IntSlice(s1)将入参转换成IntSlice类型，sort.Ints(s)函数内部进行转换
	fmt.Println(s1)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//[]int\[]float64\[]string排序本质上是由sort.Sort包进行的
	//上一句的分解
	//sort.IntSlice(s),将[]int类型强制转换为sort包当中定义的IntSlice类型，因为后者实现了sort包当中Interface接口的所有方法
	//sort.Reverse(),重新定义了Less方法，即排序的比较规则，排序时Swap方法会按照Less定义的比较方法来进行位置挪动
	//sort.Sort(),排序时调用的方法，会按照入参的类型定义的Less、Swap方法进行排序操作
	fmt.Println(s)
}
