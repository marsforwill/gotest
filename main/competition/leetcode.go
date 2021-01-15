package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type A struct {
	a int
}

// interesting
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i, ni := range nums {
		for j := i; j < l; j++ {
			if ni+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
func gameOfLife(board [][]int) {
	//fmt.Printf("%v\n",board)
	n := len(board)    // row
	m := len(board[0]) // column
	//ans := make([][]int, n,m)
	ans := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	ans[i] = make([]int, m)
	//}
	copy(ans, board)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			countLive := 0
			if i > 0 && j > 0 && board[i-1][j-1] == 1 {
				countLive++
			}
			if i > 0 && board[i-1][j] == 1 {
				countLive++
			}
			if i > 0 && j < m-1 && board[i-1][j+1] == 1 {
				countLive++
			}
			if j > 0 && board[i][j-1] == 1 {
				countLive++
			}
			if j < m-1 && board[i][j+1] == 1 {
				countLive++
			}
			if i < n-1 && j > 0 && board[i+1][j-1] == 1 {
				countLive++
			}
			if i < n-1 && board[i+1][j] == 1 {
				countLive++
			}
			if i < n-1 && j < m-1 && board[i+1][j+1] == 1 {
				countLive++
			}

			if board[i][j] == 1 {
				if countLive < 2 || countLive > 3 {
					ans[i][j] = 0
				} else {
					ans[i][j] = 1
				}
			} else {
				if countLive == 3 {
					ans[i][j] = 1
				} else {
					ans[i][j] = 0
				}

			}
		}
	}
	copy(board, ans)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num, flag = 0, 0
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := ans
	for l1 != nil || l2 != nil || flag > 0 {
		num = 0
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		num += flag
		if num >= 10 {
			num -= 10
			flag = 1
		} else {
			flag = 0
		}
		node := &ListNode{
			Val:  num,
			Next: nil,
		}
		temp.Next = node
		temp = temp.Next
	}
	return ans.Next
}

// EncodeToString returns the base64 encoding of src.

//以从上往下、从左到右进行 Z 字形排列
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	dict := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		dict[i] = make([]byte, len(s))
	}
	i, j := 0, 0
	flag := true
	for index := 0; index < len(s); index++ {
		dict[i][j] = s[index]
		if i+1 == numRows {
			flag = false
		}
		if i == 0 {
			flag = true
		}
		if flag {
			i++
		} else {
			i--
			j++
		}
	}
	var ans string
	for i := 0; i < len(dict); i++ {
		for j := 0; j < len(dict[i]); j++ {
			if dict[i][j] != ' ' && dict[i][j] != '\u0000' {
				ans += string(dict[i][j])
			}
		}
	}
	return ans
}

func myAtoi(s string) int {
	str := strings.TrimSpace(s)
	ans := ""
	i := 0
	flag := false
	for i = 0; i < len(str); i++ {
		if (str[i] == '-' || str[i] == '+') && flag == false {
			flag = true
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			ans = str[:i]
			break
		}
	}
	if i == len(str) {
		ans = str
	}
	if ans[len(ans)-1] == '-' || ans[len(ans)-1] == '+' {
		ans = ans[:len(ans)-1]
	}
	if len(ans) == 0 {
		return 0
	}
	num, _ := strconv.Atoi(ans)
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		return math.MinInt32
	}
	return num
}

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	// dont want to write more
	i, n := 0, len(words)
	for {
		start := i   // 这一行文本在 words 的起始位置
		wordLen := 0 // 统计这一行单词字符的个数
		for ; i < n && wordLen+len(words[i])+i-start <= maxWidth; i++ {
			wordLen += len(words[i])
		}
		// 循环结束后，i 为这一行文本在 words 的结束位置 +1
		if i == n { // 最后一行
			s := strings.Join(words[start:], " ")
			ans = append(ans, s+strings.Repeat(" ", maxWidth-len(s)))
			return ans
		}
		space := maxWidth - wordLen // 需要插入至单词间的空格个数
		if i-start == 1 {           // 只有一个单词的情况，在末尾加上空格即可
			ans = append(ans, words[start]+strings.Repeat(" ", space))
		} else { // 计算两单词间至少要加多少个空格，以及有多少个间隔要额外加一个空格
			avg, extra := space/(i-start-1), space%(i-start-1)
			avgSpace := strings.Repeat(" ", avg)
			s1 := strings.Join(words[start:start+extra+1], avgSpace+" ")
			s2 := strings.Join(words[start+extra+1:i], avgSpace)
			ans = append(ans, s1+avgSpace+s2)
		}
	}
}

func intToRoman(num int) string {
	ans := ""
	if num/1000 > 0 {
		n := num / 1000
		ans = ans + strings.Repeat("M", n)
	}
	if (num/100)%10 > 0 {
		n := (num / 100) % 10
		if n == 4 {
			ans = ans + "CD"
		} else if n == 9 {
			ans = ans + "CM"
		} else {
			a := n % 5
			b := 0
			if n >= 5 {
				b = 1
			}
			ans = ans + strings.Repeat("D", b) + strings.Repeat("C", a)
		}
	}
	if (num/10)%10 > 0 {
		n := (num / 10) % 10
		if n == 4 {
			ans = ans + "XL"
		} else if n == 9 {
			ans = ans + "XC"
		} else {
			a := n % 5
			b := 0
			if n >= 5 {
				b = 1
			}
			ans = ans + strings.Repeat("L", b) + strings.Repeat("X", a)
		}
	}
	if num%10 > 0 {
		n := num % 10
		if n == 4 {
			ans = ans + "IV"
		} else if n == 9 {
			ans = ans + "IX"
		} else {
			a := n % 5
			b := 0
			if n >= 5 {
				b = 1
			}
			ans = ans + strings.Repeat("V", b) + strings.Repeat("I", a)
		}
	}
	return ans
}

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			m[nums[i]] = true
		}
	}
	ans := 1
	for ; ans <= len(nums); ans++ {
		if _, ok := m[ans]; !ok {
			break
		}
	}
	return ans
}
func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	//fmt.Println(intToRoman(1994))
	//fmt.Println(myAtoi("123-"))
	//a := ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//b := ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//addTwoNumbers(&a, &b)

	//str := "U09GMbdfAAAQBAAAAAEAAAAAAAAEAAAAAAAAAAAAAABbCr0Jkud1uW3Bk1P6E6BKrdXgXL3lqFk42fhK/6uDjyPSHGNdfI0iEfdC4qgev0hEKkknGO4X4f1DEOjfSYe51sAIh8OUTi9Z4+xQrwspYxhCOEhGaqetrkrrfHqtlXGpm0UkWbhnhSpf1T0QwXJRxP68xqGS1gjdCC17mL0C5be2VaPBAWQ2fJVQVaEYpmiC2mua4MiFdgna3O26SyngcAlKiIe/xDtVNDxNDSnZvGiAyeL77S3iEy5J/WQeoS61z8LEOUwSeHHJovrkRkeiUugAUF4Vp2iAPYwkWr52srC46l0b2XynhZ2aUCog+IBlWWGHeVMpmI2VLm1R3KZj7d8jY93I9U5K0LkOtaqGrI4w2NDjGcIMpigxTs5lbVwnitIuJflhVPFzjsmoHYRK9tEKpMvNt5Dcl4ItyPQQyV0K5y8AkLVO7uECMcZ13BC6oK4eaNAleHKaB/+ZPu1XcgFe4VWc2AFnqJBiZrS9HlF9RPp1PgtwLn8EwqYaKTn0gk9Jhp/313Qv6oanSuJPB0LEzC8Kn1OPy5BvV5aWp0iIxj9ilgAd+7dK/lEwEhVVPGVVW1GIhcHp61xqS7xWSkxymdWi+lW+0q98Loj9QUcjLC/MjY4/HlBcU8M9JG/zVeKBQUGO2It5uxOwuHwGJMHWmfgYGpXs4vZtKgNXU4Wz2FG0yVCNpuTcv61o+9b2dWwV7ZxxgGgsHhXslcTKLTs6irDL3slngCRi2jCGgMWL/jfueafiwXHLvQnvZWCMZUK3gUm0TIoAFxgtmHMvw5rULE83KVBGW3IQcCVNz9RE+Jy972dTOnIYbDWE3g7IfWnl2hmh+q4SWJupZ/qT3FV/VGLJP7lhFQIEzUeLq1RKiyLsYHFOlgZFXjypxptDk8damvAVaCqVx58WEVdgeqAWXPfN2e7yN40eOZhpnLwEgu/JNeaKpo0OPCeGfR/Dh48rizbYAIRGeg3t5yPMX+x0xSh+YK6e5gy04t+wfJOJUx6JPaLbOTZ3+sYAS7xSoq3CNPsqpIUYZQLSOErkzfGpDudz4iGjifXJsec6pRq5mRHtKsE2h4TVIyppttKn7gtcIz12ZoG+TRiXh+j+eX1k8pqaFNppy2qCewzNPMuVVnLOlr/Nc7UlmmFe+w4j2xr21FZsXC7Hkl7DfiIEU8YWULxMe2lHzs7um/foL1as/PXd/5ZFQh/4PehKnLpVZ6h6gLg/hmEDclCljudPoVKeJzblDMTUKCtzEfDf/v2BHl7Xcqq5tH+TQ78GuGCyu+PziCNdhSGfnd70E4sDBJMlg6QeA/iEX7wF/BzRJ5GlIruddx8q8iMJ/fTBssysWGEBIABziJQ9mslaEKuWQ0q4fXUIpPrOwhJDKxmwmA=="
	//bytes :=[]byte(str)
	//fmt.Println(bytes)
	//str2 := string(bytes)
	//fmt.Println(str2)
	//bytes2,err := base64.StdEncoding.DecodeString(str)
	//if err != nil{
	//	fmt.Printf(err.Error())
	//}
	//fmt.Println(bytes2)
	//str3 := base64.StdEncoding.EncodeToString(bytes2)
	//fmt.Println(str3)

	//var testCase = [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	//gameOfLife(testCase)

	//
	//var ages []int = []int{1, 2, 3, 5}
	//
	//fmt.Println(ages)
	//b := test(ages)
	//fmt.Println(ages)
	//fmt.Println(b)

	//var ages int =5
	//
	//fmt.Println(ages)
	//b := test2(ages)
	//fmt.Println(ages)
	//fmt.Println(b)

}

// 参数指针传递
func test(ages []int) []int {
	var b = ages
	b[0]++
	return b
}

// 参数指针传递
func test2(ages int) int {
	var b = ages
	b++
	return b
}
