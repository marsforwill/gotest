package main

import "fmt"

//https://blog.csdn.net/helloworldchina/article/details/104465772

// 计算pattern串的next数组 最大公共前后缀的长度值 前者是在 pat 中匹配 pat[1..end]
func kmpNext(str string) []int {
	next := make([]int, len(str))
	next[0] = 0
	j := 0
	for i := 1; i < len(str); i++ {
		for j > 0 && str[j] != str[i] {
			j = next[j-1]
		}
		if str[i] == str[j] {
			j++
		}
		next[i] = j
	}

	return next
}

// 在s中匹配模式串pattern t
func kmp(s, t string, next []int) int {
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != t[j] {
			//当不相等的时候j（t串指针并不是归零而是从next数组【j-1】开始）
			j = next[j-1]
		}
		if s[i] == t[j] {
			j++
		}
		if j == len(t) {
			return i - j + 1
		}
	}
	return -1
}

func main() {
	s := "abababc"
	t := "ababc"
	next := kmpNext(t)
	fmt.Println(next)
	print(kmp(s, t, next))
}
