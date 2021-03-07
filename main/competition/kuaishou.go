package main

import (
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

func main() {

}
