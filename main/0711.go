package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	l := n
	sum := make([]int, l)
	count := 0
	for i := 0; i < l; i++ {
		count += nums[i]
		sum[i] = count
	}
	var temp []int
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			temp = append(temp, (sum[i]-sum[j])%mod)
		}
	}
	for i := 0; i < n; i++ {
		temp = append(temp, sum[i]%mod)
	}
	sort.Ints(temp)
	ans := 0
	for i := left - 1; i <= right-1; i++ {
		ans += temp[i]
	}
	return ans % mod
}

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	ans := 1000000007
	if nums[n-4]-nums[0] < ans {
		ans = nums[n-4] - nums[0]
	}
	if nums[n-3]-nums[1] < ans {
		ans = nums[n-3] - nums[1]
	}
	if nums[n-2]-nums[2] < ans {
		ans = nums[n-2] - nums[2]
	}
	if nums[n-1]-nums[3] < ans {
		ans = nums[n-1] - nums[3]
	}
	return ans

}
func winnerSquareGame(n int) bool {
	flag := make([]bool, n+1)
	for i := 1; i*i <= n; i++ {
		flag[i*i] = true
	}
	if flag[n] == true {
		return true
	}
	for i := 2; i <= n; i++ {
		if flag[i] {
			continue
		}
		f := false
		for num := 1; num*num < i; num++ {
			if flag[i-num*num] == false {
				flag[i] = true
				f = true
				break
			}
		}
		if f == false {
			flag[i] = false
		}
	}
	return flag[n]
}

func numIdenticalPairs(nums []int) int {
	l := len(nums)
	count := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func numSub(s string) int {
	mod := 1000000007
	l := len(s)
	ans := 0
	flag := 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			flag = 0
		} else {
			flag++
			ans += flag
			ans = ans % mod
		}
	}
	return ans % mod
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// 最短路径的状态数组
	var dp [][]float64
	// 先初始化
	for i := 0; i < n; i++ {
		var tmp []float64
		for j := 0; j < n; j++ {
			if i == j {
				tmp = append(tmp, 0)
			} else {
				tmp = append(tmp, -1)
			}
		}
		dp = append(dp, tmp)
	}
	// 填出边长
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		weight := succProb[i]
		// 无向图
		dp[from][to] = weight
		dp[to][from] = weight
	}
	// dp状态转移方程
	// k放在第一层是因为后面的k要依赖前面的值
	for k := 0; k < n; k++ {
		// 从i到j
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// 相同的节点不考虑
				if i == j || i == k || j == k {
					continue
				}
				// 不通的路也不考虑
				if dp[i][k] == -1 || dp[k][j] == -1 {
					continue
				}
				tmp := dp[i][k] * dp[k][j]
				if dp[i][j] == -1 || dp[i][j] < tmp {
					dp[i][j] = tmp
					dp[j][i] = tmp
				}
			}
		}
	}
	if dp[start][end] < 0 {
		return 0
	}
	return dp[start][end]

}

func maxProbability2(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// 邻接矩阵描述两点间得距离
	dijkstra := make([][]float64, n)
	for i := 0; i < n; i++ {
		dijkstra[i] = make([]float64, n)
	}
	// 初始化距离 -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				dijkstra[i][j] = 0
			}
		}
	}
	// 初始化 边距离
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		weight := succProb[i]
		// 无向图
		dijkstra[from][to] = weight
		dijkstra[to][from] = weight
	}

	from := start
	//u 待处理的节点
	u := map[int]bool{}
	for k := 0; k < n; k++ {
		if k == from {
			continue
		}
		u[k] = true
	}
	// 每个节点和from的距离
	dis := make([]float64, n)
	for k := 0; k < n; k++ {
		// 每一个节点和from的距离
		dis[k] = dijkstra[from][k]
	}

	// 处理所有节点的距离
	for {
		if len(u) == 0 {
			break
		}

		toFormaxdis := -2.0
		toFormaxdisIndex := -1
		// 找当前最大的距离节点
		for to := range u {
			if dis[to] <= 0 || toFormaxdis > dis[to] {
				continue
			}
			toFormaxdis = dis[to]
			toFormaxdisIndex = to
		}
		if toFormaxdisIndex == -1 {
			break
		}
		delete(u, toFormaxdisIndex)

		for j := 0; j < n; j++ {
			_, ok := u[j]
			if !ok || dijkstra[toFormaxdisIndex][j] <= 0 {
				continue
			}

			indirectDis := dis[toFormaxdisIndex] * dijkstra[toFormaxdisIndex][j]
			if dis[j] < 0 {
				dis[j] = indirectDis
				continue
			}
			if indirectDis > dis[j] {
				dis[j] = indirectDis
			}
		}
	}
	if dis[end] < 0 {
		return 0
	}
	return dis[end]
}
func main() {
	//fmt.Println(rangeSum([]int{1,2,3,4},4,1,5))
	//fmt.Println(winnerSquareGame(15))
	fmt.Println(maxProbability2(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
}
