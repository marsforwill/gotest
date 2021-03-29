package main

import (
	"fmt"
	"strings"
)

//151. 翻转字符串里的单词
func reverseWords(s string) string {
	var str []uint8
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str = append(str, s[i])
		} else {
			if i+1 < len(s) && s[i+1] != ' ' {
				str = append(str, s[i])
			}
		}
	}
	ansStr := string(str)
	spli := strings.Split(ansStr, " ")
	var ans string
	for i := 0; i < len(spli); i++ {
		if len(spli[i]) == 0 || spli[i] == " " {
			spli = append(spli[0:i], spli[i+1:]...)
		}
	}
	for i := len(spli) - 1; i >= 0; i-- {
		ans = ans + spli[i]
		if i != 0 {
			ans = ans + " "
		}
	}
	return ans
}

//273. 整数转换英文表示 递归 简洁处理细节
//输入：num = 1234567891
//输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	to19 := strings.Split("One Two Three Four Five Six Seven Eight Nine Ten Eleven Twelve Thirteen Fourteen Fifteen Sixteen Seventeen Eighteen Nineteen", " ")
	tens := strings.Split("Twenty Thirty Forty Fifty Sixty Seventy Eighty Ninety", " ")
	var helper func(num int) []string
	helper = func(num int) []string {
		if num <= 0 {
			return []string{}
		}
		if num < 20 {
			return to19[num-1 : num]
		}
		if num < 100 {
			ans := []string{tens[(num/10)-2]}
			return append(ans, helper(num%10)...)
		}
		if num < 1000 {
			ans := []string{to19[num/100-1]}
			ans = append(ans, "Hundred")
			ans = append(ans, helper(num%100)...)
			return ans
		}
		temp := []string{"Thousand", "Million", "Billion"}
		for i := 0; i < len(temp); i++ {
			if num < p(1000, i+2) {
				ansm := helper(num % p(1000, i+1))
				low := make([]string, 30)
				copy(low, ansm)
				ansh := helper(num / p(1000, i+1))
				ansh = append(ansh, temp[i])
				ansh = append(ansh, low...)
				return ansh
			}
		}
		return []string{}
	}
	return strings.TrimRight(strings.Join(helper(num), " "), " ")
}
func p(num int, pow int) int {
	if pow == 0 {
		return 1
	}
	ans := 1
	for i := 0; i < pow; i++ {
		ans = ans * num
	}
	return ans
}

//48. 旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

//365. 水壶问题 能否从x y升的水壶到处z的水 三水杯简化版本 bfs
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	var queue []int
	init := 0
	queue = append(queue, init)
	visited := make(map[int]bool)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = true
		nx := cur / 10000
		ny := cur % 10000
		if nx == z || ny == z || nx+ny == z {
			return true
		}
		// 增加下一次可达的状态：倒满X
		queue = append(queue, x*10000+ny)
		// 增加下一次可达的状态：倒满Y
		queue = append(queue, nx*10000+y)
		// 增加下一次可达的状态：清空X
		queue = append(queue, 0+ny)
		// 增加下一次可达的状态：清空Y
		queue = append(queue, nx*10000+0)
		// x--> y
		if nx > y-ny {
			queue = append(queue, (nx+ny-y)*10000+y)
		} else {
			queue = append(queue, 0, nx+ny)
		}
		// y--> x
		if ny > x-nx {
			queue = append(queue, x*10000+(nx+ny-x))
		} else {
			queue = append(queue, (nx+ny)*10000)
		}
	}
	return false
}
func main() {
	//println(reverseWords("  hello world  "))
	fmt.Println(numberToWords(1234567891))
}
