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

// 统计全1子矩形 不看题解想不明白系列
func numSubmat(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	left := make([][]int, n)
	for i := 0; i < n; i++ {
		left[i] = make([]int, m)
	}
	now := 0
	for i := 0; i < n; i++ {
		now = 0
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				now = 0
			} else {
				now++
			}
			left[i][j] = now
		}
	}
	count := 0
	var minx int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			minx = 1 << 20
			for k := i; k >= 0; k-- {
				minx = min(minx, left[k][j])
				count += minx
			}
		}
	}
	return count
}

func main() {
	//fmt.Println(isMatch("aa", "a*"))
	fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}
