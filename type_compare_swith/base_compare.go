package main

import (
	"fmt"
	"reflect"
)

// interface compare start
type Person interface {
	getName() string
}

type Studenti1 struct {
	Name string
}

type Studenti2 struct {
	Name string
}

func (s Studenti1) getName() string {
	return s.Name
}

func (s Studenti2) getName() string {
	return s.Name
}

func compare(s, t Person) bool {
	return s == t
}

// interface compare end

func main() {
	/*
		基本类型比较：数值、字符串、布尔值
		1.类型完全不一样的，不能比较
		2.类型再定义(type A string)不可比较，是两种不同的类型
			类型再定义，一般用在为一个类型添加一个独有的方法使用，再定义的类型，和原类型的方法，不再有交集。
		3.类型别名(type A = string)可比较，是同一种类型
			类型别名，和原类型完全一样，原类型的方法，别名类型也可以使用。type byte = uint8, type rune = int32
	*/

	//fmt.Println("1" == 1) invalid operation: "1" == 1 (mismatched types string and int)
	type A int
	var a int = 1
	var b A = 1
	//fmt.Println(a == b)  invalid operation: a == b (mismatched types int and A)
	fmt.Println(a == int(b)) //true，类型再定义，可以强制转换类型后再比较

	type C = int
	var c C = 1
	fmt.Println(a == c) //true， 类型别名，直接进行比较

	//reflect.TypeOf获取变量的类型对象，此对象的类型为reflect.Type()
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b)) //int main.A
	//通过类型对象的成语，可以获取变量的类型名和种类(底层类型)
	//类型名指系统原生数据类型，如int、string、bool等，以及使用type关键字定义的类型
	//种类指的是对象归属的品种，在reflect包中type.go有定义，比如int、float、Array、Slice等，结果是品种、品类
	fmt.Println(reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind()) // int int
	fmt.Println(reflect.TypeOf(b).Name(), reflect.TypeOf(b).Kind()) // A int

	/*
		复合类型比较：数组、struct结构体
		1.数组的长度不同，无法比较
		2.数组逐个比较类型和值，数组的每个元素，必须是可以比较的类型，否则无法比较
		3.struct逐个比较类型和值，每个对应的成员的比较遵循基本类型变量的比较规则，如有不可比较类型时，则struct无法比较
	*/

	var d = [2]int{1, 2}
	//var e = [3]int{1, 2, 3}
	var f = [2]int{1, 2}
	//fmt.Println(d == e) invalid operation: d == e (mismatched types [2]int and [3]int)
	fmt.Println(d == f)

	type Student struct {
		Name string
		Age  int
	}

	g := Student{Name: "chenli"}
	h := Student{Name: "chenli"}
	fmt.Println(g == h) //true

	//定义struct字段相同，本质也是不同类型，无法比较
	//type Student1 struct {
	//	Name string
	//	Age  int
	//}
	//i := Student1{Name: "chenli"}
	//fmt.Println(h == i)    invalid operation: h == i (mismatched types Student and Student1)

	//当struct当中含有无法比较的成员时，整个struct均无能进行比较
	//type Student2 struct {
	//	Name string
	//	Info []string
	//}
	//
	//j := Student2{Name: "beibei"}
	//k := Student2{Name: "beibei"}
	//fmt.Println(j == k)  invalid operation: j == k (struct containing []string cannot be compared)

	/*
		引用类型比较：slice、map、channel、pointer
		1.slice类型两两之间不能比较，只能与零值nil做比较
		2.map类型凉凉之间不能比较，只能与零值nil做比较

		slice不可比较有两个原因：
		1.引用类型，比较地址没有意义
		2.切片有len、cap，比较的维度不好衡量，因此go设计的时候就不允许切片可比较

		map内部是以key为hash的结构来实现的，因此顺序是混乱的，两次遍历中读取数据的顺序是不一致的。

		引用类型变量存储的是某个变量的地址，所以channel、pointer引用类型变量的比较，判断的是这两个引用类型存储的是不是同一个变量。
		如果是同一个变量，则内存地址肯定一样，则引用类型变量想到
		如果不是同一个变量，则内存地址肯定不一样。
	*/
	type Student3 struct {
		Name string
		Age  int
		Info []string //pointer比较时，struct包含slice item也可以比较的，关注点仅仅在于对象的内存地址是否一样
	}
	l := &Student3{Name: "juanjuan"}
	m := &Student3{Name: "juanjuan"}
	fmt.Println(l == m) //false
	n := l
	fmt.Println(n == l) //true

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	ch3 := ch1
	fmt.Println(ch1 == ch2) //false
	fmt.Println(ch1 == ch3) //true

	/*
		接口类型：
		接口类型变量，包含该接口变量存储的动态值和动态值的类型两部分组成，只有当值和类型都相等，两个接口值才是相等的。
		空接口interface{}，因为类型不确定，因此是无法比较是否相等的
	*/

	//https://segmentfault.com/a/1190000039005467
	//https://www.jianshu.com/p/a982807819fa

	//基本类型接口的例子
	var i1 interface{} = 0
	var i2 interface{} = 2
	var i3 interface{} = 0
	fmt.Println(i1 == i2) //false, i1和i2 动态类型相同，但动态值不相同
	fmt.Println(i1 == i3) //true, i1和i3，动态类型和动态值均相同

	//复合类型的例子
	i4 := Studenti1{"mutou"}
	i5 := Studenti2{"mutou"}
	i6 := Studenti1{"mutou"}
	fmt.Println(compare(i4, i5)) // false, 类型不同
	fmt.Println(compare(i4, i6)) //true, 类型相同，值相同
	//接口的动态类型必须是可比较的，如果不能比较(比如slice、map)，则运行时会panic
	//var i7 interface{} = []int{1, 2}
	//var i8 interface{} = []int{1, 2}
	//fmt.Println(i7 == i8)  panic: runtime error: comparing uncomparable type []int

	/*
		函数类型：
		函数也是一种类型，但不可以进行比较
	*/

	/*
		总结：
		1）可比较：int、float等数值类型，string、bool、complex、pointer、channel、interface、array
		2）不可比较：slice、map、function
		3）符合类型中如果带有不可比较的类型，那么该类型也是不可比较的(不可比较类型具有传递性)。
		4）可以通过reflect包来比较任意类型，包括slice、map类型
	*/
}
