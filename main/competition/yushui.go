package main

import "fmt"

func trap(height []int) int {

	length := len(height)
	if length < 3 {
		return 0
	}
	var leftMax = make([]int, length)
	var rightMax = make([]int, length)
	lmax := height[0]
	for i := 0; i < length-1; i++ {
		if height[i] > lmax {
			lmax = height[i]
		}
		leftMax[i+1] = lmax
	}
	rmax := height[length-1]
	for i := length - 2; i > 0; i-- {
		rightMax[i] = rmax
		if height[i] > rmax {
			rmax = height[i]
		}
	}
	sum := 0
	for i := 1; i < length-1; i++ {
		if height[i] > leftMax[i] || height[i] > rightMax[i] {
			continue
		}
		if leftMax[i] > rightMax[i] {
			sum += rightMax[i] - height[i]
		} else {
			sum += leftMax[i] - height[i]
		}
	}
	return sum
}

//[ 0,1,0,2,1,0,1,3,2,1,2,1]
//
//l:0,0,1,1,2,2,2,2,3,3,3,3
//r:0,3,3,3,3,3,3,2,2,2,1,0

func setZeroes(matrix [][]int) {
	col := make([]int, 5000)
	row := make([]int, 5000)
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
	return
}

// 11.盛最多水的容器 虽然暴力也能过 但是最优应该是首尾双指针向中间移动
func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}
	ans := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if height[i] > height[j] {
				ans = max(ans, (j-i)*height[j])
			} else {
				ans = max(ans, (j-i)*height[i])
			}
		}
	}
	return ans
}
func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(a))
}
