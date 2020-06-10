package main

import (
	"sync"
	"time"
)

type Stu struct {
	lock sync.RWMutex
	num  int
}

func write(num int, s *Stu) {
	time.Sleep(time.Second)
	s.num = num
}

func read(s *Stu) {
	s.lock.Lock()
	defer s.lock.Unlock()
	time.Sleep(time.Second / 5)
	println(s.num)

}

func main() {
	s := new(Stu)
	for i := 0; i < 10; i++ {
		go write(i, s)
	}
	for i := 0; i < 10; i++ {
		go read(s)
	}
	time.Sleep(time.Second * 6)
}
