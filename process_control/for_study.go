package main

func main() {
	//go语言循环语言仅支持for关键字，而不支持while和do-while结构
	//和条件语句、分支语句一样，左花括号{必须与for处于同一行
	//可以使用for-range结构对可迭代集合进行遍历
	//支持基于条件判断进行循环迭代
	//允许在循环条件中定义和初始化变量，且支持多重赋值
	//for循环支持contine、break来控制循环

	//基本使用
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += 1
	//}
	//fmt.Println(sum)

	//基于条件判断进行循环
	//sum := 0
	//i := 0
	//for i < 100 {
	//	i++
	//	sum += i
	//}
	//fmt.Println(sum)

	//无限循环
	//sum := 0
	//i := 0
	//for {
	//	i++
	//	if i > 100 {
	//		break
	//	}
	//	sum += 1
	//}
	//fmt.Println(sum)

	//	多重赋值
	//a := []int{1, 2, 3, 4, 5, 6}
	//for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	//	a[i], a[j] = a[j], a[i]
	//}
	//fmt.Println(a)

	//可迭代的集合(数组、切片、字典、字符串、channel)
	//string/array/slice返回index和value
	//map返回key和value
	//channel返回元素【仅且仅能返回一个值】，当channel被close后range退出
	//for可以遍历array、slice、string和key为整形递增的map
	//for range可以完成所有for的事情，并可以遍历key为string类型的map以及遍历channel

	//同时取出索引/键对应的值
	//for k, v := range a {
	//	fmt.Println(k, v)
	//}

	//取出值，忽略索引/键
	//for _, v := range a {
	//	fmt.Println(v)
	//}

	//取出索引/键，忽略元素值
	//for k := range a {
	//	fmt.Println(k)
	//}
}
