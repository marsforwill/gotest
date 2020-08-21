package main

import (
	"fmt"
	"sort"
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

//leetcode submit region begin(Prohibit modification and deletion)
// 输入：s = "3242415"
//输出：5
//解释："24241" 是最长的超赞子字符串，交换其中的字符后，可以得到回文 "24142"
// 状态压缩dp
func longestAwesome(s string) int {
	// 压缩状态 -> index
	m := make(map[int]int)
	cur := 0
	ans := 1
	m[cur] = -1
	for i := 0; i < len(s); i++ {
		ch := s[i] - '0'
		cur = cur ^ (1 << ch)
		// 奇数次
		for j := 0; j < 10; j++ {
			// 可容忍差一个的回文状态
			next := cur ^ (1 << j)
			// 得到index 取相差长度最大
			index, ok := m[next]
			if ok {
				ans = max(ans, i-index)
			}
		}
		// 偶数次
		index, ok := m[cur]
		if !ok {
			m[cur] = i
		} else {
			ans = max(ans, i-index)
		}
	}
	return ans
}

func max(ans int, i int) int {
	if ans > i {
		return ans
	} else {
		return i
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}


// 输入：n = 9, cuts = [5,6,1,4,2]
//输出：22
//解释：如果按给定的顺序切割，则总成本为 25 。总成本 <= 25 的切割顺序很多，例如，[4，6，5，2，1] 的总成本 = 22，是所有可能方案中成本最
//小的。
// 区间dp dp[l][r] 代表切cut[l] cut[r]的最小成本
func minCost(n int, cuts []int) int {
	dp := make([][]int,103)
	for i := 0; i < 103; i++ {
		for j := 0; j < 103; j++ {
			dp[i] = append(dp[i],-1)
		}
	}
	cuts = append(cuts, 0)
	cuts = append(cuts, n)
	sort.Ints(cuts)
	return dfsRangeDp(0,len(cuts)-1, cuts, &dp)

}

func dfsRangeDp(l int, r int, cuts []int, dp *[][]int) int {
	if (*dp)[l][r] != -1 {
		return (*dp)[l][r]
	}
	if l+1 == r {
		(*dp)[l][r] = 0
		return 0
	}
	for i := l+1; i < r; i++ {
		cost := dfsRangeDp(l,i,cuts,dp) + dfsRangeDp(i,r,cuts,dp) + cuts[r] - cuts[l]
		if (*dp)[l][r] == -1 || (*dp)[l][r] > cost {
			(*dp)[l][r] = cost
		}
	}
	return (*dp)[l][r]
}

func main() {
	//fmt.Println(isMatch("aa", "a*"))
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(longestAwesome("3242415"))
}
