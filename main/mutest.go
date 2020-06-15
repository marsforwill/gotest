package main

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

//只有当只存在读锁的时候不会互相阻塞
func main() {
	data = map[string]string{"hoge": "fuga"}
	mu = sync.RWMutex{}
	go read()
	go read()
	go read()
	go read()
	go read()
	go read()
	go write()
	go read()
	go read()
	go read()
	go read()
	time.Sleep(5 * time.Second)
}

// 读方法
func read() {
	println("read_wait")
	mu.RLock()
	println("read_start")
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	println("read_complete", data["hoge"])
}

// 写方法
func write() {
	//仔细看下这两行代码,此处是注释掉的
	println("write_wait")
	mu.Lock()
	println("write_start")
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data["hoge"] = "piyo"
	println("write_complete")
}
