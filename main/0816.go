package main

import (
	"container/list"
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
	for i := 2; i < len(arr); i++ {
		if arr[i]%2 == 1 && arr[i-1]%2 == 1 && arr[i-2]%2 == 1 {
			return true
		}
	}
	return false
}

// 1 3 5 7 9 11
// 1 3 5
func minOperations(n int) int {
	ans := 0
	for i := 1; i < n; i += 2 {
		ans += n - i
	}
	return ans
}

func minDays(n int) int {
	if n == 1 {
		return n
	}
	l := list.New()
	l.PushBack(n)
	count := 0
	m := make(map[int]bool)

	for true {
		count++
		k := list.New()
		for l.Len() != 0 {
			num := l.Front().Value.(int)
			m[num] = true
			l.Remove(l.Front())
			_, ok := m[num-1]
			if !ok {
				k.PushBack(num - 1)
			}
			if num-1 == 1 {
				return count + 1
			}
			if num%2 == 0 {
				_, ok := m[num/2]
				if !ok {
					k.PushBack(num / 2)
				}
				if num == 2 {
					return count + 1
				}
			}
			if num%3 == 0 {
				_, ok := m[num/3]
				if !ok {
					k.PushBack(num / 3)
				}
				if num == 3 {
					return count + 1
				}
			}

		}
		l = k
	}
	return count
}

func main() {
	//for i := 1; i < 500; i++ {
	//	fmt.Printf("%v %v\n",i,minDays(i))
	//}
	fmt.Println(minDays(9209408))
}
