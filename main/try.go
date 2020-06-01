package main

import "fmt"

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
func main() {
	s := [][]int{{2, 2}, {3, 3}, {4, 4}, {1, 5}, {1, 5}}
	fmt.Println(maxEvents3(s))
}
