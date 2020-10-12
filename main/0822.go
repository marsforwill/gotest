package main

import (
	"fmt"
	"sort"
	"strconv"
)

func thousandSeparator(n int) string {
	str := strconv.FormatInt(int64(n), 10)
	l := len(str)
	count := 0
	for i := l - 1; i > 0; i-- {
		count++
		if count == 3 {
			count = 0
			str = str[0:i] + "." + str[i:]
		}
	}
	return str
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	flag := make([]int, n+5)
	for i := 0; i < len(edges); i++ {
		flag[edges[i][1]]++
	}
	var ans []int
	for i := 0; i < n; i++ {
		if flag[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func minOperations(nums []int) int {
	ans := 0
	max := -1
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		ans += countNumsOfOne(nums[i])
	}
	ans += countBits(max) - 1
	return ans
}

func countBits(m int) int {
	ans := 0
	for m >= 1 {
		m = m >> 1
		ans++
	}
	return ans
}

func countNumsOfOne(n int) int {
	res := 0
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return res
}

//["b","a","c"]
//["c","a","c"]
//["d","d","c"]
//["b","c","c"]
func containsCycle(grid [][]byte) bool {
	n := len(grid)
	if n == 0 {
		return false
	}
	m := len(grid[0])
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfsGrid(grid, i, j, &vis) == true {
				return true
			}
			vis[i][j] = false
		}
	}
	return false
}

func dfsGrid(grid [][]byte, i int, j int, vis *[][]bool) bool {
	(*vis)[i][j] = true
	count := 0
	if i+1 < len(grid) && grid[i+1][j] == grid[i][j] {
		if (*vis)[i+1][j] {
			count++
		} else {
			return dfsGrid(grid, i+1, j, vis)
		}
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == grid[i][j] {
		if (*vis)[i][j+1] {
			count++
		} else {
			return dfsGrid(grid, i, j+1, vis)
		}
	}
	if i-1 >= 0 && grid[i-1][j] == grid[i][j] {
		if (*vis)[i-1][j] {
			count++
		} else {
			return dfsGrid(grid, i-1, j, vis)
		}
	}
	if j-1 >= 0 && grid[i][j-1] == grid[i][j] {
		if (*vis)[i][j-1] {
			count++
		} else {
			return dfsGrid(grid, i, j-1, vis)
		}
	}

	if count >= 2 {
		return true
	}
	(*vis)[i][j] = false
	return false
}

func mostVisited(n int, rounds []int) []int {
	var ans []int
	for i := rounds[0]; ; i++ {
		if i > n {
			i -= n
		}
		ans = append(ans, i)
		if i == rounds[len(rounds)-1] {
			break
		}
	}
	sort.Ints(ans)
	return ans
}

func maxCoins(piles []int) int {
	sort.Ints(piles)
	ans := 0
	flag := false
	count := 0
	for i := len(piles) - 1; i >= 0; i-- {
		if flag == true {
			ans += piles[i]
			count++
		}
		if flag == false {
			flag = true
		} else {
			flag = false
		}
		if count == len(piles)/3 {
			return ans
		}
	}
	return 0
}

//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
// 流程想清楚再 code 写得难受
func removeKdigits(num string, k int) string {
	var stack []uint8
	var result string
	for i := 0; i < len(num); i++ {
		number := num[i] - '0'
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		if number != 0 || len(stack) != 0 {
			stack = append(stack, number)
		}
	}
	for len(stack) != 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	for _, v := range stack {
		result += string('0' + v)
	}
	if result == "" {
		return "0"
	}
	return result
}


func maximalNetworkRank(n int, roads [][]int) int {
	city := make([][]int, n)
	for i := 0; i < len(roads); i++ {
		city[roads[i][0]] = append(city[roads[i][0]], roads[i][1])
		city[roads[i][1]] = append(city[roads[i][1]], roads[i][0])
	}
	ans := -1
	for i := 0; i < n; i++ {
		flag := make([]bool, n)
		for index,j := range city[i]{
			ans = max(ans, len(city[i]) + len(city[j]) - 1)
			flag[city[i][index]] = true
		}
		for j := i+1; j < n; j++ {
			if !flag[j] {
				ans = max(ans, len(city[i]) + len(city[j]))
			}
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


func main() {
	//fmt.Println(thousandSeparator(123456789))
	//fmt.Println(countNumsOfOne(7))
	//fmt.Println(countBits(5))
	//fmt.Println(containsCycle([][]byte{{'b','a','c'},{'c','a','c'},{'d','d','c'},{'b','c','c'}}))
	//fmt.Println(mostVisited(4,[]int{1,3,1,2}))
	//fmt.Println(stoneGameV([]int{68,75,25,50,34,29,77,1,2,69}))
	//str := "1432219"
	//fmt.Println(str[:len(str)-1])
	fmt.Println(removeKdigits("5337", 2))
}
