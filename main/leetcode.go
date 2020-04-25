package main

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

	temp := ListNode{
		Val:  0,
		Next: nil,
	}
	head := &temp
	a := 0
	for i := 1; l1 != nil; i++ {
		a += l1.Val * i
		l1 = l1.Next
	}
	b := 0
	for i := 1; l2 != nil; i++ {
		b += l2.Val * i
		l2 = l2.Next
	}
	ans := a + b
	for i := 0; ans != 0; i++ {

		num := ans % 10

		temp2 := ListNode{
			Val:  num,
			Next: nil,
		}
		temp.Next = &temp2
		temp = *temp.Next

		ans = ans / 10
	}
	return head
}

// EncodeToString returns the base64 encoding of src.

func main() {

	a := ListNode{
		Val:  5,
		Next: nil,
	}
	b := ListNode{
		Val:  5,
		Next: nil,
	}
	addTwoNumbers(&a, &b)

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
