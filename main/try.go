package main

import (
	"fmt"
	"strings"
)

// 输入：events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
//输出：4

// 在每一个时间点 维护一个这个时间点可以参加的会议集合
// 每遍历一个时间点 集合入in 出out

type Event struct {
	start int
	end   int
}
type PairList []Event

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].end < p[j].end }

func maxEvents3(events [][]int) int {
	leng := len(events)
	var list []Event
	in := make([][]int, 100009)
	out := make([][]int, 100009)
	mp := map[Event]int{} // event 对应的数量
	maxDay := 0
	count := 0
	for i := 0; i < leng; i++ {
		list = append(list, Event{
			start: events[i][0],
			end:   events[i][1],
		})
		if events[i][1] > maxDay {
			maxDay = events[i][1]
		}
		in[events[i][0]] = append(in[events[i][0]], i)
		out[events[i][1]+1] = append(out[events[i][1]+1], i)
	}
	for day := 1; day <= maxDay; day++ {
		for i := 0; i < len(in[day]); i++ {
			event := list[in[day][i]]
			c, ok := mp[event]
			if ok {
				mp[event] = c + 1
			} else {
				mp[event] = 1
			}
		}
		for i := 0; i < len(out[day]); i++ {
			event := list[out[day][i]]
			delete(mp, event)
		}
		if len(mp) > 0 {
			count++
			temp := Event{
				start: 100009,
				end:   100009,
			}
			for e := range mp {
				if temp.end > e.end {
					temp = e
				}
			}
			c, _ := mp[temp]
			if c == 1 {
				delete(mp, temp)
			} else {
				mp[temp] = c - 1
			}

		}
	}
	return count

}

//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
func oneEditAway(first string, second string) bool {
	if strings.Compare(first, second) == 0 {
		return true
	}
	l := len(first)
	r := len(second)
	if l-r >= 2 || r-l >= 2 {
		return false
	}
	if l == r {
		count := 0
		for i := 0; i < l; i++ {
			if first[i] != second[i] {
				count++
			}
			if count >= 2 {
				return false
			}
		}
		return true
	}
	var m int
	var long, short string
	if l > r {
		m = l
		long = first
		short = second
	} else {
		m = r
		long = second
		short = first
	}
	for i := 0; i < m; i++ {
		if i == m-1 {
			return true
		}
		if long[i] != short[i] {
			long = long[:i] + long[i+1:]
			break
		}
	}
	return strings.Compare(long, short) == 0

}

//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]

// 0,1 -> 1,3   1,3 -> 3,2 1,2 2,2
func rotate(matrix [][]int) {
	l := len(matrix)
	if l <= 1 {
		return
	}
	temp := make([][]int, l)
	for i := 0; i < l; i++ {
		temp[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			temp[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			matrix[j][l-1-i] = temp[i][j]
		}
	}
	return
}

// dont want write!!!!
type StackOfPlates struct {
	cap   int
	stack [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		cap:   cap,
		stack: make([][]int, 0),
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	if len(this.stack) == 0 {
		newStack := []int{val}
		this.stack = append(this.stack, newStack)
		return
	}
	last := this.stack[len(this.stack)-1]
	if len(last) == this.cap {
		newStack := []int{val}
		this.stack = append(this.stack, newStack)
		return
	}
	last = append(last, val)
	this.stack[len(this.stack)-1] = last
}

func (this *StackOfPlates) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	plate := this.stack[len(this.stack)-1]
	v := plate[len(plate)-1]
	plate = plate[0 : len(plate)-1]
	this.stack[len(this.stack)-1] = plate
	if len(plate) == 0 {
		this.stack = this.stack[0 : len(this.stack)-1]
	}
	return v
}

func (this *StackOfPlates) PopAt(index int) int {
	n := len(this.stack)
	if index >= n {
		return -1
	}
	plate := this.stack[index]
	v := plate[len(plate)-1]
	plate = plate[0 : len(plate)-1]
	this.stack[index] = plate
	if len(plate) == 0 {
		tmp := this.stack[index+1:]
		this.stack = this.stack[:index]
		this.stack = append(this.stack, tmp...)
	}
	return v
}
func main() {
	//s := [][]int{{2, 2,3}, {3, 3,3}}
	//fmt.Println(maxEvents3(s))
	//fmt.Println(oneEditAway("asdf", "asdff"))
	//rotate([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//for i := 0; i < 20; i++ {
	//	fmt.Printf("%v %v\n", i, i*i)
	//}
	//a:=[]int{1,2,3,4}
	//fmt.Println(append(a[:2],a[3:]...))
	//["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
	//[[1], [1], [2], [1], [], []]
	c := Constructor(1)
	c.Push(1)
	c.Push(2)
	fmt.Println(c.PopAt(1))
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
}
