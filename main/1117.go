package main

// 加两个数组存map
func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	ans := 0
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := A[i] + B[j]
			v, ok := m[sum]
			if ok {
				m[sum] = v + 1
			} else {
				m[sum] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := C[i] + D[j]
			v, ok := m[0-sum]
			if ok {
				ans += v
			}
		}
	}
	return ans
}

func main() {

}
