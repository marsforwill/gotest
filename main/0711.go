package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	l := n
	sum := make([]int,l)
	count := 0
	for i := 0; i < l; i++ {
		count += nums[i]
		sum[i] = count
	}
	var temp []int
	for i := 0; i < l; i++ {
		for j := 0; j< i; j++ {
			temp = append(temp, (sum[i] - sum[j])%mod)
		}
	}
	for i := 0; i < n; i++ {
		temp = append(temp,sum[i]%mod)
	}
	sort.Ints(temp)
	ans := 0
	for i := left-1; i <= right-1 ; i++ {
		ans += temp[i]
	}
	return ans % mod
}

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	ans := 1000000007
	if nums[n-4] - nums[0] < ans {
		ans =  nums[n-4] - nums[0]
	}
	if nums[n-3] - nums[1] < ans {
		ans =  nums[n-3] - nums[1]
	}
	if nums[n-2] - nums[2] < ans {
		ans =  nums[n-2] - nums[2]
	}
	if nums[n-1] - nums[3] < ans {
		ans =  nums[n-1] - nums[3]
	}
	return ans

}
func winnerSquareGame(n int) bool {
	flag := make([]bool, n+1)
	for i := 1; i*i <= n; i++ {
		flag[i*i] = true
	}
	if flag[n] == true {
		return true
	}
	for i := 2; i <= n ; i++ {
		if flag[i] {
			continue
		}
		f := false
		for num :=1;num*num < i;num++ {
			if flag[i-num*num] == false {
				flag[i] = true
				f = true
				break
			}
		}
		if f == false {
			flag[i] = false
		}
	}
	return flag[n]
}
func main() {
	//fmt.Println(rangeSum([]int{1,2,3,4},4,1,5))
	fmt.Println(winnerSquareGame(15))
}
