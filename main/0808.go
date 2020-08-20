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

//// S1 = "0"
//// S2 = "011"
//// S3 = "0111001"
//// S4 = "011100110110001"
func findKthBit(n int, k int) byte {
	if n == 1 || k == 1 {
		return '0'
	}
	l := (1 << n) - 1
	mid := (l >> 1) + 1
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		return invert(findKthBit(n-1, l-k+1))
	}
}

func invert(bit byte) byte {
	if bit == '0' {
		return '1'
	} else {
		return '0'
	}
}

//返回 非空不重叠 子数组的最大数目，且每个子数组中数字和都为 target
// 输入：nums = [-1,3,5,1,4,2,-9], target = 6
//输出：2
//解释：总共有 3 个子数组和为 6 。
//([5,1], [4,2], [3,5,1,4,2,-9]) 但只有前 2 个是不重叠的。
// // 前缀和 + hash 不需要dp 贪心可行！！！！！！！！！！！！！！！！！
func maxNonOverlapping(nums []int, target int) int {
	m := make(map[int]int)
	sum := 0
	ans := 0
	m[0] = -1
	from := -1
	// 前缀和 + hash
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		index, ok := m[sum-target]
		if ok && index >= from {
			ans++
			from = i
		}
		m[sum] = i
	}
	return ans
}

func main() {
	//fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
	//fmt.Printf("%c",findKthBit(4,11))
	fmt.Println(maxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
}
