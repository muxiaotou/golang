//+build ignore

//同一个包下面多个main函数，需要加上+build ignore
package main

import (
	"fmt"
	"sync"
)

//互斥锁： sync.Mutex类型的零值表示未被锁定的互斥量，  Lock Unlock
//对于同一个互斥锁的锁定操作和解锁操作应该成对出现，如果锁定了一个已锁定的互斥锁，那么进行重复锁定操作的goroutine将被阻塞，直到该互斥锁
//回到解锁状态。对一个未锁定的互斥锁进行解锁操作，panic
//锁不能拷贝，因为变量资源本身带状态，操作一定要配套，不然会死锁。

func main() {
	var mutex sync.Mutex //零值表示未被锁定的互斥量
	//defer func() {
	//	fmt.Println("Try to recover the panic")
	//	if p := recover(); p != nil {
	//		fmt.Printf("Recovered the panic (%#v).\n", p)
	//	}
	//}()
	fmt.Println("Lock the lock.")
	mutex.Lock()
	fmt.Println("The lock is locked.")
	fmt.Println("Unlock the lock.")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.")
	fmt.Println("Unlock the lock again.")
	//mutex.Unlock()  //unlock一个未lock的mutx，panic

}
