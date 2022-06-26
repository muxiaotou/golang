//https://geekr.dev/posts/go-anonymous-function-and-closure【仅仅看懂了文章中的一到三(三当中的部分)内容，剩余的函数式编程篇
//博大精深，暂时看不懂，先留着以后看吧】
package main

import (
	"fmt"
	"sync"
	"time"
)

//import "fmt"

func main() {

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

	//匿名函数赋值给变量
	//var add = func(a, b int) int {
	//	return a + b
	//}
	//fmt.Println(add(1, 2))
	//

	//定义匿名函数时直接调用匿名函数
	//func(a, b int) {
	//	fmt.Println(a + b)
	//}(1, 3)

	//	匿名函数一，最终结果  vol4  vol4  vol4 vol4
	keys := []string{"vol1", "vol2", "vol3", "vol4"}
	wg1 := sync.WaitGroup{}
	wg1.Add(len(keys))
	for _, k := range keys {
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println(k) //k使用的函数外全局变量，最终执行的时候k已经变成了vol4,
			wg1.Done()
		}()
	}
	wg1.Wait()

	//	匿名函数二，最终结果  vol1  vol2  vol3  vl4（顺序每次不一定一致）
	wg2 := sync.WaitGroup{}
	wg2.Add(len(keys))
	for _, k := range keys {
		go func(k string) {
			time.Sleep(5 * time.Second)
			fmt.Println(k) //k使用的是函数内局部变量，每次协程起匿名函数的时候即已经赋好了值
			wg2.Done()
		}(k)
	}
	wg2.Wait()
}
