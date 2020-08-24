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

func containsCycle(grid [][]byte) bool {
	n := len(grid)
	if n == 0 {
		return false
	}
	m := len(grid[0])
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] == false && dfsGrid(grid, i, j, &vis) == true {
				return true
			}
		}
	}
	return false
}

func dfsGrid(grid [][]byte, i int, j int, vis *[][]bool) bool {
	if (*vis)[i][j] == true {
		return false
	}
	(*vis)[i][j] = true
	count := 0
	if i+1 < len(grid) && grid[i+1][j] == grid[i][j] {
		if (*vis)[i+1][j] {
			count++
		} else {
			if dfsGrid(grid, i+1, j, vis) {
				return true
			}
		}
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == grid[i][j] {
		if (*vis)[i][j+1] {
			count++
		} else {
			if dfsGrid(grid, i, j+1, vis) {
				return true
			}
		}
	}
	if i-1 >= 0 && grid[i-1][j] == grid[i][j] {
		if (*vis)[i-1][j] {
			count++
		} else {
			if dfsGrid(grid, i-1, j, vis) {
				return true
			}
		}
	}
	if j-1 >= 0 && grid[i][j-1] == grid[i][j] {
		if (*vis)[i][j-1] {
			count++
		} else {
			if dfsGrid(grid, i, j-1, vis) {
				return true
			}
		}
	}

	if count >= 2 {
		return true
	}
	return false
}

func main() {
	//for i := 1; i < 500; i++ {
	//	fmt.Printf("%v %v\n",i,minDays(i))
	//}
	//fmt.Println(minDays(9209408))
	//{'b','a','c'},
	//{'c','a','c'},
	//{'d','d','c'},
	//{'b','c','c'}
	fmt.Println(containsCycle([][]byte{{'a', 'g', 'b', 'e', 'c', 'b', 'd', 'd', 'c', 'c', 'e', 'd', 'b', 'd', 'a', 'a', 'h', 'c', 'g', 'f'}, {'b', 'h', 'e', 'e', 'c', 'a', 'f', 'e', 'h', 'c', 'h', 'c', 'a', 'g', 'a', 'd', 'b', 'f', 'g', 'g'}, {'b', 'b', 'c', 'b', 'a', 'a', 'b', 'a', 'a', 'f', 'e', 'f', 'f', 'g', 'e', 'g', 'e', 'h', 'b', 'e'}, {'a', 'c', 'a', 'd', 'f', 'c', 'g', 'b', 'h', 'b', 'e', 'd', 'c', 'h', 'a', 'b', 'd', 'e', 'h', 'b'}, {'a', 'a', 'd', 'a', 'c', 'b', 'b', 'h', 'g', 'f', 'a', 'a', 'g', 'a', 'd', 'd', 'd', 'd', 'd', 'g'}, {'f', 'f', 'h', 'b', 'c', 'a', 'e', 'e', 'c', 'c', 'g', 'f', 'g', 'c', 'c', 'g', 'h', 'b', 'd', 'g'}, {'d', 'a', 'a', 'a', 'e', 'b', 'g', 'h', 'a', 'b', 'g', 'h', 'c', 'd', 'h', 'g', 'c', 'h', 'f', 'a'}, {'h', 'h', 'a', 'h', 'f', 'h', 'c', 'f', 'g', 'b', 'c', 'a', 'a', 'g', 'f', 'h', 'c', 'h', 'f', 'a'}, {'d', 'f', 'c', 'f', 'e', 'c', 'c', 'd', 'd', 'e', 'b', 'g', 'd', 'g', 'f', 'f', 'f', 'h', 'f', 'b'}, {'d', 'g', 'h', 'e', 'e', 'h', 'g', 'e', 'f', 'h', 'g', 'g', 'h', 'f', 'c', 'b', 'f', 'b', 'b', 'f'}, {'g', 'f', 'h', 'g', 'a', 'e', 'c', 'b', 'f', 'd', 'g', 'a', 'g', 'h', 'a', 'h', 'd', 'g', 'h', 'c'}}))
}
