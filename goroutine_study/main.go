package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine1() {
	fmt.Println("Hello goroutine1")
}

// channel 是引用类型，此处传递值即可
func goroutine2(isDone chan bool) {
	fmt.Println("goroutine2 begin...")
	time.Sleep(2 * time.Second)
	fmt.Println("goroutine2 end...")
	isDone <- true
}

func goroutine3(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("gotoutine3 %d begin ...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("goroutine3 %d end... \n", i)
}

func main() {
	// 函数前添加go关键字，就创建一个新的goroutine
	go goroutine1()
	// 如果此处不sleep，如果main理解销毁，goroutine1可能没有机会执行
	time.Sleep(1 * time.Second)
	fmt.Println("Hello main")

	// 使用channel阻塞来控制goroutine的执行，比直接sleep优雅许多
	isDone := make(chan bool)
	go goroutine2(isDone)
	<-isDone
	close(isDone)
	fmt.Println("main goroutine end...")

	// 可以使用sync.WaitGroup同步机制来控制多个goroutine的退出
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go goroutine3(i, &wg)
	}
	wg.Wait()

	//goroutine传入外部变量，保证每次printf的正确性，如若未传入，最终由于外部变量变化，导致printf打印非预期
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s !\n", who)
		}(name)
	}
	time.Sleep(2 * time.Second)
}
