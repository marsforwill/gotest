package main

import (
	"fmt"
	"sort"
)

func findKthPositive(arr []int, k int) int {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
	}
	for i := 1; i < 1001; i++ {
		_, has := m[i]
		if has {
			continue
		} else {
			k--
		}
		if k == 0 {
			return i
		}
	}
	return 0
}

// 输入：position = [1,2,3,4,7], m = 3
//输出：3
//解释：将 3 个球分别放入位于 1，4 和 7 的三个篮子，两球间的磁力分别为 [3, 3, 6]。最小磁力为 3 。我们没办法让最小磁力大于 3 。
//万万没有想到会是二分 ! check函数 ！ 最大化最小！
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	hi := (position[len(position)-1] - position[0]) / (m - 1)
	lo := 1
	ans := 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if checkDis(position, mid, m) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func checkDis(position []int, mid int, m int) bool {
	count := 1
	i := 0
	for j := 1; j < len(position); j++ {
		if position[j]-position[i] >= mid {
			count++
			i = j
		}
	}
	return count >= m
}

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
}
