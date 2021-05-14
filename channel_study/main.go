//https://www.cnblogs.com/secondtonone1/p/11843392.html
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//非缓冲chan，要求一端读取，一端写入，channel的大小为零，所以读写操作一定要匹配
	//如若不使用子goroutine，无论第11行还是14行先执行谁，主goroutine都会阻塞(golang会panic，提示dead lock)；所以此处应该先启动一个子goroutine读(或写)，再在主goroutine当中写(或读)
	nochan := make(chan int)
	go func(ch chan int) {
		data := <-ch
		fmt.Println("recive data", data)
	}(nochan)
	nochan <- 1
	fmt.Println("send data", 1)

	go func(ch chan int) {
		ch <- 2
		fmt.Println("send data", 2)
	}(nochan)
	data := <-nochan
	fmt.Println("recive data", data)

	//WaitGroup,内部维护一个计数器，Add() +1，Done() -1, Wait() 阻塞代码，直到计数器的值减为0
	nochan1 := make(chan int)
	waiter1 := &sync.WaitGroup{} //sync.WaitGroup是值类型，切记需要传递地址
	waiter1.Add(2)

	go func(ch chan int, wt *sync.WaitGroup) {
		data := <-ch
		fmt.Println("receive data ", data)
		wt.Done()
	}(nochan1, waiter1)

	go func(ch chan int, wt *sync.WaitGroup) {
		ch <- 3
		fmt.Println("send data ", 3)
		wt.Done()
	}(nochan1, waiter1)

	waiter1.Wait()

	//for range从channel中自动读取，当有channel被close后，for 循环退出
	nochan2 := make(chan int, 3)
	go func(ch chan int) {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("send data is ", i)
		}
		//	需要关闭channel，否则主协程当中的for循环无法退出, dead lock, 程序panic
		close(ch)
		fmt.Println("goroutine1 exited")
	}(nochan2)

	for data := range nochan2 {
		fmt.Println("receive data is ", data)
	}

	fmt.Println("main goroutine exited")

	//缓冲channel内部其实是一个加锁的队列，先进先出(FIFO  First Input First Output)，先写入的数据优先读出来
	nochan3 := make(chan int, 3)
	go func(ch chan int) {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("send nochan3 data is ", i)
		}
	}(nochan3)

	//for i := 0; i < 2; i++ {, chan队列里面的数据没有被接受完也可以正常退出
	for i := 0; i < 3; i++ {
		data := <-nochan3
		fmt.Println("receive nochan3 data is ", data)
	}

	//从关闭的channel中读取数据， 优先读出没有被取出的数据，然后读出存储类型的空值
	//循环读取关闭的channel不会阻塞，会一直读取空值，可以通过读取结果的bool值判断是否关闭
	nochan4 := make(chan int)
	go func(ch chan int) {
		ch <- 100
		fmt.Println("send nochan4 data is ", 100)
		close(ch)
		fmt.Println("close channel nochan4, and goroutine exited")
	}(nochan4)

	for i := 0; i < 3; i++ {
		data := <-nochan4
		fmt.Println("receive nochan4 data is ", data)
	}

	data, ok := <-nochan4
	if !ok {
		fmt.Println("channe nochan4 close")
		fmt.Println("receive close channel nochan4 data is ", data)
	}
	fmt.Println("main existed")
	//重复关闭channel会导致panic
	nochan5 := make(chan int)
	go func(ch chan int) {
		close(ch)
	}(nochan5)

	data, ok = <-nochan5
	if !ok {
		fmt.Println("nochan5 close")
	}
	//close(nochan5), 二次close已经被close的channel会panic，提示close of closed channel
	//nochan5 <- 100, 想被close的channel里再次发生数据会panic，提示send on closed channel
	fmt.Println("main exited")

	//通过for、select、channel实现超时机制
	nochan6 := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select { //select 语句阻塞，等待最先返回数据的channel
			case num := <-nochan6: //如果nochan6当中有数据，则执行fmt语句
				fmt.Println("recevie num is ", num)
			case <-time.After(3 * time.Second): //如果nochan6中3秒仍读取不到数据，则走timeout逻辑
				fmt.Println("timeout...")
				quit <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		nochan6 <- i
		time.Sleep(time.Second)
	}

	<-quit //暂时阻塞，直到可读
	fmt.Println("main exist")
}
