package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	num1 := strings.Split(version1, ".")
	num2 := strings.Split(version2, ".")
	var i int
	for i = 0; i < len(num1) || i < len(num2); i++ {
		var n1, n2 int
		if i >= len(num1) {
			n1 = 0
		} else {
			n1, _ = strconv.Atoi(num1[i])
		}

		if i >= len(num2) {
			n2 = 0
		} else {
			n2, _ = strconv.Atoi(num2[i])
		}
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
	}
	return 0
}

// 可用map存储换时间优化到O（n）
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	ans := 0
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			count = 0
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

func subsets(nums []int) [][]int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	var numq []int
	for k := range m {
		numq = append(numq, k)
	}
	var ans [][]int
	dfsset(&ans, numq, []int{}, 0)
	return ans
}

func dfsset(ans *[][]int, numq []int, temp []int, index int) {
	if index == len(numq) {
		*ans = append(*ans, append([]int(nil), temp...))
		return
	}
	dfsset(ans, numq, temp, index+1)
	temp = append(temp, numq[index])
	dfsset(ans, numq, temp, index+1)
}

// 乘积最大连续子数组 自然的dp定义不满足最优子结构 需要根据正负来分类讨论转移
func maxProduct(nums []int) int {
	maxF, minF := nums[0], nums[0]
	maxn := []int{nums[0]}
	minn := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF // max for cur +，min for cur -
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		maxn = append(maxn, maxF)
		minn = append(minn, minF)
	}
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		if ans < maxn[i] {
			ans = maxn[i]
		}
	}
	fmt.Printf("%v\n", maxn)
	fmt.Printf("%v\n", minn)
	return ans
}

// 300.最长递增子序列
func lengthOfLIS(nums []int) int {
	count := make([]int, len(nums))
	count[0] = 1
	for i := 1; i < len(nums); i++ {
		temp := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && temp < count[j]+1 {
				temp = count[j] + 1
			}
		}
		count[i] = temp
	}
	ans := 0
	for i := 0; i < len(count); i++ {
		if ans < count[i] {
			ans = count[i]
		}
	}
	return ans
}

func main() {
	//ans := subsets([]int{9, 0, 3, 5, 7})
	//fmt.Println(ans)
	fmt.Println(maxProduct([]int{1, 0, -1, 7, 8, 0, -2}))
}
