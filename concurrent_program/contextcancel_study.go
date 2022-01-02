//+build ignore
package main

import (
	"context"
	"fmt"
	"time"
)

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

func HandelRequest(ctx context.Context) {
	go WriteRedisk(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedisk(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	ctx, cacel := context.WithCancel(context.Background())
	go HandelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all goroutine")
	cacel() //cancel取消了所有子的routine

	time.Sleep(5 * time.Second)

}
