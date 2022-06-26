package main

import "fmt"

//https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html

/*
	1.defer用于资源的释放，会在函数返回之前进行调用，比如defer f.close()
	2.多个defer，调用顺序类似于栈，越后面的defer表达式先被执行
	3.函数的返回过程：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中，return语句并不是一条原子指令，因此
	可能在设置函数返回值之后，在返回到调用函数之前，修改返回值。defer可以改写成
		返回值=xxx
		调用defer函数
		空的return
*/

//例1
func a() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//例1改写成，先设置return的值0给result，然后执行defer当中的函数，然后返回给调用函数，此时defer当中函数修改了result的值
func a1() (result int) {
	result = 0
	func() {
		result++
	}()
	return
}

//例2
func b() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//例2改写成，先设置r=5，然后执行defer内函数修改t的值，但并未修改r的值，return返回调用函数时r的值仍为5
func b1() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
	return
}

//例3
func c() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//例3改写成，先设置r=1，然后执行defer内函数，由于外部defer内函数参数为值传递，函数内局部变量修改不影响外部变量值，因此最终返回调用函数值为1
func c1() (r int) {
	r = 1
	func(r int) {
		r = r + 5
	}(r)
	return
}

func main() {
	fmt.Println(a())  // 1
	fmt.Println(a1()) // 1

	fmt.Println(b())  //5
	fmt.Println(b1()) //5

	fmt.Println(c())  //1
	fmt.Println(c1()) //1

	//类似栈，先入后出，打印输出为 4  3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //此处i实际是外部变量i
	}

	//类似栈，先入后出，但是由于匿名函数执行时，i已经变成了5，因此打印输出 5 5 5 5 5
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i) //此处i实际是外部变量i
		}()
	}

	//类似栈，先入后出，由于匿名函数每次已经对内部变量进行了赋值，因此打印输出 4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i) //此处i实际是函数内的局部变量i
		}(i)
	}
}
