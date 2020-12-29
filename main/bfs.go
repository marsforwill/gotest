package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// 模版 BFS
func slidingPuzzle(board [][]int) int {
	m := 2
	n := 3
	target := "123450"
	var start string
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			start += strconv.Itoa(board[i][j])
		}
	}
	//fmt.Println(start)
	//fmt.Println(target)
	// 六个位置的相邻遍历 方便六个位置的0与下一步的位置交换
	neighbor := [][]int{{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	q := list.New()
	// 处理过的状态存储
	visited := make(map[string]bool)
	// 双向list当作bfs队列
	q.PushBack(start)
	visited[start] = true

	step := 0
	for q.Len() > 0 {
		sz := q.Len()
		// bfs每一层的遍历
		for i := 0; i < sz; i++ {
			cur := q.Front().Value.(string)
			q.Remove(q.Front())
			if target == cur {
				return step
			}
			idx := 0
			for idx = 0; cur[idx] != '0'; idx++ {
			}
			for _, adj := range neighbor[idx] {
				newBoard := []byte(cur)
				// swap
				temp := newBoard[idx]
				newBoard[idx] = newBoard[adj]
				newBoard[adj] = temp
				if _, ok := visited[string(newBoard)]; !ok {
					visited[string(newBoard)] = true
					q.PushBack(string(newBoard))
				}
			}
		}
		step++
	}
	return -1
}

// 获取所有钥匙的最短路径 todo 感觉实在不怎么想写
func shortestPathAllKeys(grid []string) int {
	return 0
}

func main() {
	fmt.Println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
}
