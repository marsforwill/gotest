package main

import (
	"fmt"
	"math"
	"sort"
)

//二分是我万万没有想到的
//要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花
func minDays(bloomDay []int, m int, k int) int {
	l := len(bloomDay)
	if m*k > l {
		return -1
	}
	left := 1
	right := 0
	for i := 0; i < l; i++ {
		if bloomDay[i] > right {
			right = bloomDay[i]
		}
	}
	for left+1 < right {
		mid := (left + right) / 2
		if checkMid(mid, bloomDay, m, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if checkMid(left, bloomDay, m, k) {
		return left
	}
	return right
}

// 校验mid天是不是能满足开花的条件要求
func checkMid(mid int, day []int, m int, k int) bool {
	l := len(day)
	count := 0
	ans := 0
	for i := 0; i < l; i++ {
		if day[i] <= mid {
			count++
		} else {
			count = 0
		}
		if count == k {
			ans++
			count = 0
		}
	}
	return ans > m
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

//875. 爱吃香蕉的珂珂 二分算法
func minEatingSpeed(piles []int, H int) int {
	sort.Ints(piles)
	left, right := 1, piles[len(piles)-1]
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		if checkSpeed(piles, H, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// 判断以speed速度能不能在h小时内吃完
func checkSpeed(piles []int, h int, speed int) bool {
	hour := 0
	for i := 0; i < len(piles); i++ {
		if speed >= piles[i] {
			hour++
		} else {
			hour += int(math.Ceil(float64(piles[i]) / float64(speed)))
		}
		if hour > h {
			return false
		}
	}
	return true
}

func main() {
	//minDays([]int{1, 10, 3, 10, 2}, 3, 1)
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
