package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁由结构体类型sync.RWMutex表示，零值即表示可用的读写锁实例了。
//读写锁控制下的多个写操作之间是互斥的，写操作和读操作之间是互斥的，多个读操作之间却不存在互斥关系。
//对一个未被写锁定的读写锁写解锁，对一个未被读锁定的读写锁读解锁，panic
//读锁的作用？
//Lock 写加锁  Unlock写解锁
//RLock 读加锁  RUnlock 读解锁

func main() {

	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock() //read lock
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock() //read unlock
			fmt.Printf("Unlock for reading. [%d]\n", i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writeing...")
	rwm.Lock() //write lock,虽然未解锁，但是程序退出了，因此无需再解锁了
	fmt.Println("Locked for writing.")
}
