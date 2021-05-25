// 空结构体
package main

import (
	"fmt"
	"time"
	"unsafe"
)

//各类型值的零值：
//false: bool
//0: integer
//0.0: float
//"": string
//nil: pointer, function, interface, slice, channel, map

var a int
var b string
var c bool
var d [3]int32
var e []string
var f = make([]string, 1)
var g map[string]bool
var h = make(map[string]bool)

type S struct {
}

type T struct {
}

func (t *T) Call() {
	fmt.Println("hello, Call")
}

type Set map[string]struct{}

func (s Set) Append(k string) {
	s[k] = struct{}{} //struct{}是一种类型，那空的结构体就是struct{}{}
	fmt.Println(s)
}

func (s Set) Remove(k string) {
	delete(s, k)
	fmt.Println(s)
}

func (s Set) Exist(k string) bool {
	_, ok := s[k]
	fmt.Println(ok)
	return ok
}

func main() {
	// 以下仅仅单纯的声明，亦然会占用宽度(空间)
	fmt.Printf("a to h lenth is %d %d %d %d %d %d %d %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d), unsafe.Sizeof(e), unsafe.Sizeof(f), unsafe.Sizeof(g), unsafe.Sizeof(h))

	//空结构体的宽度比较特殊，是0
	var s S
	fmt.Println("empty struct lenth is ", unsafe.Sizeof(s))

	// 空结构体宽度是0，不占用空间，被当做占位符来经常使用，因此会被用到的有三个场景： 1.实现方法的接收者；2.实现集合类型；3.实现空通道

	// 方法接收者
	var t T
	t.Call()

	//	集合：本质就是一个list，只是list的元素不重复，由于map的key不重复，因此可以利用这一点来实现一个set，
	//	map的value由于用不到，可以用空struct来占位，因为其他类型做为value可能会占用内存，而空struct占位符不占用内存空间，提高性能
	set := Set{}
	set.Append("chen")
	set.Append("li")
	set.Exist("chen")
	set.Remove("chen")
	set.Exist("chen")

	// 实现空通道
	ch := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	fmt.Println("wait goroutine......")
	<-ch
	fmt.Println("goroutine end...")
}
