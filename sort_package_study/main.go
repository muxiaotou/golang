//+build ingore
package main

//https://learnku.com/articles/38269
//https://www.cxyzjd.com/article/weixin_40580582/93900141

//sort包提供了三个可以直接使用的函数Ints(a []int)\Float64s(a []float64)\Strings(a []string),对三种类型的slice进行升序排序
//如果要降序排序，则需要sort.Sort(sort.Reverse(sort.IntSlice(a)))

//sort.Slice(s, func(i, j int) bool{return s[i] > s[j]}) 使用自定义less对slice进行排序
import (
	"fmt"
	"sort"
	"strings"
)

type Student struct {
	Age  int
	Name string
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

	//对字符串排序，两种方法：1）将string转成[]string,排序后再拼接成string；2）构造less方法对slice进行排序进行排序
	s2 := "badc"
	tmp := strings.Split(s2, "")
	fmt.Println(tmp)
	sort.Strings(tmp)
	s2 = strings.Join(tmp, "")
	fmt.Println("s2 sorted is  ", s2)

	s3 := "agih"
	s3tmp := []byte(s3)
	sort.Slice(s3tmp, func(i, j int) bool { //Slice函数，会按照用户自定义的less方法，对s3tmp slice进行排序
		//func(i, j int) bool 为自定义的less函数，下面的return是具体的实现放，i、j int类型代表index，而非s3tmp的slice元素
		//s3tmp必须为slice
		return s3tmp[i] < s3tmp[j]
	})

	s3 = string(s3tmp)
	fmt.Println("s3 sorted is ", s3)

	//sort.Slice是使用自定义的less函数对slice进行排序，比如可以实现逆序排序
	s4 := []int{1, 6, 7, 4}
	sort.Ints(s4) //默认升序
	fmt.Println(s4)

	s5 := []int{1, 6, 7, 4}
	sort.Slice(s5, func(i, j int) bool { //通过自定义的less函数，实现降序排序
		return s5[i] > s5[j]
	})
	fmt.Println(s5)

	//对map排序
	//对struct排序
	//上面通过sort.Slice排序的是slice，也可以排序map、struct组成的slice
	//排序顺序在于自定义的less函数
	//func(i, j int) bool { //通过自定义的less函数，实现降序排序
	//		return s5[i] > s5[j]
	//	}里面的return后面怎么来写，比如s5[i].age > s5[j].age

	//sort.Sort排序任意实现sort.Interface的数据结构(就是实现了Less\Swap\Len三个方法)
	//比如实现一个struct slice的排序
	//type Student struct {
	//	Age  int
	//	Name string
	//}
	//
	//type Students []Student
	//
	//func (s Students) Len() int {return len(s)}
	//func (s Students) Less(i,j int) bool {return s[i].Age < s[j].Age}
	//func (s Students) Swap(i,j int) {s[i],s[j]=s[j],s[i]}

	stu := []Student{
		{10, "beibei"},
		{5, "juanjuan"},
		{20, "mutou"},
	}

	sort.Sort(Students(stu))
	fmt.Println(stu)

}
