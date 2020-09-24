package main


// 计算pattern串的next数组
func kmpNext(str string) []int  {
	next := make([]int, len(str))
	next[0] = 0
	k :=0
	for j := 1; j < len(str); j++ {
		for k > 0 && str[k] != str[j] {
			 k = next[k-1]
		}
		if str[j] == str[k] {
			k++
		}
		next[j] = k
	}

	return next
}

func kmp(s,t string, next []int) int {
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != t[j] {
			j = next[j-1]
		}
		if s[i] == t[j] {
			j++
		}
		if j == len(t) {
			return i-j+1
		}
	}
	return -1
}

func main() {
	s := "abababc"
	t := "ababc"
	next := kmpNext(t)
	print(kmp(s,t,next))
}