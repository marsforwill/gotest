package main

func findKthPositive(arr []int, k int) int {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
	}
	for i := 1; i < 1001; i++ {
		_, has := m[i]
		if has {
			continue
		} else {
			k--
		}
		if k == 0 {
			return i
		}
	}
	return 0
}
