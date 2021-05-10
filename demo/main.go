//refer to : https://www.bookstack.cn/read/hunterhug-goa.c/golang-basic.md
package main

import (
	"beibei.com/golang/demo/diy"
	"fmt"
)

//init 函数先于main函数执行，无参数和返回值
//同一文件内多个init按照声明顺序先后执行
//同一个包内多个init按照文件名字典序依次执行(网上看的，不知道对错)
//不同包的init函数按照包导入的依赖关系决定执行顺序
func init() {
	//声明并且初始化三个值
	var i, j, k = 1, 2, 3
	//println打印字符串和变量，打印字符串且行尾加换行符\n
	//print类似println打印字符串和变量，但行尾不加换行符\n
	//printf格式化打印字符串，行尾需要自行加上换行符\n, %T type,%v value in default format, %c character, %d base 10
	//sprintf 返回格式化后的字符串，
	fmt.Println("init hello world")
	fmt.Println(i, j, k)
}

//两个数相加
func sum(a, b int64) int64 {
	return a + b
}

//main函数无参数和返回值
func main() {
	//未使用的变量，不允许声明，否则语法报错
	//noused := 8

	p := true                             //bool
	a := 3                                //int
	b := 6.0                              //float64
	c := "ni"                             // string
	d := [2]string{"1", "2"}              //array
	e := []int64{1, 2, 3}                 //slice
	f := map[string]int64{"a": 3, "b": 4} //map

	fmt.Printf("type: %T, value:%v\n", p, p)
	fmt.Printf("type: %T, value:%v\n", a, a)
	fmt.Printf("type: %T, value:%v\n", b, b)
	fmt.Printf("type: %T, value:%v\n", c, c)
	fmt.Printf("type: %T, value:%v\n", d, d)
	fmt.Printf("type: %T, value:%v\n", e, e)
	fmt.Printf("type: %T, value:%v\n", f, f)

	//slice 赋值
	e[0] = 5
	//	slice增加值
	e = append(e, 3)
	fmt.Println(e)

	//增加map值
	f["c"] = 5
	//查找map键值
	v, ok := f["c"]
	fmt.Println(v, ok)
	v, ok = f["ff"]
	fmt.Println(v, ok)

	//	判断语句
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		fmt.Println("a <= 0")
	}

	//死循环
	a = 0
	for {
		if a >= 10 {
			fmt.Println("out")
			//退出循环
			break
		}

		a = a + 1 //a++
		if a > 5 {
			continue
		} else {
			fmt.Println(a)
		}
	}

	//	循环语句
	for i := 9; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}

	//	循环slice
	for k, v := range e {
		fmt.Println(k, v)
	}

	//	循环map
	for k, v := range f {
		fmt.Println(k, v)
	}

	//定义int64变量
	var h, i int64 = 4, 6

	//	使用函数
	s := sum(h, i)
	fmt.Printf("sum(%v + %v)=%v\n", h, i, s)

	//	新建结构体，值
	g := diy.Diy{A: 3}

	fmt.Printf("type:%T, value:%v\n", g, g)

	//结构体值变化
	g.Set(1, 1)
	fmt.Printf("type:%T, value:%v\n", g, g)

	//结构体值未变化
	g.Set2(3, 3)
	fmt.Printf("type:%T, value:%v\n", g, g)

	//	新建结构体，引用
	m := new(diy.Diy) //返回*diy.Diy
	//fmt.Printf("%T", m)
	m.A = 2
	fmt.Printf("type:%T, value:%v\n", m, m)

	s1 := make([]int64, 5)
	s2 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v, len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%v,cap:%v, len:%v\n", s1, cap(s1), len(s1))

	fmt.Printf("%#v,cap:%#v, len:%#v\n", s2, cap(s2), len(s2))
	fmt.Printf("%#v, len:%#v\n", m1, len(m1))
	fmt.Printf("%#v, len:%#v\n", m2, len(m2))

	var ll []int64
	fmt.Printf("%#v \n", ll)
	ll = append(ll, 1)
	fmt.Printf("%#v \n", ll)
	ll = append(ll, 2, 3, 4, 5, 6, 7, 8) //append后面添加1个或多个元素
	fmt.Printf("%#v \n", ll)
	ll = append(ll, []int64{9, 10, 11}...) //...语法糖打散slice，等价于ll = append(ll,9, 10, 11)
	fmt.Printf("%#v \n", ll)

	fmt.Println(ll[0:2]) //包前不包后
	fmt.Println(ll[:2])  //默认从0开始
	fmt.Println(ll[0:])
	fmt.Println(ll[:])
}
