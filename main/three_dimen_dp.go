package main

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// 一维house 二维当前target 三维颜色 值value cost
	// 二三维度顺序不固定 但是一唯应该是固定的
	var f [125][125][25]int
	var i, j, k, l int
	ans := 1 << 29
	for i = 0; i < 125; i++ {
		for j = 0; j < 125; j++ {
			for k = 0; k < 25; k++ {
				f[i][j][k] = 1 << 30
			}
		}
	}
	f[0][0][0] = 0
	//处理当前第i个house
	for i = 0; i < m; i++ {
		//第i个house的target j
		for j = 0; j <= i; j++ {
			//第i个house颜色 k
			for k = 0; k <= n; k++ {
				if houses[i] > 0 {
					if k == houses[i] { // 相等target+1
						f[i+1][j][houses[i]] = min(f[i+1][j][houses[i]], f[i][j][k])
					} else {
						f[i+1][j+1][houses[i]] = min(f[i+1][j+1][houses[i]], f[i][j][k])
					}
				} else {
					for l = 1; l <= n; l++ {
						if k == l {
							f[i+1][j][l] = min(f[i+1][j][l], f[i][j][k]+cost[i][l-1])
						} else {
							f[i+1][j+1][l] = min(f[i+1][j+1][l], f[i][j][k]+cost[i][l-1])
						}
					}
				}
			}
		}
	}
	// 选m个house target 满足的最小value
	for l = 1; l <= n; l++ {
		ans = min(ans, f[m][target][l])
	}
	if ans == 1<<29 {
		return -1
	}
	return ans
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	} else {
		return i2
	}
}

func main() {

}
