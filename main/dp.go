package main

import (
	"fmt"
)

//两边dp
//注意从起点开始能走的范围
//走过的点需要标记
func cherryPickup(grid [][]int) int {
	return 0
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firMatch && isMatch(s[1:], p))
	} else {
		return firMatch && isMatch(s[1:], p[1:])
	}

}

// 枚举中心点
func longestPalindrome(s string) string {
	l := len(s)
	print(l)
	// ans[i]表示已i为终点的最长回文串的长度
	if l <= 1 {
		return s
	}
	// 用中心点再扫一下
	maxm := 0
	ansmi := 0
	for i := 0; i < l; i++ {
		var delta int
		for delta = 1; i-delta >= 0 && i+delta < l; delta++ {
			if s[i-delta] == s[i+delta] {
			} else {
				break
			}
		}
		delta--
		if delta*2+1 > maxm {
			ansmi = i
			maxm = delta*2 + 1
		}
	}
	//中心虚点
	maxn := 0
	ansni := 0
	for i := 1; i < l; i++ {
		if s[i] != s[i-1] {
			continue
		}
		var delta int
		for delta = 1; i-delta-1 >= 0 && i+delta < l; delta++ {
			if s[i-delta-1] == s[i+delta] {
			} else {
				break
			}
		}
		delta--
		if delta*2+1 > maxn {
			ansni = i
			maxn = delta*2 + 1
		}
	}
	fmt.Println(maxm)
	fmt.Println(maxn)
	if maxm > maxn {
		return s[ansmi-(maxm/2) : ansmi+(maxm/2)+1]
	} else {
		return s[ansni-(maxn/2)-1 : ansni+(maxn/2)+1]
	}
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// 一唯 house，二维选择的颜色，value存target
	//dp := make([][]int,109)
	return 0
}

func main() {
	//fmt.Println(isMatch("aa", "a*"))
	fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}
