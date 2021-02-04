package main

import (
	"fmt"
	"math"
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
	dp := make([][]int, 103)
	for i := 0; i < 103; i++ {
		for j := 0; j < 103; j++ {
			dp[i] = append(dp[i], -1)
		}
	}
	cuts = append(cuts, 0)
	cuts = append(cuts, n)
	sort.Ints(cuts)
	return dfsRangeDp(0, len(cuts)-1, cuts, &dp)

}

func dfsRangeDp(l int, r int, cuts []int, dp *[][]int) int {
	if (*dp)[l][r] != -1 {
		return (*dp)[l][r]
	}
	if l+1 == r {
		(*dp)[l][r] = 0
		return 0
	}
	for i := l + 1; i < r; i++ {
		cost := dfsRangeDp(l, i, cuts, dp) + dfsRangeDp(i, r, cuts, dp) + cuts[r] - cuts[l]
		if (*dp)[l][r] == -1 || (*dp)[l][r] > cost {
			(*dp)[l][r] = cost
		}
	}
	return (*dp)[l][r]
}

//区间dp 贪心算不可取
func stoneGameV(stoneValue []int) int {
	dp := make([][]int, 501)
	for i := 0; i < 501; i++ {
		for j := 0; j < 501; j++ {
			dp[i] = append(dp[i], -1)
		}
	}
	sum := make([]int, 501)
	sum[0] = 0
	for i := 0; i < len(stoneValue); i++ {
		sum[i+1] = sum[i] + stoneValue[i]
	}
	return storeDfs(1, len(stoneValue), &dp, &sum)
}

func storeDfs(l int, r int, dp *[][]int, sum *[]int) int {
	if (*dp)[l][r] != -1 {
		return (*dp)[l][r]
	}
	if l == r {
		(*dp)[l][r] = 0
	} else {
		val := 0
		for i := l; i < r; i++ {
			s1 := (*sum)[i] - (*sum)[l-1]
			s2 := (*sum)[r] - (*sum)[i]
			if s1 < s2 {
				val = max(val, s1+storeDfs(l, i, dp, sum))
			} else if s1 > s2 {
				val = max(val, s2+storeDfs(i+1, r, dp, sum))
			} else {
				val = max(val, max(storeDfs(l, i, dp, sum), storeDfs(i+1, r, dp, sum))+s1)
			}
		}
		(*dp)[l][r] = val
	}
	return (*dp)[l][r]
}

//1690. 石子游戏 VII 又是一道区间dp
func stoneGameVII(stones []int) int {
	n := len(stones)
	if n <= 1 {
		return 0
	}
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		temp := stones[i]
		sum[i][i] = temp
		for j := i + 1; j < n; j++ {
			temp += stones[j]
			sum[i][j] = temp
			sum[j][i] = temp
		}
	}
	//dp[i][j] 表示轮到这一个人（无论是 A 或者 B）选时，他能得到与另一个人最大的得分差；
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 区间dp 所以循环要按照从已知区间往未知区间推的方向
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j-i == 1 {
				dp[i][j] = max(stones[i], stones[j])
			} else {
				// point!!!
				dp[i][j] = max(sum[i+1][j]-dp[i+1][j], sum[i][j-1]-dp[i][j-1])
			}
			dp[j][i] = dp[i][j]
		}
	}
	return dp[0][n-1]
}

//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数【插入/删除/替换】 。
// 最短编辑距离 dp[i][j] 表示 word1到i位置转换成word2到j位置的编辑次数
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	// init
	for i := 1; i <= len1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	// dp
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[len1][len2]
}

// 2出现的次数 数位dp
//以dp[i]表示n的1~i位组成的数字所包含2的个数，关键点在于推导出dp[i]与dp[i-1]的关系
/**
0：numberOf2sInRange(02) = numberOf2sInRange(2)
1：numberOf2sInRange(178) = numberOf2sInRange(99) + numberOf2sInRange(78)
2：numberOf2sInRange(233) = 2 * numberOf2sInRange(99) + numberOf2sInRange(33) + 33 + 1
>2:numberOf2sInRange(478) = 4 * numberOf2sInRange(99) + numberOf2sInRange(78) + 100 【200～299 百为】
*/
func numberOf2sInRange(n int) int {
	if n == 0 {
		return 0
	}
	digit := int(math.Log10(float64(n)) + 1)
	dp := make([]int, digit+1)  //numberOf2sInRange(n % pow(10, i)) 保存0~n的1-i位组成的数包含2的个数
	dp9 := make([]int, digit+1) //numberOf2sInRange(99..9) 保存i位均为9包含2的个数
	if n%10 >= 2 {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	dp9[1] = 1
	for i := 2; i <= digit; i++ {
		k := n / int(math.Pow(10, float64(i-1))) % 10
		dp[i] = k*dp9[i-1] + dp[i-1]
		if k == 2 {
			dp[i] += n%int(math.Pow(10, float64(i-1))) + 1
		} else if k > 2 {
			dp[i] += int(math.Pow(10, float64(i-1)))
		}
		// dp9999 = 9*dp999 + 999 + 1000
		dp9[i] = 10*dp9[i-1] + int(math.Pow(10, float64(i-1)))
	}
	return dp[digit]
}

// 元素和最大子矩阵 二维最大子段和 枚举首行和末行 枚举后二维缩成一维求
func getMaxMatrix(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	b := make([]int, m)
	anssumm := math.MinInt16
	var x1, y1 int
	ans := make([]int, 4)
	for i := 0; i < n; i++ { // first row
		for t := 0; t < m; t++ { //  //每次更换子矩形上边，就要清空b，重新计算每列的和
			b[t] = 0
		}
		for j := i; j < n; j++ { // last row
			// 确定枚举出i，j为首行，末行后，按列加和转化为一维b[]最大子段和
			sum := 0
			for k := 0; k < m; k++ {
				b[k] += matrix[j][k]
				if sum > 0 {
					sum += b[k]
				} else {
					sum = b[k]
					x1 = i
					y1 = k
				}
				if sum > anssumm {
					anssumm = sum
					ans = []int{x1, y1, j, k}
				}
			}
		}

	}
	return ans
}

// leetcode 44 通配符匹配
//s = "adceb"
//p = "*a*b"
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

//面试题 08.11. 硬币
////硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
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

//198. 打家劫舍
func rob(nums []int) int {
	// ans[i][0]代表不拿i房的最大 ans[i][1]代表拿i房的最大
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, 2)
	}
	ans[0][0], ans[0][1] = 0, nums[0]
	ans[1][0], ans[1][1] = nums[0], nums[1]
	for i := 2; i < n; i++ {
		ans[i][0] = max(ans[i-1][0], ans[i-1][1])
		ans[i][1] = max(ans[i-1][0]+nums[i], ans[i-2][1]+nums[i])
	}
	return max(ans[n-1][0], ans[n-1][1])
}
func max(ans int, i int) int {
	if ans > i {
		return ans
	} else {
		return i
	}
}

//123. 买卖股票的最佳时机 III 不太好想 先定义状态 然后通过状态找到转移方程
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		//对于buy1而言，在第 ii 天我们可以不进行任何操作，保持不变，也可以在未进行任何操作的前提下以 \textit{prices}[i]prices[i] 的价格买入股票
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
func main() {
	//fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
	//fmt.Println(isMatch("aa", "a*"))
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	//fmt.Println(longestAwesome("3242415"))
	//fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))
	//fmt.Println(minDistance("intention", "execution"))
	//fmt.Println(numberOf2sInRange(25))
	ans := waysToChange(5)
	fmt.Printf("ans:%v", ans)
}
