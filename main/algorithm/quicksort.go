package main

import "fmt"

func quickSort(a *[]int, low int, high int) {
	if low >= high {
		return
	}
	keyPos := partitions(a, low, high)
	quickSort(a, low, keyPos-1)
	quickSort(a, keyPos+1, high)
}

func partitions(a *[]int, low int, high int) int {
	key := (*a)[low]
	for low < high {
		// 从后往前比key大需交换的数
		for high > low && (*a)[high] >= key {
			high--
		}
		//换到前面
		(*a)[low] = (*a)[high]
		// 从前往后比key小的数
		for low < high && (*a)[low] <= key {
			low++
		} // 换到后面
		(*a)[high] = (*a)[low]
	}
	(*a)[low] = key
	return low
}

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	targetIndex := len(nums) - k
	keyPos := partitions(&nums, 0, len(nums)-1)
	if keyPos == targetIndex {
		return nums[keyPos]
	}
	if targetIndex > keyPos {
		return findKthLargest(nums[keyPos+1:], k)
	} else {
		return findKthLargest(nums[0:keyPos], k-(len(nums)-keyPos))
	}
}

// 1,2
func main() {
	//a := []int{3, 5, 8, 6, 2}
	//quickSort(&a, 0, 4)
	//fmt.Println(a)
	fmt.Println(findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
}
