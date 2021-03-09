package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x + y)
			fmt.Println("done")
			return
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	fmt.Println("begin")
	c := make(chan int, 3)
	quit := make(chan int)
	go func() {
		fmt.Println("get channel")
		<-c
		quit <- 5
	}()
	go fibonacci(c, quit)
	time.Sleep(time.Second * 5)
}
