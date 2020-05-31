package main

import (
	"fmt"
	"math"
)

/*
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 1 1 1 1 2 2 2 2 2 3 3 3 3 3 4 4 4 4 4 5 5 5 5 5 6 6
1 1 1 1 1 2 2 2 2 2 4 4 4 4 4 6 6 6 6 6 9 9 9 9 9 12 12
1 1 1 1 1 2 2 2 2 2 4 4 4 4 4 6 6 6 6 6 9 9 9 9 9 13 13
ans:13
f(4,90)=f(3,90)+f(3,90-25)+f(3,90-2*25)+f(3,90-3*25)
*/
func waysToChange(n int) int {
	var ans = make([]int, n+5)
	coin := []int{1, 5, 10, 25}
	ans[0] = 1
	for i := 0; i < 4; i++ {
		for j := 0; j <= n; j++ {
			if j-coin[i] >= 0 {
				ans[j] += ans[j-coin[i]]
			}
		}
		for j := 0; j <= n; j++ {
			fmt.Printf("%v ", ans[j])
		}
		fmt.Println()
	}
	return ans[n]

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[uint8]int{}
	ans := 0
	j := 0
	k := 0
	for i := 0; i < len(s); i++ {
		var char = s[i]
		_, ok := m[char]
		if ok {
			for ; j < i; j++ {
				if s[j] == char {
					for ; k <= j; k++ {
						delete(m, s[k])
					}
					j++
					break
				}
			}
		} else {
			if i-j+1 > ans {
				ans = i - j + 1
			}
		}
		m[char] = i
	}
	return ans
}

// 输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 说明:
//
// 假设你总是可以到达数组的最后一个位置。
// Related Topics 贪心算法 数组
func jump(nums []int) int {
	lenth := len(nums)
	count := make([]int, lenth)
	count[lenth-1] = 0
	// 从后往前循环的index
	for i := lenth - 2; i >= 0; i-- {
		if nums[i] == 0 {
			count[i] = -1
			continue
		}
		min := math.MaxInt64
		// 可以跳的步数
		for j := 1; j <= nums[i]; j++ {
			if j+i == lenth-1 {
				min = 0
				continue
			}
			if j+i > lenth-1 {
				continue
			}
			if nums[i+j] == 0 {
				continue
			}
			if count[i+j] < min && count[i+j] >= 0 {
				min = count[i+j]
			}
		}
		count[i] = min + 1
	}
	//fmt.Println(count)
	return count[0]
}

func mySqrt(x int) int {
	ans := math.Sqrt(float64(x))
	return int(ans)
}

//[1,4],[4,4],[2,2],[3,4],[1,1]
func main() {
	//ans := waysToChange(5)
	//fmt.Printf("ans:%v",ans)
	//fmt.Println(lengthOfLongestSubstring("asdfsdf"))
	s := [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}
	fmt.Println(maxEvents(s))
	//fmt.Println(jump(s))

}
