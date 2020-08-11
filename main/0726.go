package main

func restoreString(s string, indices []int) string {
	ans := make([]uint8, len(s))
	for i := 0; i < len(indices); i++ {
		ans[indices[i]] = s[i]
	}
	return string(ans)
}

func minFlips(target string) int {
	count := 0
	flag := "0"
	for i := 0; i < len(target); i++ {
		if target[i] != flag[0] {
			count++
			if flag == "0" {
				flag = "1"
			} else {
				flag = "0"
			}
		}
	}
	return count
}

func getLengthOfOptimalCompression(s string, k int) int {
	var ch []uint8
	var count []int
	before := s[0] + 1
	for i := 0; i < len(s); i++ {
		if s[i] != before {
			ch = append(ch, s[i])
			count = append(count, 1)
			before = s[i]
		} else {
			count[len(count)-1]++
		}
	}
	//fmt.Println(ch)
	//fmt.Println(count)
	delta := 0
	for k > 0 && delta < len(s) {
		delta++
		for i := 0; i < len(ch); i++ {
			if count[i] == delta {
				k -= count[i]
				ch = append(ch[:i], ch[i+1:]...)
				count = append(count[:i], count[i+1:]...)
			}
		}
	}
	return 0
}

func main() {
	getLengthOfOptimalCompression("aaabcccd", 2)
}
