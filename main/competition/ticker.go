package main

import (
	"fmt"
	"math"
)

// 示例 2：
//
// 输入：days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
//输出：17
//解释：
//例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
//在第 1 天，你花了 costs[2] = $15 买了一张为期 30 天的通行证，它将在第 1, 2, ..., 30 天生效。
//在第 31 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 31 天生效。
//你总共花了 $17，并完成了你计划的每一天旅行。
//
//
//
//
// 提示：
//
//
//i 1 <= days.length <= 365
// 1 <= days[i] <= 365
// days 按顺序严格递增
// costs.length == 3
// 1 <= costs[i] <= 1000
//
// Related Topics 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
var temp = make([]int, 400)

func mincostTickets(days []int, costs []int) int {
	for i, _ := range temp {
		temp[i] = -1
	}
	ans := dfs(days[0], days, costs)
	return ans
}

// 从day天开始走的后续的花费
func dfs(day int, days []int, costs []int) int {
	if temp[day] > 0 {
		return temp[day]
	}
	value := []int{1, 7, 30}
	lenday := len(days)
	if day > days[lenday-1] {
		return 0
	}
	for i := 0; i < lenday; i++ {
		if day == days[i] {
			break
		}
		if day > days[i] && i+1 < lenday && day < days[i+1] {
			day = days[i+1]
			break
		}
	}
	min := math.MaxInt32
	var count int
	for i := 0; i < 3; i++ {
		nextday := day + value[i]
		if nextday > days[lenday-1] {
			if costs[i] < min {
				min = costs[i]
			}
		} else {
			count = dfs(nextday, days, costs)
			temp[nextday] = count
			if count+costs[i] < min {
				min = count + costs[i]
			}
		}
	}
	return min
}

func main() {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))
	//test := dfs(1,days,costs)
	//fmt.Println(test)
}
