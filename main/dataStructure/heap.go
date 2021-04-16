package main

import (
	"fmt"
	"sort"
)

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
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
//func maxEvents(events [][]int) int {
//	leng := len(events)
//	count := 0
//	max := 0
//	h := &IntHeap{}
//	in := make([][]int, 100009)
//	out := make([][]int, 100009)
//	mp := map[int]int{}
//	for i := 0; i < leng; i++ {
//		if events[i][1] > max {
//			max = events[i][1]
//		}
//		in[events[i][0]] = append(in[events[i][0]], i)
//		out[events[i][1]+1] = append(out[events[i][1]+1], i)
//	}
//	for day := 1; day <= max; day++ {
//		for i := 0; i < len(in[day]); i++ {
//			heap.Push(h, events[in[day][i]])
//		}
//		sort.Sort(mp)
//		//for i := 0; i < len(out[day]); i++ {
//		//	heap.Remove(h,events[out[day][i]])
//		//}
//	}
//
//	return count
//}

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

type SortedStack struct {
	heap []int
	len  int
}

func Constructor() SortedStack {
	return SortedStack{
		heap: []int{},
		len:  0,
	}
}

func (this *SortedStack) Push(val int) {
	this.heap = append(this.heap, val)
	i := len(this.heap) - 1
	for i != 0 {
		pi := (i - 1) / 2
		if this.heap[pi] <= val {
			break
		}
		//把父亲节点的值下放并继续往上层迭代
		this.heap[i] = this.heap[pi]
		i = pi
	}
	this.heap[i] = val
	this.len++
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	if this.len == 1 {
		this.heap = make([]int, 0)
	} else {
		// swap 把最底层元素放到堆顶并逐渐往下调整
		this.heap[0], this.heap[this.len-1] = this.heap[this.len-1], this.heap[0]
		this.heap = this.heap[:this.len-1]
		// down
		i := 0
		for {
			j := i*2 + 1
			if j >= len(this.heap) || j < 0 {
				break
			}
			j1 := j      // left child
			j2 := j1 + 1 // right child
			if j2 < len(this.heap) && this.heap[j2] < this.heap[j1] {
				j = j2
			}
			// 找到j定义为左右儿子的min值并 与头节点i交换
			if this.heap[i] <= this.heap[j] {
				break
			}
			this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
			i = j
		}
	}
	this.len--
}

func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.heap[0]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.heap) == 0
}

/**
测试结果:[null,-1,-1,null,true,-1,null,null,null,19,null,19,null,null,null,false,null,8,false,null,8,8,false,null,false,8,false,null,42,null,null,null,false,null,false,63,null,null,null,null,false,null,null,null,17,false,null,52,null,null,6,false,false,false,false,false,null,null,null,null,null,2,2,false,null]
期望结果:[null,-1,-1,null,true,-1,null,null,null,19,null,19,null,null,null,false,null,8,false,null,8,8,false,null,false,8,false,null,25,null,null,null,false,null,false,42,null,null,null,null,false,null,null,null,17,false,null,52,null,null,6,false,false,false,false,false,null,null,null,null,null,2,2,false,null
*/
func main() {
	//s := [][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 5}, {1, 6}, {1, 7}}
	//fmt.Println(maxEvents(s))
	//h := &IntHeap{2,5,6,8,3}
	//heap.Init(h)
	//fmt.Printf("%d",heap.Pop(h))
	//fmt.Printf("%d",heap.Pop(h))
	//fmt.Print(h.Len())
	//heap.Push(h,1)
	//fmt.Printf("%d",heap.Pop(h))
	obj := Constructor()
	obj.Push(40)
	obj.Push(19)
	obj.Push(44)
	obj.Push(8)
	obj.Push(42)
	obj.Push(29)
	fmt.Println(obj.Peek())
	obj.Pop()
	fmt.Println(obj.Peek())

}
