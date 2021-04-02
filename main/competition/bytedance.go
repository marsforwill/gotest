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

func main() {
	nextPermutation([]int{1, 5, 1})

}
