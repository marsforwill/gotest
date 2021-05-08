package main

import (
	"context"
	"fmt"
	"time"
)

//多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作。
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, cancel, 1500*time.Millisecond)
	go handle2(ctx, cancel, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, cancel context.CancelFunc, duration time.Duration) {
	cancel()
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func handle2(ctx context.Context, cancel context.CancelFunc, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle2", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process2 request with", duration)
	}
}
