package main

//https://geekr.dev/posts/go-type-assertion

//基本类型断言
//无需使用反射，直接使用variable.(type)表达式即可获取variable变量对应的类型值

//空接口   interface{}，任何类型都可以赋值给空接口类型的变量
//var v1 interface{} =1
//var v2 interface{} = true
