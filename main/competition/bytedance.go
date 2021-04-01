package main

import "sort"

/**
[1,3,2]
输出：
[3,1,2]
预期结果：
[2,1,3]
*/
//31. 下一个排列
func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}
	sort.Ints(nums)
	return
}
func main() {

}
