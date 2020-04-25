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
func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(a))
}
