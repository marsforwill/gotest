package main

import (
	"fmt"
	"sort"
	"strings"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) <= 2 {
		return true
	}
	delta := arr[0] - arr[1]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] != delta {
			return false
		}
	}
	return true
}

func getLastMoment(n int, left []int, right []int) int {
	l := 0
	for i := 0; i < len(left); i++ {
		if left[i] > l {
			l = left[i]
		}
	}
	r := n
	for i := 0; i < len(right); i++ {
		if right[i] < r {
			r = right[i]
		}
	}
	if l > n-r {
		return l
	} else {
		return n - r
	}
}

func numSubmat(mat [][]int) int {
	var m = len(mat)
	var n = len(mat[0])
	var res = 0

	var f [][]int
	for i := 0; i < m; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		f = append(f, tmp)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				f[i][j] = 0
			} else if i == 0 {
				if j == 0 {
					f[i][j] = mat[i][j]
				} else {
					f[i][j] = f[i][j-1] + 1
				}
			} else if j == 0 {
				f[i][j] = f[i-1][j] + 1
			} else {
				last := f[i][j-1] + f[i-1][j] - f[i-1][j-1]
				if last > 0 {
					f[i][j] = last + 1
				} else {
					f[i][j] = 1
				}
			}
			res = res + f[i][j]
		}
	}
	return res
}

func main() {
	//n := getLastMoment(4,[]int{4,3},[]int{0,1})
	//fmt.Println(n)
	//fmt.Println(numSubmat([][]int{{0,0,0},{0,0,0},{0,1,1},{1,1,0},{0,1,1}}))
	//u,err := url.Parse("internal-proxy.proxy:10002")
	//fmt.Println(u)
	//fmt.Println(err)
	fmt.Println(strings.ReplaceAll("/argus/operation/v1/db/static_group", "/operation", ""))
}
