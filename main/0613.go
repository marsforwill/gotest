package main

import (
	"fmt"
	"math"
	"sort"
)

func finalPrices(prices []int) []int {
	l := len(prices)
	var count int
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		count = 0
		for j := i + 1; j < l; j++ {
			if prices[j] < prices[i] {
				count = prices[j]
				break
			}
		}
		ans[i] = prices[i] - count
	}
	return ans
}

type SubrectangleQueries struct {
	rect [][]int
}

//func Constructor(rectangle [][]int) SubrectangleQueries {
//	s := new(SubrectangleQueries)
//	s.rect = rectangle
//	return *s
//}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rect[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rect[row][col]
}

func runningSum(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		ans[i] = sum
	}
	return ans
}
func findLeastNumOfUniqueInts(arr []int, k int) int {
	mapcount := make(map[int]int)
	l := len(arr)
	for i := 0; i < l; i++ {
		c, ok := mapcount[arr[i]]
		if ok {
			mapcount[arr[i]] = c + 1
		} else {
			mapcount[arr[i]] = 1
		}
	}
	var count []int
	for _, v := range mapcount {
		count = append(count, v)
	}
	sort.Ints(count)
	fmt.Println(count)
	for i := 0; i < len(count); i++ {
		k -= count[i]
		if k < 0 {
			return len(count) - i
		}
		if k == 0 {
			return len(count) - i - 1
		}
	}
	return 0
}

//[1,2,3,10,2,5,6,10], m = 3 花, 相邻 k = 2
//[2,3,10,10,5,6,10,
func minDays(bloomDay []int, m int, k int) int {
	l := len(bloomDay)
	if m*k > l {
		return -1
	}
	return 0
}

// 1 <= k <= n <= 5*10^4
// parent[0] == -1 表示编号为 0 的节点是根节点。
// 对于所有的 0 < i < n ，0 <= parent[i] < n 总成立
// 0 <= node < n
// 至多查询 5*10^4 次
type TreeAncestor struct {
	dp [50009][18]int
	n  int
}

func Constructor(n int, parent []int) TreeAncestor {
	// 一维度元素下标示 二唯2的j次方 value是元素下标i的（2的j次方个）祖先的值
	var ans [50009][18]int
	for i := 0; i < 50009; i++ {
		for j := 0; j < 18; j++ {
			ans[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		ans[i][0] = parent[i]
	}
	for i := 1; i < 18; i++ {
		for j := 0; j < n; j++ {
			if ans[j][i-1] == -1 {
				continue
			}
			// point！！
			// 下标j往上找2的i次方的祖先 = 下标j往上找2的（i-1）次方，再往上找2的（i-1）次方
			// 因为 pow(2,i) = pow(2,i-1) + pow(2,i-1)
			ans[j][i] = ans[ans[j][i-1]][i-1]
		}
	}
	return TreeAncestor{
		n:  n,
		dp: ans,
	}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := 0; i < 18; i++ {
		if ((1 << i) & k) > 0 {
			node = this.dp[node][i]
			if node == -1 {
				return node
			}
		}
	}
	return node
}

// 输入：arr = [3,1,1,1,5,1,2,1], target = 3
//输出：3
//解释：注意子数组 [1,2] 和 [2,1] 不能成为一个方案因为它们重叠了。
// 滑动窗口 加 dp
func minSumOfLengths(arr []int, target int) int {
	l := len(arr)
	j := 0
	sum := 0
	var sta, end []int
	// 滑动窗口双指针
	for i := 0; i < l; i++ {
		sum += arr[i]
		if sum == target {
			sta = append(sta, j)
			end = append(end, i)
		}
		for sum >= target {
			sum -= arr[j]
			j++
			if sum == target {
				sta = append(sta, j)
				end = append(end, i)
			}
		}
	}
	lc := len(end)
	if lc < 2 {
		return -1
	}

	// dp求不相交最小
	index := 0
	min := 1 << 24
	//arr[i]的左边最小长度
	left := make([]int, l)
	//arr[i]的右边最小长度
	right := make([]int, l)
	for i := 0; i < l; i++ {
		if index >= len(sta) {
			break
		}
		if i == end[index] {
			min = minNum(min, end[index]-sta[index]+1)
			index++
		}
		left[i] = min
	}
	index = len(sta) - 1
	min = 1 << 24
	for i := l - 1; i >= 0; i-- {
		if index < 0 {
			break
		}
		if i == sta[index] {
			min = minNum(min, end[index]-sta[index]+1)
			index--
		}
		right[i] = min
	}
	ans := 1 << 24
	//吐血调试dp
	for i := 0; i < l-1; i++ {
		ans = minNum(ans, left[i]+right[i+1])
	}
	if ans == 1<<24 {
		return -1
	}
	return ans
}

func minNum(m int, i int) int {
	if m < i {
		return m
	} else {
		return i
	}
}

//
//type TreeAncestor struct {
//	nums   []int
//	parent []int
//	resMap map[int]int
//}
//
//func Constructor(n int, parent []int) TreeAncestor {
//	tmpNums := make([]int, n)
//	tmpParent := make([]int, n)
//	resMap := make(map[int]int, 0)
//	for i := 0; i < n; i++ {
//		tmpNums[i] = i
//		tmpParent[i] = parent[i]
//		resMap[parent[i]] = 1
//	}
//	return TreeAncestor{nums: tmpNums, parent: tmpParent, resMap: resMap}
//}
//
//func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
//	if k > node {
//		return -1
//	}
//	if len(this.resMap) == len(this.parent) {
//		return node - k
//	}
//	for k >= 1 && node >= 0 {
//		node = this.parent[node]
//		k--
//	}
//	return node
//}

func main() {
	//a := []int{487,25,59,16,2,2,1,308,29,254,1,302,150,127,4,6,2,1,346,64,10,74,11,37,39,4,6,1,476,14,60,30,3,7,2,481,58,3,26,10,11,2,1,114,450,7,1,9,8,2,1,162,174,205,26,24,1,381,103,14,73,14,5,1,1,545,10,6,16,3,9,3,147,160,156,121,1,1,2,3,1,333,70,155,3,5,21,5,360,151,35,5,29,7,3,1,1,478,99,2,4,7,1,1,158,372,20,24,16,2,563,7,5,4,11,1,1,280,120,7,141,11,4,10,8,2,1,4,3,1,235,10,136,78,126,7,403,167,20,1,1,156,94,338,2,1,1,6,475,80,13,10,5,2,1,210,208,124,35,15,30,156,256,30,22,20,71,7,308,107,47,14,32,40,11,11,20,1,1,244,338,2,4,3,1,227,51,97,81,23,14,48,15,18,1,3,9,4,1,237,111,41,197,1,1,4,165,178,60,169,6,3,5,6,488,54,48,1,1,491,85,10,4,2,499,32,45,11,3,1,1,327,173,8,6,56,18,4,309,267,2,7,6,1,331,168,33,39,20,1,37,338,166,8,43,461,93,38,369,72,50,15,57,27,2,440,23,82,22,14,5,2,1,1,2,520,7,44,4,1,8,5,3,266,10,277,30,6,2,1,286,68,229,5,3,1,143,369,10,31,31,8,556,14,17,5,383,194,12,2,1,428,53,103,7,1,274,121,99,36,3,7,6,20,6,7,4,2,6,1,550,7,10,8,5,11,1,581,9,2,164,51,127,194,13,4,33,3,2,1,231,174,166,6,10,1,1,3,396,98,19,37,9,13,20,447,8,75,12,3,3,18,9,10,5,1,1,113,71,45,328,29,6,370,43,58,115,5,1,512,49,7,3,3,12,4,1,1,351,48,180,12,1,47,274,176,84,2,3,2,3,1,227,109,191,3,60,2,123,464,3,2,475,17,37,16,27,17,3,192,116,199,3,81,1,255,78,61,196,2,45,540,3,3,1,124,419,12,7,10,7,2,11,381,110,77,15,7,2,1,30,228,185,109,16,5,17,1,413,142,18,13,4,2,237,214,130,7,3,1,475,5,106,2,1,1,1,1,253,126,80,109,12,5,7,516,4,5,14,53,289,193,52,37,4,17,264,140,124,37,11,13,1,2,464,97,18,12,1,540,20,9,15,5,2,1,123,22,439,4,4,320,156,55,26,12,10,4,5,4,266,25,194,101,5,1,155,281,65,69,4,4,2,6,6,116,266,36,11,29,8,122,2,1,1,147,268,59,31,52,11,23,1,298,84,115,25,38,9,18,1,4,171,163,257,1,286,290,6,8,2,309,54,204,10,5,1,6,1,1,1,482,73,7,14,15,1,379,208,5,214,65,137,141,1,22,5,7,537,18,14,12,4,4,2,1,471,65,14,36,6,42,4,325,86,22,98,11,4,105,154,200,67,6,6,44,9,1,263,289,27,13,352,81,56,87,15,1,504,8,58,13,9,185,34,222,133,11,7,328,41,177,33,11,1,1,488,99,2,3,125,240,33,80,114,334,75,13,145,4,5,4,10,2,81,430,39,2,10,8,20,2,205,170,107,38,19,26,14,8,3,1,1,22,171,83,82,28,204,1,1,48,423,77,14,18,6,4,1,1,355,234,2,1,59,203,140,69,37,5,8,7,52,3,2,1,5,1,168,89,7,182,94,23,5,5,12,6,1,501,61,13,16,1,116,452,13,8,2,1,272,313,2,2,1,2,483,52,10,3,14,2,16,9,2,1,169,52,53,145,157,8,5,2,1,143,320,24,57,6,40,1,1,46,519,17,4,4,1,1,354,75,131,17,6,7,2,198,319,10,48,1,1,5,4,2,4,284,127,22,158,1,303,188,25,76,404,55,71,25,5,16,5,1,5,2,1,2,109,45,179,9,18,211,1,14,1,5,20,446,106,14,3,1,2,166,73,295,36,6,9,5,2,563,13,5,10,1,488,35,18,51,254,222,44,21,9,29,8,4,1,129,413,40,8,1,1,570,6,1,13,2,92,152,280,14,13,4,37,456,18,91,14,10,1,2,68,182,136,171,1,5,14,10,4,1,101,422,57,5,5,1,1,27,470,69,10,16,401,150,21,8,10,1,1,370,67,58,27,9,61,393,173,10,8,1,1,6,104,209,234,20,11,4,2,6,2,475,4,102,3,5,3,195,241,150,4,1,1,218,122,113,105,1,9,4,4,3,11,2,369,149,32,25,7,7,1,1,1,271,61,105,61,24,31,14,1,19,5,401,160,4,2,18,2,2,3,74,173,105,7,134,81,1,3,1,4,4,3,2,213,283,76,2,12,2,1,3,539,38,12,1,1,1,53,43,26,208,224,37,1,491,48,51,2,96,261,15,118,94,6,1,1,458,13,89,27,4,1,398,65,91,17,20,1,549,1,7,20,1,7,1,2,3,1,410,42,122,8,9,1,4,100,178,216,65,25,2,1,1,20,406,46,92,7,19,2,592,565,27,327,261,1,2,1,586,1,2,2,1,193,184,143,40,15,3,13,1,131,268,78,26,76,4,8,1,388,38,161,3,2,61,101,42,258,77,4,23,13,10,3,227,85,41,42,101,69,25,1,1,538,3,9,29,4,3,2,3,1,428,45,83,7,28,1,207,10,225,130,19,1,60,480,20,22,9,1,101,402,89,286,58,146,81,10,9,2,464,84,27,9,8,405,128,13,7,17,21,1,555,7,2,28,424,68,4,64,17,9,2,4,2,10,508,61,6,4,1,281,303,4,3,1,362,162,11,45,9,3,186,234,115,43,11,2,1,289,240,60,3,388,37,77,18,49,4,10,7,1,1,253,96,114,97,10,14,2,6,252,61,117,21,95,18,18,1,9,536,44,3,4,5,321,206,10,55,226,146,125,46,27,2,10,5,4,1,171,421,228,131,146,41,23,5,16,1,1,95,186,154,102,27,24,2,2,322,34,223,6,3,3,1,381,14,177,3,12,2,1,2,60,268,264,298,204,79,5,1,4,1,350,163,62,11,1,3,1,1,383,165,14,27,1,1,1,409,131,26,8,16,2,240,151,149,7,19,26,216,301,31,20,2,10,8,2,2,501,57,14,8,6,2,3,1,275,26,118,43,39,21,33,17,7,9,1,1,2,351,168,56,6,7,2,1,1,319,129,69,25,43,2,1,4,585,1,5,1,4,93,118,338,36,2,1,547,32,12,1,111,60,6,334,65,16,255,194,69,65,7,2,167,249,145,17,10,1,3,250,78,119,74,44,7,9,3,5,1,1,1,174,219,73,68,34,8,3,7,6,118,472,1,1,491,93,1,5,2,12,138,388,4,17,8,22,1,1,1,25,523,12,12,9,1,9,1,311,9,206,7,25,30,4,141,138,293,6,7,3,4,112,470,2,8,246,248,9,58,2,6,6,13,3,1,224,101,81,172,9,4,1,457,97,12,4,21,1,184,374,10,13,9,2,486,63,16,4,16,4,2,1,257,39,114,171,3,5,1,1,1,414,134,14,24,6,345,231,1,2,11,2,533,45,1,6,1,3,1,1,1,73,85,145,194,87,8,109,98,161,44,159,2,17,2,335,226,20,7,1,3,289,166,110,9,15,1,1,1,560,17,5,8,2,533,42,1,6,4,3,1,1,1,91,314,109,1,70,4,1,2,329,169,12,22,26,31,3,62,70,352,23,51,1,13,11,5,3,1,187,206,155,38,6,548,19,5,3,13,2,2,579,5,1,6,1,479,51,27,33,1,1,573,8,6,3,1,1,476,79,9,24,4,465,5,4,2,75,2,35,2,1,1,218,258,107,9,26,154,19,392,1,550,36,2,2,2,481,61,29,20,1,522,68,2,80,247,245,13,6,1,256,213,46,52,2,3,11,1,4,2,1,1,116,12,78,265,13,60,46,1,1,112,194,262,6,13,4,1,388,10,84,38,10,29,16,7,10,316,141,117,15,2,1,132,184,56,146,64,8,1,1,424,88,75,1,4,337,60,90,33,31,30,2,6,1,1,1,77,225,63,65,96,63,3,512,42,32,5,1,402,178,7,3,2,214,143,93,120,4,5,1,10,2,44,184,316,40,4,1,1,2,538,20,11,19,1,1,2,89,77,407,9,8,2,487,42,7,13,6,2,27,5,3,342,202,9,18,14,2,1,4,76,142,6,117,142,71,9,21,4,4,302,227,61,1,1,516,74,1,1,308,47,67,133,26,4,5,1,1,344,219,10,5,5,3,5,1,47,143,150,25,183,44,290,146,9,111,4,21,4,1,5,1,499,55,28,4,3,3,401,190,1,199,81,133,84,82,7,1,3,2,522,10,43,8,1,6,2,336,135,88,1,31,1,550,27,1,2,1,11,62,341,43,109,19,6,6,5,1,523,38,22,2,5,2,73,40,195,122,4,33,121,4,234,279,55,1,20,3,326,150,29,14,36,22,15,532,51,1,1,6,1,234,353,1,2,2,231,299,45,13,3,1,424,5,71,10,79,3,558,9,17,4,2,1,1,344,234,11,3,268,322,1,1,62,296,180,20,32,1,1,213,331,22,23,3,364,100,52,55,12,9,129,332,87,37,2,5,79,23,60,227,43,16,34,82,11,14,2,1,116,352,40,48,30,3,3,170,243,96,56,23,3,1,501,45,36,9,1,17,424,25,109,10,4,2,1,478,17,71,20,2,4,35,42,494,18,3,445,38,45,8,10,38,3,5,468,66,1,14,15,5,22,1,166,397,25,1,3,350,90,142,6,3,1,588,4,141,116,11,70,112,59,29,39,12,1,1,1,34,551,6,1,160,268,158,1,3,2,384,174,30,2,1,1,31,230,248,66,11,3,1,2,211,335,15,31,283,250,5,16,18,19,1,483,18,37,18,12,12,8,1,2,1,65,287,61,129,5,28,9,3,4,1,480,27,85,186,177,206,2,12,3,2,3,1,350,221,13,6,2,303,9,238,7,9,7,14,5,373,182,32,3,2,294,244,28,17,2,4,2,1,284,298,1,2,3,3,1,300,286,3,2,1,470,36,34,19,30,2,1,29,173,190,41,88,7,6,15,6,4,29,4,278,218,81,14,1,304,52,217,2,1,8,4,2,2,191,130,243,28,146,113,171,46,76,29,7,2,1,1,568,22,1,1,381,191,14,2,2,1,1,149,183,36,148,48,17,9,2,240,56,260,20,13,3,78,318,1,195,565,22,2,1,2,525,28,24,14,1,280,97,38,31,134,11,1,161,240,159,3,10,2,6,1,3,3,2,1,1,376,181,24,7,2,2,194,359,28,5,2,3,1,336,54,156,31,8,7,269,267,4,13,27,4,7,1,282,118,29,63,70,28,1,1,440,6,8,120,12,1,5,143,210,123,111,4,1,304,93,19,44,24,77,15,2,1,10,3,319,8,193,68,3,1,263,58,258,7,3,2,1,523,36,18,1,11,3,349,198,7,38,343,112,133,1,1,1,1,192,91,48,147,113,1,299,144,133,5,9,2,3,65,176,221,88,14,4,3,9,3,4,2,406,186,440,8,9,97,7,5,19,1,3,3,115,130,120,27,143,33,23,1,387,133,12,10,17,11,6,9,6,1,15,373,203,1,140,423,10,5,14,355,227,2,4,1,1,2,90,322,46,57,21,21,27,3,4,1,475,68,22,9,5,1,4,5,3,205,106,245,32,3,1,445,48,54,45,530,23,20,7,4,6,1,1,469,82,27,6,3,1,3,1,461,17,36,49,6,4,5,8,5,1,302,72,63,107,27,12,9,476,115,1,363,75,134,9,6,2,1,1,1,229,2,85}
	//
	//lengths := minSumOfLengths(a, 592)
	//fmt.Println(lengths)
	//fmt.Println(findLeastNumOfUniqueInts([]int{4,1,3,2,4},  2))
	//fmt.Println(minSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 1}, 3))
	//var num uint64 = 234234234
	//fmt.Println(strconv.FormatUint(num, 10))
	fmt.Printf("%v", (int64(math.Pow(2, 16))))
}
