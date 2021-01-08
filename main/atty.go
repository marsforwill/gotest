package main

import (
	"fmt"
	"math"
	"strconv"
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
func reverse(x int) int {
	bound := 1 << 31
	print(bound)
	if x < bound*-1 || x > bound-1 {
		return 0
	}
	str := strconv.Itoa(x)
	var ans string
	var i int
	for i = len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	for ; i >= 0; i-- {
		if str[i] != '-' {
			ans = ans + string(str[i])
		}
	}
	if x < 0 {
		ans = "-" + ans
	}
	a, _ := strconv.Atoi(ans)
	return a
}

// 双指针扫描 维护一个 含有尽可能多元素且长度最小的集合
func minWindow(s string, t string) string {
	// t串中各个字符出现的次数
	tmap := make(map[uint8]int)
	// 当前扫描中各个字符出现的次数
	vis := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	if len(tmap) == 0 || len(s) == 0 {
		return ""
	}
	check := func() bool {
		for k, v := range tmap {
			if vis[k] < v {
				return false
			}
		}
		return true
	}

	left, right := -1, -1
	slen := len(s)
	ansl := math.MaxUint32
	//在 s 上滑动窗口，通过移动 r 指针不断扩张窗口。当窗口包含 t 全部所需的字符后，如果能收缩，我们就移动l收缩窗口直到得到最小窗口
	for l, r := 0, 0; r < slen; r++ {
		if r < slen && tmap[s[r]] > 0 {
			vis[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < ansl {
				ansl = r - l + 1
				left, right = l, r+1
			}
			if _, ok := tmap[s[l]]; ok {
				vis[s[l]] -= 1
			}
			l++
		}
	}
	if left == -1 {
		return ""
	}

	return s[left:right]
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
	//s := "aa"
	//p := "a"
	//p := "mis*is*p*."
	//fmt.Println(isMatch(s, p))
	fmt.Println(reverse(1534236469))

}
