package main

import "math"

//单源最短路 可环无负边 format dij code
//787. K 站中转内最便宜的航班
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := make([][]int, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < len(flights); i++ {
		graph[flights[i][0]][flights[i][1]] = flights[i][2]
	}
	var fakepq [][]int
	vis := make(map[int]int) // node+k*1000, cost
	// 放入初始源节点 维护可达未访问最小节点集合
	fakepq = append(fakepq, []int{0, 0, src}) // cost k node
	// 当前处理节点不为空
	for len(fakepq) > 0 {
		// 此处可以从n简化位logN fakepq 模拟最小优先队列
		min := math.MaxInt32
		mini := -1
		for i := 0; i < len(fakepq); i++ {
			if fakepq[i][0] < min {
				min = fakepq[i][0]
				mini = i
			}
		}
		// 选取当前cost最小节点
		item := fakepq[mini]
		fakepq = append(fakepq[0:mini], fakepq[mini+1:]...)
		cost, k, node := item[0], item[1], item[2]
		if k > K+1 {
			continue
		}
		c, ok := vis[k*1000+node]
		if ok && cost > c {
			continue
		}
		// 每一步遍历处理的都是最小的
		if node == dst {
			return cost
		}
		// 将进一步可达且未访问符合条件的节点加入
		for nei := 0; nei < n; nei++ {
			if graph[node][nei] > 0 {
				newCost := cost + graph[node][nei]
				v, ok := vis[nei+(k+1)*1000]
				if (ok && newCost < v) || !ok {
					fakepq = append(fakepq, []int{newCost, k + 1, nei})
					vis[nei+(k+1)*1000] = newCost
				}
			}
		}
	}
	return -1
}

func main() {

}
