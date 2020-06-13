package main

import (
	"fmt"
	"sort"
)

func shuffle(nums []int, n int) []int {
	var ans []int
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[i+n-1])
	}
	return ans
}

type Info struct {
	value int
	index int
	flag  int
}

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	n := len(arr)
	mid := arr[(n-1)/2]
	delta := make([]Info, n)
	for i := 0; i < n; i++ {
		if arr[i] > mid {
			delta[i].value = arr[i] - mid
			delta[i].flag = 1
		} else {
			delta[i].value = mid - arr[i]
			delta[i].flag = -1
		}
		delta[i].index = i
	}
	sort.Slice(delta, func(i, j int) bool {
		if delta[i].value == delta[j].value {
			return delta[i].flag < delta[j].flag
		}
		return delta[i].value < delta[j].value
	})
	var ans []int
	for i := 1; i <= k; i++ {
		ans = append(ans, arr[delta[n-i].index])
	}
	return ans
}

type BrowserHistory struct {
	cur   int
	pages []string
}

func Constructor(homepage string) BrowserHistory {
	a := BrowserHistory{
		cur:   1,
		pages: []string{homepage},
	}
	return a
}

func (this *BrowserHistory) Visit(url string) {
	if this.cur == len(this.pages) {
		this.pages = append(this.pages, url)
		this.cur++
		return
	}
	this.pages[this.cur] = url
	this.cur++
	var copage []string
	for i := 0; i < this.cur; i++ {
		copage = append(copage, this.pages[i])
	}
	this.pages = copage
}

func (this *BrowserHistory) Back(steps int) string {
	if this.cur-steps <= 0 {
		this.cur = 1
		return this.pages[0]
	}
	this.cur = this.cur - steps
	return this.pages[this.cur-1]
}

func (this *BrowserHistory) Forward(steps int) string {
	leng := len(this.pages)
	if this.cur+steps > leng {
		this.cur = leng
		return this.pages[this.cur-1]
	}
	this.cur = this.cur + steps
	return this.pages[this.cur-1]

}

func main() {
	//arr := []int{6,7,11,7,6,8}
	//fmt.Println(getStrongest(arr,5))
	//["BrowserHistory","visit","forward","back","visit","visit","visit","visit","back","visit","back","forward","visit","visit","visit"]
	//[["hdqqhm.com"],["yltqtsj.com"],[1],[1],["cyv.com"],["vbpvni.com"],["opy.com"],["kbyzp.com"],[7],["fchhxaz.com"],[6],[9],["rg.com"],["oemqn.com"],["hyqsb.com"]]
	obj := Constructor("hdqqhm.com")
	obj.Visit("yltqtsj.com")
	fmt.Println(obj.Forward(1))
	fmt.Println(obj.Back(1))

	obj.Visit("cyv.com")
	obj.Visit("vbpvni.com")
	obj.Visit("opy.com")
	obj.Visit("kbyzp.com")
	fmt.Println(obj.Back(7))

	fmt.Println(obj.Back(1))
	fmt.Println(obj.Forward(1))
}
