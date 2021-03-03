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

// 获取所有钥匙的最短路径 todo 感觉实在不怎么想写 [a..f]key [A..F]lock
func shortestPathAllKeys(grid []string) int {
	// init
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	lx := len(grid)
	ly := len(grid[0])
	var sx, sy int
	keys := 0
	// vis[x][y][钥匙状态] 标记访问
	vis := make([][][]bool, len(grid))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([][]bool, 32)
		for j := 0; j < len(vis[i]); j++ {
			vis[i][j] = make([]bool, 64)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j, ch := range grid[i] {
			if ch == '@' {
				sx = i
				sy = j
			} else if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				keys = keys | (1 << (grid[i][j] - 'a'))
			}
		}
	}
	q := list.New()
	// input first head
	q.PushBack((sy << 16) | (sx << 8))
	vis[sx][sy][0] = true

	// bfs
	res := 0
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for q.Len() > 0 {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			y, x, sta := cur>>16, (cur>>8)&0xFF, cur&0xFF
			if sta == keys {
				return res
			}
			for idx := 0; idx < 4; idx++ {
				nx := x + dx[idx]
				ny := y + dy[idx]
				newSta := sta
				if nx >= 0 && nx < lx && ny >= 0 && ny < ly && grid[nx][ny] != '#' {
					flag, canThroughLock := newSta&(1<<(grid[nx][ny]-'A')), false
					if flag != 0 {
						canThroughLock = true
					}
					if grid[nx][ny] >= 'A' && grid[nx][ny] <= 'F' && !canThroughLock {
						continue
					}
					if grid[nx][ny] >= 'a' && grid[nx][ny] <= 'f' {
						newSta = newSta | (1 << (grid[nx][ny] - 'a'))
					}
					if vis[nx][ny][newSta] {
						continue
					}
					q.PushBack((ny << 16) | (nx << 8) | newSta)
					vis[nx][ny][newSta] = true
				}
			}
		}
		res++
	}

	return -1
}

func bfs(a int, b int, n int) int {
	startx, starty := 0, 0
	delx := []int{a, b, -a, -b, a, b, -a, -b}
	dely := []int{b, a, b, a, -b, -a, -b, -a}
	var queue []int
	m := make(map[int]bool)
	// count*10000+x*100+y
	queue = append(queue, startx*100+starty)
	for len(queue) > 0 {
		info := queue[0]
		queue = queue[1:]
		x := (info / 100) % 100
		y := info % 100
		count := info / 10000
		//fmt.Printf("%v %v %v\n",x,y,count)
		if x < 0 || y < 0 || x >= n || y >= n {
			continue
		}
		if ok, _ := m[x*100+y]; ok {
			continue
		}
		if x == n-1 && y == n-1 {
			return count
		}
		m[x*100+y] = true
		for i := 0; i < 8; i++ {
			newx := x + delx[i]
			newy := y + dely[i]
			if newx < 0 || newy < 0 || newx >= n || newy >= n {
				continue
			}
			queue = append(queue, (count+1)*10000+newx*100+newy)
		}

	}
	return -1
}

func main() {
	//fmt.Println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
	//@.a.#
	//###.#
	//b.A.B
	fmt.Println(shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
}
