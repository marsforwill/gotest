package main

import "fmt"

func xorOperation(n int, start int) int {
	num := make([]int, n)
	for i := 0; i < n; i++ {
		num[i] = start + 2*i
	}
	ans := num[0]
	for i := 1; i < n; i++ {
		ans = ans ^ num[i]
	}
	return ans
}

func getFolderNames(names []string) []string {
	leng := len(names)
	mapi := make(map[string]int)
	ans := make([]string, leng)
	for i := 0; i < leng; i++ {
		count, ok := mapi[names[i]]
		if ok {
			flag := true
			for flag == true {
				name := fmt.Sprintf("%v(%v)", names[i], count)
				_, ok := mapi[name]
				if ok {
					count++
				} else {
					ans[i] = name
					mapi[name] = 1
					mapi[names[i]] = count + 1
					flag = false
					break
				}
			}
		} else {
			ans[i] = names[i]
			mapi[names[i]] = 1
		}
	}
	return ans
}

// 第i天抽的水 贪心算 一定抽 有水&&后面排得靠前
//
func avoidFlood(rains []int) []int {
	l := len(rains)
	ans := make([]int, l)
	mapi := make(map[int]int)
	var zero []int
	for i := 0; i < l; i++ {
		if rains[i] > 0 {
			ans[i] = -1
			last, ok := mapi[rains[i]]
			// 当前下雨的池有水
			if ok {
				if len(zero) == 0 {
					return []int{}
				} else {
					var j int
					flag := false
					for j = 0; j < len(zero); j++ {
						if zero[j] < i && zero[j] > last {
							ans[zero[j]] = rains[i]
							zero = append(zero[:j], zero[j+1:]...)
							flag = true
							break
						}
					}
					if flag == false {
						return []int{}
					}
				}
			}
			mapi[rains[i]] = i
			ans[i] = -1
		} else {
			zero = append(zero, i)
			ans[i] = 1
		}
	}
	return ans
}
func main() {
	//a := []string{"kaido","kaido(1)","kaido","kaido(1)","kaido(2)"}
	//fmt.Println(getFolderNames(a))
	fmt.Println(avoidFlood([]int{0, 1, 1}))

}
