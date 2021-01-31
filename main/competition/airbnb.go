package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func validate_xml(xml string) string {
	// Write your code here
	var tags []string
	for i := 0; i < len(xml); i++ {
		if xml[i] == '>' {
			return "parse error"
		}
		if xml[i] == '<' {
			// start tag
			if i+1 < len(xml) && xml[i+1] != '/' {
				var j int
				for j = i + 1; j < len(xml) && xml[j] != '>'; j++ {
					if xml[j] == '<' {
						return "parse error"
					}
				}
				if j == len(xml) || j == i+1 {
					return "parse error"
				}
				startTag := xml[i+1 : j]
				tags = append(tags, startTag)
				i = j
			} else if i+1 < len(xml) && xml[i+1] == '/' {
				// end tag
				var j int
				for j = i + 2; j < len(xml) && xml[j] != '>'; j++ {
					if xml[j] == '<' {
						return "parse error"
					}
				}
				if j == len(xml) || j == i+2 {
					return "parse error"
				}
				endTag := xml[i+2 : j]
				if len(tags) > 0 && tags[len(tags)-1] == endTag {
					tags = tags[:len(tags)-1]
				} else {
					return "encountered closing tag without matching open tag for </" + endTag + ">"
				}
				i = j
			} else {
				return "parse error"
			}
		}
	}
	if len(tags) > 0 {
		return "missing closing tag for <" + tags[0] + ">"
	}
	return "valid"
}

//39. 组合总和
// 深度搜索回溯 可加剪枝优化
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var comb []int
	var dfsComb func(target int, idx int)
	dfsComb = func(target int, idx int) {
		// 递归出口
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//不选当前数
		dfsComb(target, idx+1)
		//选当前数
		if target >= candidates[idx] {
			comb = append(comb, candidates[idx])
			dfsComb(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfsComb(target, 0)
	return ans
}

//322. 零钱兑换 可以贪心dfs 或者动态规划
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数
func coinChange(coins []int, amount int) int {

	sort.Ints(coins)
	min := math.MaxInt64
	var dfsCoin func(coins []int, idx int, amount int, count int)
	dfsCoin = func(coins []int, idx int, amount int, count int) {
		if amount == 0 {
			if min > count {
				min = count
			}
			return
		}
		if amount < 0 || idx < 0 {
			return
		}
		for i := amount / coins[idx]; i >= 0; i-- {
			if i+count > min {
				break
			}
			dfsCoin(coins, idx-1, amount-i*coins[idx], count+i)
		}
		return
	}
	dfsCoin(coins, len(coins)-1, amount, 0)
	if min == math.MaxInt64 {
		return -1
	}
	return min
}

// 为什么会有这么奇怪的ip细节题
// 751. IP 到 CIDR 给定一个起始 IP 地址 ip 和一个我们需要包含的 IP 的数量 n，返回用列表（最小可能的长度）表示的 CIDR块的范围。
func ipToCIDR(ip string, n int) []string {
	start := ipToInt(ip)
	var ans []string
	for n > 0 {
		// start 和 n的最大掩码
		// 在一般情况下啊，我们使用 n 和 start&-start（start 的最低非零位）的位长度来计算能表示 2^{32 - \text{mask}}2 32−mask
		//  个 ip 地址的掩码。然后，我们动态调整 start 和 n。
		mask := max(33-bitLength((start&-start)),
			33-bitLength(n))
		if mask > 32 {
			mask = 32
		}
		ans = append(ans, intToIp(start)+"/"+strconv.Itoa(mask))
		start += 1 << (32 - mask)
		n -= 1 << (32 - mask)

	}
	return ans
}

func bitLength(num int) int {
	ans := 0
	for num > 0 {
		num = num >> 1
		ans++
	}
	return ans
}

func ipToInt(ip string) int {
	ans := 0
	strs := strings.Split(ip, ".")
	for _, x := range strs {
		v, _ := strconv.Atoi(x)
		ans = 256*ans + v
	}
	return ans
}
func intToIp(x int) string {
	var ans string
	for i := 0; i < 25; i += 8 {
		ans = strconv.Itoa((x>>i)%256) + ans
		if i != 24 {
			ans = "." + ans
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

//324. 摆动排序 II
func wiggleSort(nums []int) {
	sort.Ints(nums)
	var min, max []int
	for i := 0; i < len(nums); i++ {
		if i < len(nums)/2 {
			min = append(min, nums[i])
		} else {
			max = append(max, nums[i])
		}
	}
	if len(max) > len(min) {
		temp := max[0]
		max = max[1:]
		min = append(min, temp)
	}
	j := len(min) - 1
	k := len(max) - 1
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = min[j]
			j--
		} else {
			nums[i] = max[k]
			k--
		}
	}
	return
}

//1166. 设计文件系统
type FileSystem struct {
	value int
	m     map[string]FileSystem
}

func Constructor() FileSystem {
	return FileSystem{
		value: 0,
		m:     make(map[string]FileSystem),
	}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if path[0] == '/' {
		path = path[1:]
	} else {
		return false
	}
	paths := strings.Split(path, "/")
	tempMap := this.m
	i := 0
	for i = 0; i < len(paths)-1; i++ {
		v, ok := tempMap[paths[i]]
		if !ok {
			return false
		}
		tempMap = v.m
	}
	_, ok := tempMap[paths[i]]
	if ok {
		return false
	} else {
		tempMap[paths[i]] = FileSystem{
			value: value,
			m:     make(map[string]FileSystem),
		}
		return true
	}
}

func (this *FileSystem) Get(path string) int {
	if path[0] == '/' {
		path = path[1:]
	} else {
		return -1
	}
	paths := strings.Split(path, "/")
	tempMap := this.m
	i := 0
	for i = 0; i < len(paths); i++ {
		v, ok := tempMap[paths[i]]
		if !ok {
			return -1
		}
		tempMap = v.m
		if i == len(paths)-1 {
			return v.value
		}
	}
	return -1
}

// 212. 单词搜索 II dfs  // 深度搜索回溯
func findWords(board [][]byte, words []string) []string {
	var ans []string
	for i := 0; i < len(words); i++ {
		str := words[i]
		vis := make([][]bool, len(board))
		for i := 0; i < len(board); i++ {
			vis[i] = make([]bool, len(board[0]))
		}
		for j := 0; j < len(board); j++ {
			flag := false
			for k := 0; k < len(board[j]); k++ {
				if board[j][k] == str[0] {
					if dfsWord(board, words[i], 0, j, k, &vis) || len(str) == 1 {
						ans = append(ans, words[i])
						flag = true
						break
					}
					vis[j][k] = false
				}
			}
			if flag {
				break
			}
		}
	}
	return ans
}

func dfsWord(board [][]byte, s string, index int, x int, y int, vis *[][]bool) bool {
	if index == len(s) {
		return true
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || (*vis)[x][y] == true {
		return false
	}
	if board[x][y] != s[index] {
		return false
	}
	(*vis)[x][y] = true

	if dfsWord(board, s, index+1, x+1, y, vis) {
		return true
	}
	if dfsWord(board, s, index+1, x-1, y, vis) {
		return true
	}
	if dfsWord(board, s, index+1, x, y+1, vis) {
		return true
	}
	if dfsWord(board, s, index+1, x, y-1, vis) {
		return true
	}

	(*vis)[x][y] = false
	return false
}

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	var phoneMap map[string]string = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var dfsDigit func(digits string, index int, temp string)
	dfsDigit = func(digits string, index int, temp string) {
		if index == len(digits) {
			ans = append(ans, temp)
			return
		}
		if digits[index] != '1' {
			str, ok := phoneMap[string(digits[index])]
			if ok {
				for i := 0; i < len(str); i++ {
					dfsDigit(digits, index+1, temp+string(str[i]))
				}
			}
		} else {
			dfsDigit(digits, index+1, temp)

		}
		return
	}
	dfsDigit(digits, 0, "")
	return ans
}

//43. 字符串相乘 大数乘法 竖式优化
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			//num1[i] x num2[j] 的结果为 tmp(位数为两位，"0x","xy"的形式)，其第一位位于 res[i+j]，第二位位于 res[i+j+1]
			sum := res[i+j+1] + int(n1*n2)
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	var ans string
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		ans = ans + strconv.Itoa(res[i])
	}
	return ans
}

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		ans[i][0] = 1
		ans[i][i] = 1
	}
	for i := 2; i < numRows; i++ {
		for j := 1; j < len(ans[i])-1; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
func main() {
	//fmt.Println(validate_xml("<a></a>"))
	//	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//fmt.Println(ipToCIDR("255.0.0.7", 10))
	//wiggleSort([]int{1, 1, 2, 1, 2, 2, 1})
	//obj := Constructor()
	//param_1 := obj.CreatePath("/a", 1)
	//param_2 := obj.Get("/a")
	//fmt.Println(param_1)
	//fmt.Println(param_2)
	//fmt.Println(findWords([][]byte{{'a', 'a'}}, []string{"aaa"}))
	fmt.Println(multiply("123", "45"))
}
