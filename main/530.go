package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func canBeEqual(target []int, arr []int) bool {
	lt := len(target)
	la := len(arr)
	if lt != la {
		return false
	}
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < la; i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}

func hasAllCodes(s string, k int) bool {
	count := math.Pow(float64(2), float64(k))
	n := int(count)
	for i := 0; i < n; i++ {
		str := fmt.Sprintf("%b", i)
		if len(str) < k {
			str = strings.Repeat("0", k-len(str)) + str
		}
		if !strings.Contains(s, str) {
			return false
		}
	}
	return true
}
func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	// ans[i] 存 i的一级后续
	ans := make([][]int, n+5)
	for i := 0; i < len(prerequisites); i++ {
		ans[prerequisites[i][0]] = append(ans[prerequisites[i][0]], prerequisites[i][1])
	}
	l := len(queries)
	ansf := make([][]bool, n+5)
	// 循环每一个begin，
	for i := 0; i <= n; i++ {
		dfsAdd(i, ans, &ansf)
	}
	var final []bool
	for i := 0; i < l; i++ {
		if ansf[queries[i][0]][queries[i][1]] {
			final = append(final, true)
		} else {
			final = append(final, false)
		}
	}
	return final
}

func dfsAdd(num int, mp [][]int, ans *[][]bool) {
	for i := 0; i < len(mp[num]); i++ {
		(*ans)[num][mp[num][i]] = true
		dfsAdd(mp[num][i], mp, ans)
	}
	return
}

func dfsQuery(begin int, end int, attr [][]int) bool {
	if begin == end {
		return true
	}
	for i := 0; i < len(attr[begin]); i++ {
		if dfsQuery(attr[begin][i], end, attr) {
			return true
		}
	}
	return false
}

func main() {
	//a := []int{2,4,3,5}
	//b := []int{2,5,4,3}
	//fmt.Println(canBeEqual(a,b))
	//hasAllCodes("2342",3)
	pre := [][]int{{1, 2}, {1, 3}}
	checkIfPrerequisite(3, pre, nil)
}
