//+build ignore
package main

import (
	"fmt"
	"sync"
	"time"
)

//go语言提供了三种解决方案来实现并发控制
//channel
//waitgroup
//context

//Sync.waitGroup方式适用于明确多个goroutine协同干一件事情
func waitcontrol() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("I am sync.waitgroup one")
		wg.Done()
	}()

	go func() {
		fmt.Println("I am sync.waitgroup two")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all group done, exit")
}

//channel+select 方式，类似通知机制，告诉某个go routine执行某些代码片段
//	select只能用于channel
//	select{} select语句中不包含case语句和default语句时，那么协程将陷入永久性阻塞
//  当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
//  当不存在可以收发的 Channel 时，执行 default 中的语句；
//	https://studygolang.com/articles/12992?fr=sidebar
//	select {} 死循环，但是前面需要再写点代码，要不然代码panic，死锁
//	for {}  死循环
func chancontrol() {
	var stop = make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("recive stop singnal, exit")
				return
			default:
				fmt.Println("please wait 2 second")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("start notify end")
	stop <- true
	fmt.Println("over, sleep 2 second")
}

// context
//https://zhuanlan.zhihu.com/p/68792989
//https://www.flysnow.org/2017/05/12/go-in-action-go-context.html
//context包提供了一个Context interface，interface内定义了4个方法：
//1）Deadline，返回一个time.Time，表示context应该结束的时间
//2）Done，当Context被取消或者超时时候返回的一个close的channel
//3）Err，context被取消的原因
//4）Value，context实现共享数据存储的地方
//context包提供了四个函数，用来创建功能各异的context：
//1）WithCancel，基于父context，创建一个可以取消的新的context
//2）WithDeadline，基于父context，创建一个具有截止时间的新的context
//3）WithTimeout，基于父context，创建一个具有超时时间的新的context
//4）WithValue，基于某个context，创建并存储对应的上下文信息

//context的例子
//https://mp.weixin.qq.com/s/6yYsWgZ0zRrYBSQyMhmGSg

func main() {
	waitcontrol()
	chancontrol()
}
