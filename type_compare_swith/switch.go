package main

import (
	"fmt"
	"reflect"
)

/*
	类型转换是用来在类型不同但相互兼容的类型之间的相互转换的方式，如果不兼容，则无法相互转换，写法是a(b)，把b转换成a
	类型断言是在接口之间进行，本质也是类型转换，写法是a.(b)，含义是把a转换成b，有两种方式：
	1）s, ok := x.(T)，ok为true，则转换成功
	2）switch i:= s.(type) { case int: ... case string: ... default:... }
	https://cloud.tencent.com/developer/article/1581011
	https://learnku.com/articles/42797
*/

func main() {

	//reflect.Typeof()  返回变量的类型对象
	var s string = "AB"
	fmt.Println(reflect.TypeOf(s[0])) //uint8
	fmt.Printf("%T\n", s[0])
	for _, v := range s {
		fmt.Println(reflect.TypeOf(v)) //int32
		fmt.Printf("%T\n", v)          //int32
	}

	/*
		类型转换的一些例子
	*/

	//类型不同，也不兼容，所以编译会报错
	//s1 := "ab"
	//i := int(s1)  cannot convert s1 (type string) to type int

	//类型不同，但兼容，所以可以转换
	var j int8 = 1
	m := int(j)
	fmt.Printf("%v %T\n", m, m)

	//运行时检测到k的类似是string，与int不兼容，panic
	var k interface{} = "s"
	//l := k.(int) panic: interface conversion: interface {} is string, not int
	//fmt.Printf("%v %s", l, l)
	l, f := k.(int)               //增加参数来加以判断，是由当f为true时，才会转换成功
	fmt.Printf("%v, %v \n", l, f) //0  false
	p, f := k.(string)
	fmt.Printf("%v %v\n", p, f) //s  true

	//string和[]byte可以相互进行互转
	//string和数字可以使用strconv.Atoi\strconv.Itoa相互转换

	//switch 断言
	var i interface{} = 1398
	switch t := i.(type) {
	case int:
		fmt.Printf("x is int, values is %v\n", t)
	case string:
		fmt.Printf("x is string, values is %x\n", t)
	default:
		fmt.Printf("don`t know x type\n")
	}

}
