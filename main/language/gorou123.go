package main

import (
	"sync"
	"time"
)

func dealImage(wg *sync.WaitGroup) {
	println("deal image ok")
	wg.Done()
}

func dealVoice(wg *sync.WaitGroup) {
	println("deal voice ok")
	wg.Done()
}

func dealText(wg *sync.WaitGroup) {
	println("deal text ok")
	wg.Done()
}

func main() {

	//10ms内返回 有多少结果返回多少
	wg := sync.WaitGroup{}
	ch := make(chan bool)
	wg.Add(3)
	go dealImage(&wg)
	go dealVoice(&wg)
	go dealText(&wg)
	go func() {
		wg.Wait()
		ch <- true
	}()
	select {
	case <-time.After(time.Second):
		println("time 10 ms out")
	case <-ch:
		println("task ok")
	}

	//// test wait group
	//// print in go routine with channel order
	//wg := sync.WaitGroup{}
	//wg.Add(10)
	//for i := 1; i <= 10; i++ {
	//	ch := make(chan int)
	//	go func(num int) {
	//		fmt.Println(num)
	//		wg.Done()
	//		ch <- i
	//	}(i)
	//	<-ch
	//}
	//wg.Wait()
	////time.Sleep(time.Second)
	//println("done")
	//// test for select
	//c := make(chan bool)
	////c <- true 不能自己写入数据 需要等待其他协程来写
	//go func() {
	//	time.Sleep(time.Second)
	//	c <- true
	//}()
	//select {
	//case ch1 := <-c:
	//	println(ch1)
	//case <-time.After(time.Second):
	//	println("time out")
	//}
}
