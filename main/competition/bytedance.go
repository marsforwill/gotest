package main

/**
输入：nums = [1,2,3]
输出：[1,3,2]
*/
//31. 下一个排列
func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		// 找到第一个逆序升序index
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i && j >= 0; j-- {
				if nums[i] < nums[j] {
					// 交换 比index大的最小元素（末尾）
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			// reverse 升序变降序
			reverse(nums[i+1:])
			return
		}
	}
	reverse(nums)
	return
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 46. 全排列 dfs 搜索树回溯
func permute(nums []int) [][]int {
	var res [][]int
	vis := make(map[int]bool)
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(path, temp)
			res = append(res, temp)
			return
		}
		for _, i := range nums {
			if vis[i] {
				continue
			}
			path = append(path, i)
			vis[i] = true
			dfs(path)
			path = path[:len(path)-1]
			vis[i] = false
		}
	}
	dfs([]int{})
	return res
}
func main() {
	nextPermutation([]int{1, 5, 1})

}
