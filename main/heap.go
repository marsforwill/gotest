package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap [][]int

func (h IntHeap)Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1]{
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 按事件扫描/按天扫描

// 输入：events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
//输出：4

// 在每一个时间点 维护一个这个时间点可以参加的会议集合
// 每遍历一个时间点 集合入in 出out
func maxEvents(events [][]int) int {
	leng := len(events)
	count := 0
	max := 0
	h := &IntHeap{}
	in := make([][]int,100009)
	out := make([][]int,100009)
	mp := map[int]int{}
	for i := 0; i < leng; i++ {
		if events[i][1] > max {
			max = events[i][1]
		}
		in[events[i][0]] = append(in[events[i][0]], i)
		out[events[i][1]+1] = append(out[events[i][1]+1],i)
	}
	for day := 1; day <= max ; day++ {
		for i := 0; i < len(in[day]); i++ {
			heap.Push(h,events[in[day][i]])
		}
		sort.Sort(mp)
		//for i := 0; i < len(out[day]); i++ {
		//	heap.Remove(h,events[out[day][i]])
		//}
	}

	return count
}

func maxEvents2(events [][]int) int {

	eLen := len(events)
	if eLen <= 1 || eLen >= 99999 {
		return eLen
	}

	// minDay := events[0][0]
	sort.Slice(events, func(i int, j int) bool {
		return events[i][1] < events[j][1]
	})
	// fmt.Println(minDay)
	maxDay := events[eLen-1][1]

	days := make([]bool, maxDay+1)
	maxEvent := 0
	for i := 0; i < eLen; i++ {
		// 极端case需要把这步降低到logN
		for j := events[i][0]; j <= events[i][1]; j++ {
			if !days[j] { //每个会议只占用一天
				days[j] = true
				maxEvent++
				break
			}
		}
	}

	return maxEvent

}

func main() {
	s := [][]int{{1,5},{1,5},{1,5},{2,3},{2,5},{1,6},{1,7}}
	fmt.Println(maxEvents(s))
	//h := &IntHeap{2,5,6,8,3}
	//heap.Init(h)
	//fmt.Printf("%d",heap.Pop(h))
	//fmt.Printf("%d",heap.Pop(h))
	//fmt.Print(h.Len())
	//heap.Push(h,1)
	//fmt.Printf("%d",heap.Pop(h))
}
