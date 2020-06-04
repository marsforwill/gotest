package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// test wait group
	// print in go routine with channel order
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		ch := make(chan int)
		go func(num int) {
			fmt.Println(num)
			wg.Done()
			ch <- i
		}(i)
		<-ch
	}
	wg.Wait()
	//time.Sleep(time.Second)
	println("done")
	// test for select
	c := make(chan bool)
	//c <- true 不能自己写入数据 需要等待其他协程来写
	go func() {
		time.Sleep(time.Second)
		c <- true
	}()
	select {
	case ch1 := <-c:
		println(ch1)
	case <-time.After(time.Second):
		println("time out")
	}
}
