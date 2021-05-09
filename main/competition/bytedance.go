package main

import (
	"strconv"
	"strings"
)

/**
输入：nums = [1,2,3]
输出：[1,3,2]
*/
//31. 下一个排列
func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		// 找到第一个逆序升序index
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i && j >= 0; j-- {
				if nums[i] < nums[j] {
					// 交换 比index大的最小元素（末尾）
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			// reverse 升序变降序
			reverse(nums[i+1:])
			return
		}
	}
	reverse(nums)
	return
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 46. 全排列 dfs 搜索树回溯
func permute(nums []int) [][]int {
	var res [][]int
	vis := make(map[int]bool)
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(path, temp)
			res = append(res, temp)
			return
		}
		for _, i := range nums {
			if vis[i] {
				continue
			}
			path = append(path, i)
			vis[i] = true
			dfs(path)
			path = path[:len(path)-1]
			vis[i] = false
		}
	}
	dfs([]int{})
	return res
}

//79. 单词搜索 dfs搜索回溯
func exist(board [][]byte, word string) bool {
	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfsBoard(i, j, board, word, vis, 0) {
				return true
			}
		}
	}
	return false
}

var dx = []int{0, 0, -1, 1}
var dy = []int{-1, 1, 0, 0}

func dfsBoard(x int, y int, board [][]byte, word string, vis [][]bool, idx int) bool {
	if idx == len(word) {
		return true
	}
	if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] == word[idx] && vis[x][y] == false {
		for i := 0; i < 4; i++ {
			vis[x][y] = true
			if dfsBoard(x+dx[i], y+dy[i], board, word, vis, idx+1) {
				return true
			}
			vis[x][y] = false
		}
	}
	return false
}

//93. 复原 IP 地址
func restoreIpAddresses(s string) []string {
	var ans []string
	for i := 1; i < len(s) && i <= 3; i++ {
		for j := i + 1; j < len(s) && j < i+4; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				str1 := s[0:i]
				str2 := s[i:j]
				str3 := s[j:k]
				str4 := s[k:]
				if len(str1)+len(str2)+len(str3)+len(str4) == len(s) {
					num1, er1 := strconv.Atoi(str1)
					num2, er2 := strconv.Atoi(str2)
					num3, er3 := strconv.Atoi(str3)
					num4, er4 := strconv.Atoi(str4)
					if (str1[0] == '0' && num1 > 0) || (num1 == 0 && len(str1) > 1) {
						continue
					}
					if (str2[0] == '0' && num2 > 0) || (num2 == 0 && len(str2) > 1) {
						continue
					}
					if (str3[0] == '0' && num3 > 0) || (num3 == 0 && len(str3) > 1) {
						continue
					}
					if (str4[0] == '0' && num4 > 0) || (num4 == 0 && len(str4) > 1) {
						continue
					}
					if er1 == nil && er2 == nil && er3 == nil && er4 == nil && num1 <= 255 && num2 <= 255 && num3 <= 255 && num4 <= 255 {
						ans = append(ans, strings.Join([]string{str1, str2, str3, str4}, "."))
					}
				}
			}
		}
	}
	return ans
}

//32. 最长有效括号 匹配括号 dp/栈/
func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

//89. 格雷编码
func grayCode(n int) []int {
	ans := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			num := ans[j] + head
			ans = append(ans, num)
		}
		head = head << 1
	}
	return ans
}
func main() {
	nextPermutation([]int{1, 5, 1})

}
