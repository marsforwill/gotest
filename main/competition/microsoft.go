package main

import "strings"

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

/**
def numberToWords(self, num: int) -> str:
        to19 = 'One Two Three Four Five Six Seven Eight Nine Ten Eleven Twelve ' \
               'Thirteen Fourteen Fifteen Sixteen Seventeen Eighteen Nineteen'.split()
        tens = 'Twenty Thirty Forty Fifty Sixty Seventy Eighty Ninety'.split()

        def helper(num):
            if num < 20:
                return to19[num - 1:num]
            if num < 100:
                return [tens[num // 10 - 2]] + helper(num % 10)
            if num < 1000:
                return [to19[num // 100 - 1]] + ["Hundred"] + helper(num % 100)
            for p, w in enumerate(["Thousand", "Million", "Billion"], 1):
                if num < 1000 ** (p + 1):
                    return helper(num // 1000 ** p) + [w] + helper(num % 1000 ** p)

        return " ".join(helper(num)) or "Zero"

*/
//273. 整数转换英文表示 递归 简洁处理细节
//输入：num = 1234567891
//输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
func numberToWords(num int) string {
	to19 := strings.Split("One Two Three Four Five Six Seven Eight Nine Ten Eleven Twelve Thirteen Fourteen Fifteen Sixteen Seventeen Eighteen Nineteen", " ")
	tens := strings.Split("Twenty Thirty Forty Fifty Sixty Seventy Eighty Ninety", " ")
	var helper func(num int) []string
	helper = func(num int) []string {
		if num < 20 {
			return to19[num-1:num]
		}
		if num < 100 {
			ans := []string{tens[(num/10)-2]}
			return append(ans, helper(num%10)...)
		}
		if num < 1000 {
			ans := []string{to19[num/100-1]}
			ans = append(ans,"Hundred")
			ans = append(ans, helper(num%100)...)
			return ans
		}
		temp := []string{"Thousand", "Million", "Billion"}
		for i := 0; i < len(temp); i++ {
			if num < 1000*(i+1) {
				return helper()
			}
		}
	}
	return strings.Join(helper(num), " ")
}

func main() {
	println(reverseWords("  hello world  "))
}
