package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	c <- 1
	go func() {
		for s := range c {
			fmt.Println("channel data is ", s)
		}
		fmt.Println("i exit")
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("close channel")
	close(c)
	time.Sleep(1 * time.Second)
}
