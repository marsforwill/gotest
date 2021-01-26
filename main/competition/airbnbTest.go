package main

import "fmt"

//39. 组合总和
// 深度搜索回溯 可加剪枝优化
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var comb []int
	var dfsComb func(target int, idx int)
	dfsComb = func(target int, idx int) {
		// 递归出口
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//不选当前数
		dfsComb(target, idx+1)
		//选当前数
		if target >= candidates[idx] {
			comb = append(comb, candidates[idx])
			dfsComb(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfsComb(target, 0)
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
