package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)
	}

	if WaitTimeout(&wg, time.Second*5) {
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(w *sync.WaitGroup, duration time.Duration) bool {
	ch := make(chan bool)
	select {
	case <-time.After(duration):
		go func() {
			ch <- true
		}()
	default:
		w.Wait()
		go func() {
			ch <- false
		}()
	}
	return <-ch
}
