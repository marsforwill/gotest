package main

import (
	"fmt"
	"sort"
)

// 加两个数组存map
func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	ans := 0
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := A[i] + B[j]
			v, ok := m[sum]
			if ok {
				m[sum] = v + 1
			} else {
				m[sum] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := C[i] + D[j]
			v, ok := m[0-sum]
			if ok {
				ans += v
			}
		}
	}
	return ans
}

// 统计字符出现的种类和字数 blabla
func closeStrings(word1 string, word2 string) bool {
	l1 := len(word1)
	l2 := len(word2)
	if l1 != l2 {
		return false
	}
	m := make(map[uint8]int)
	for i := 0; i < l1; i++ {
		a := word1[i]
		v,ok := m[a]
		if ok {
			m[a] = v+1
		} else {
			m[a] = 1
		}
	}
	m2 := make(map[uint8]int)
	for i := 0; i < l1; i++ {
		a := word2[i]
		_,o := m[a]
		if o == false {
			return false
		}
		v,ok := m2[a]
		if ok {
			m2[a] = v+1
		} else {
			m2[a] = 1
		}
	}
	var list1,list2 []int
	for k := range m {
		list1 = append(list1, m[k])
	}
	for k:= range m2 {
		list2 = append(list2, m2[k])
	}
	sort.Ints(list1)
	sort.Ints(list2)
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

/**
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用
双指针扫描 就是中间逻辑有点恶心
*/
func minOperations(nums []int, x int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	ans := -1
	target := sum-x
	if target == 0 {
		return len(nums)
	}
	left, right, s := 0,0,nums[0]
	for left < len(nums)  && right < len(nums) {
		if s < target {
			right++
			if right == len(nums) {
				continue
			}
			s+=nums[right]
		} else if s > target {
			s-=nums[left]
			left++
		} else if s == target {
			if  right-left > ans {
				ans = right-left
			}
			right++
			if right == len(nums) {
				continue
			}
			s+=nums[right]
		}
	}
	if ans == -1 {
		return -1
	}
	return len(nums) - ans -1
}

func main() {
	fmt.Println(minOperations([]int{500,1,4,2,3},500))
}
