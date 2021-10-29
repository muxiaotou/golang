//golang 方法是值接收者还是指针接收者的问题
package main

// 函数的变长参数
//	func myfunc(numbers ...int) {
//		for _,number := range numbers {
//			fmt.Println(number)
//		}
//}

//举例一：
//	myfunc(1,2,3,4,5)
//举例二
//	slice := []int{1,2,3,4,5}
// myfunc(slice...)

//三个点：
//1）函数的变长参数...type语法糖，用于函数有多个不定参数，本质上是一个切片[]type
//2）slice...语法糖，slice被打散进行传递，用于将slice作为参数直接传递给变长参数的函数
//3）表示数组的元素个数，array:=[...]int{1,2,3,4,5}

//任意类型的变长参数(泛型) ,例如fmt包的Println函数，可以接收任意类型的变长参数，fmt.Println("aaa")
//func Println(a ...interface{}) (n int, err error) {
//}

//多返回值
//func add(a,b *int) (int, error) {
//	...
//	return *a+*b, nil
//}

//多返回值---命名返回值
//func add(a, b *int) (c int, err error) {
//	if *a < 0 || *b < 0 {
//		err = errors.New("only fei fu shu")
//		return
//	}
//	*a *= 2
//	*b *= 3
//	c = *a + *b
//	return
//}

//	匿名函数
//
var add = func(a, b int) int {
	return a + b
}
