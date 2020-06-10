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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	num := make([]int, len1+len2)
	i, j, cur := 0, 0, 0
	for len1 != 0 && len2 != 0 {
		if nums1[i] < nums2[j] {
			num[cur] = nums1[i]
			i++
			cur++
			if i == len1 {
				for k := j; k < len2; k++ {
					num[cur] = nums2[k]
					cur++
				}
				break
			}
		} else {
			num[cur] = nums2[j]
			j++
			cur++
			if j == len2 {
				for k := i; k < len1; k++ {
					num[cur] = nums1[k]
					cur++
				}
				break
			}
		}
	}
	if len1 == 0 {
		num = nums2
		cur = len2
	}
	if len2 == 0 {
		num = nums1
		cur = len1
	}
	index := int(cur / 2)
	if cur%2 == 0 {
		return float64(num[index]+num[index-1]) / 2
	} else {
		return float64(num[index])
	}
}

//
// 输入:
//s = "mississippi"
//p = "mis*is*p*."
//输出: false
// Related Topics 字符串 动态规划 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	lenp := len(p)
	lens := len(s)
	j := 0
	// 处理* 需要截取 判断
	for i := 0; i < lens; i++ {
		if j+1 < lenp && p[j+1] == '*' && p[j] != s[i] {
			j = j + 2
			i--
			continue
		}
		if j+1 < lenp && p[j+1] == '*' {
			var k, h int
			for k = i; s[k] == s[i]; k++ {
				println(k)
			}
			temps := s[i:k]
			for h = j; p[h] == p[j] || p[h] == '*'; h++ {
				println(h)
			}
			tempp := p[j:h]
			if check(temps, tempp) == false {
				return false
			}
			i = k
			j = h + 1
		} else {
			if j > lenp-1 {
				return false
			}
			if s[i] != p[j] && p[j] != '.' {
				return false
			}
			j++
		}
	}
	return true
}

func check(temps string, tempp string) bool {
	fmt.Println(tempp)
	fmt.Println(temps)
	ls := len(temps)
	lp := len(tempp)
	for i := lp - 1; i >= 0; i-- {
		if i > lp-1 || i > ls-1 {
			return true
		}
		if tempp[i] == '*' {
			return true
		}
		if tempp[i] != temps[i] {
			return false
		}
	}
	return true
}

//[1,4],[4,4],[2,2],[3,4],[1,1]
func main() {
	//ans := waysToChange(5)
	//fmt.Printf("ans:%v",ans)
	//fmt.Println(lengthOfLongestSubstring("asdfsdf"))
	//s := [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}
	//fmt.Println(maxEvents(s))
	//fmt.Println(jump(s))
	//num1 := []int{}
	//num2 := []int{3, 4}
	//fmt.Println(findMedianSortedArrays(num1, num2))
	s := "aa"
	p := "a"
	//p := "mis*is*p*."
	fmt.Println(isMatch(s, p))

}
