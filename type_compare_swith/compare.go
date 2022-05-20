package main

import (
	"fmt"
	"reflect"
)

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
	//通过类型对象的成语，可以获取变量的类型名和种类
	//类型名指系统原生数据类型，如int、string、bool等，以及使用type关键字定义的类型
	//种类指的是对象归属的品种，在reflect包中type.go有定义，比如int、float、Array、Slice等，结果是品种、品类
	fmt.Println(reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind()) // int int
	fmt.Println(reflect.TypeOf(b).Name(), reflect.TypeOf(b).Kind()) // A int

	/*
		复合类型比较：数组、struct结构体
		1.数组的长度不同，无法比较
		2.数组逐个比较类型和值，数组的每个元素，必须是可以比较的类型，否则无法比较
		3.struct逐个比较类型和值，每个对应的成语的比较遵循基本类型变量的比较规则，如有不可比较类型时，则struct无法比较
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
	*/

}
