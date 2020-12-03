package main

func maxDepth(s string) int {
	ans := 0
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cur++
		}
		if s[i] == ')' {
			cur--
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func maximalNetworkRank(n int, roads [][]int) int {
	city := make([]int, n)
	for i := 0; i < len(roads); i++ {
		city[roads[i][0]]++
		city[roads[i][1]]++
	}
	m := 0
	for i := 0; i < n; i++ {
		if city[i] > m {
			m = city[i]
		}
	}
	count := 0
	second := 0
	for i := 0; i < n; i++ {
		if city[i] == m {
			count++
		} else {
			if city[i] > second {
				second = city[i]
			}
		}
	}
	if count >= 2 {

	}
	return 0
}

func checkPalindromeFormation(a string, b string) bool {
	if isHuiwen(a) || isHuiwen(b) {
		return true
	}
	l := 0
	r := len(a) - 1
	flag := false
	for l < r {
		if flag == true {
			if a[l] != a[r] && b[l] != b[r] {
				break
			}
		}
		if a[l] != b[r] {
			if a[l] != a[r] && b[l] != b[r] {
				break
			} else {
				flag = true
			}
		}
		l++
		r--
	}
	if l >= r {
		return true
	}
	l = 0
	r = len(a) - 1
	flag = false
	for l < r {
		if flag == true {
			if b[l] != b[r] && a[l] != a[r] {
				break
			}
		}
		if b[l] != a[r] {
			if b[l] != b[r] && a[l] != a[r] {
				break
			} else {
				flag = true
			}
		}
		l++
		r--
	}
	if l >= r {
		return true
	}
	return false
}

func isHuiwen(str string) bool {
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// out of time
func palindromePairs(words []string) [][]int {
	var ans [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if isHuiwen(words[i] + words[j]) {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
