package main

import (
	"fmt"
)

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

var INF = 100000

// 点修改 求区间最小值
//p v  表示设a[p]=v l,r,p 都是a数组坐标,index是node的存储实际数组坐标
func update(l int, r int, index int, p int, v int, nodes *[]int) {
	m := (r + l) / 2
	if l == r {
		(*nodes)[index] = v
	} else {
		if p <= m {
			update(l, m, index*2, p, v, nodes)
		} else {
			update(m+1, r, index*2+1, p, v, nodes)
		}
		(*nodes)[index] = min((*nodes)[index*2], (*nodes)[index*2+1])
	}

}

func query(l int, r int, index int, ql int, qr int, nodes *[]int) int {
	m := (r + l) / 2
	ans := INF
	if ql <= l && qr >= r {
		return (*nodes)[index]
	}
	if ql <= m {
		ans = min(ans, query(l, m, index*2, ql, qr, nodes))
	}
	if qr > m {
		ans = min(ans, query(m+1, r, index*2+1, ql, qr, nodes))
	}
	return ans
}

//区间修改 一段区间加上一个值，求区间的和，最大值，最小值
var addv = make([]int, 50)
var sumv = make([]int, 50)
var minv = make([]int, 50)
var maxv = make([]int, 50)

func updateRange(l int, r int, index int, v int, ql int, qr int) {
	lc := index * 2
	rc := index*2 + 1
	if ql <= l && qr >= r {
		addv[index] += v
	} else {
		m := (l + r) / 2
		if ql <= m {
			updateRange(l, m, lc, v, ql, qr)
		}
		if qr > m {
			updateRange(m+1, r, rc, v, ql, qr)
		}
	}
	// maintain
	sumv[index], minv[index], maxv[index] = 0, 0, 0
	if r > l {
		sumv[index] = sumv[lc] + sumv[rc]
		minv[index] = min(minv[lc], minv[rc])
		maxv[index] = max(maxv[lc], maxv[rc])
	}
	if addv[index] > 0 {
		sumv[index] = addv[index] * (r - l + 1)
		maxv[index] += addv[index]
		minv[index] += addv[index]
	}
}

func queryRange(l int, r int, index int, add int, ql int, qr int) (s, mi, ma int) {
	if ql <= l && qr >= r {
		s += sumv[index] + add*(r-l+1)
		mi = min(mi, minv[index]+add)
		ma = max(ma, maxv[index]+add)
	} else {
		m := (l + r) / 2
		if ql <= m {
			s2, mi2, ma2 := queryRange(l, r, index*2, addv[index], ql, qr)
			s += s2
			mi = min(mi, mi2)
			ma = max(ma, ma2)
		}
		if qr > m {
			s2, mi2, ma2 := queryRange(l, r, index*2+1, addv[index], ql, qr)
			s += s2
			mi = min(mi, mi2)
			ma = max(ma, ma2)
		}
	}
	return s, mi, ma
}

//http://blog.csdn.net/zhulei19931019/article/details/38706259 线段树模板（刘汝佳）
func main() {
	// 点修改 求区间最小值
	num := make([]int, 50)
	update(1, 10, 1, 1, 1, &num)
	update(1, 10, 1, 2, 2, &num)
	update(1, 10, 1, 3, 3, &num)
	update(1, 10, 1, 4, 4, &num)
	update(1, 10, 1, 6, 6, &num)
	update(1, 10, 1, 7, 7, &num)
	update(1, 10, 1, 8, 8, &num)
	update(1, 10, 1, 9, 9, &num)
	update(1, 10, 1, 10, 10, &num)
	fmt.Println(num)
	fmt.Println(query(1, 10, 1, 1, 4, &num))

	// 区间修改

}
