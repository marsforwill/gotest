package main

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// sync.pool 多协程安全 存放暂存对象 对于很多需要重复分配、回收内存的地方，sync.Pool 是一个很好的选择。
//频繁地分配、回收内存会给 GC 带来一定的负担，严重的时候会引起 CPU 的毛刺，而 sync.Pool 可以将暂时不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，复用对象的内存，减轻 GC 的压力，提升系统的性能
func main() {
	// 初始化一个pool
	pool := &sync.Pool{
		// 默认的返回值设置，不写这个参数，默认是nil
		New: func() interface{} {
			return 1024
		},
	}

	// 看一下初始的值，这里是返回0，如果不设置New函数，默认返回nil
	init := pool.Get()
	fmt.Println(init)

	// 设置一个参数1
	for i := 0; i < 10; i++ {
		pool.Put(i)
	}

	// 获取查看结果
	for i := 0; i < 10; i++ {
		go func() {
			num := pool.Get()
			fmt.Println(num)
		}()
	}

	time.Sleep(300000)
	// 再次获取，会发现，已经是空的了，只能返回默认的值。
	num := pool.Get()
	fmt.Println(num)
	io.WriteString()
}
