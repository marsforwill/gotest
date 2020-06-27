package main

import "fmt"

func average(salary []int) float64 {
	l := len(salary)
	if l <= 2 {
		return 0
	}
	var max int
	var min int
	max = -1
	min = 100000009
	sum := 0
	for i := 0; i < l; i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		sum += salary[i]
	}
	return float64(sum-max-min) / float64(l-2)
}

func kthFactor(n int, k int) int {
	k--
	if k <= 0 {
		return 1
	}
	for i := 2; i <= n; i++ {
		flag := false
		if n%i == 0 {
			flag = true
			n = n / i
		}
		if flag {
			k--
		}
		if k == 0 {
			return i
		}
	}
	return -1
}

func longestSubarray(nums []int) int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	left[0] = 0
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}
	}
	right[l-1] = 0
	for i := l - 2; i >= 0; i-- {
		if nums[i+1] > 0 {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 0
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		if left[i]+right[i] > ans {
			ans = left[i] + right[i]
		}
	}
	return ans
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	flag := make([]bool, n+1)
	before := make(map[int]map[int]bool)
	ld := len(dependencies)
	for i := 0; i < ld; i++ {
		value, ok := before[dependencies[i][1]]
		if ok {
			value[dependencies[i][0]] = true
		} else {
			tm := make(map[int]bool)
			tm[dependencies[i][0]] = true
			before[dependencies[i][1]] = tm
		}
	}
	ans := 0
	count := 0
	// 遍历每一层
	for count < n {
		temp := k
		// 遍历每一个节点 可选则选
		for i := 1; i <= n; i++ {
			value := before[i]
			// 找到一个
			if len(value) == 0 && flag[i] == false {
				temp--
				flag[i] = true
				for num := 1; num <= n; num++ {
					v := before[num]
					delete(v, num)
				}
				count++
			}
			if temp == 0 {
				break
			}
		}
		ans++
	}
	return ans
}
func main() {
	//fmt.Println(longestSubarray([]int{1,1,0,1}))
	fmt.Println(minNumberOfSemesters(4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2))
}
