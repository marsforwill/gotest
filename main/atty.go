package main

import "fmt"
/*
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 1 1 1 1 2 2 2 2 2 3 3 3 3 3 4 4 4 4 4 5 5 5 5 5 6 6
1 1 1 1 1 2 2 2 2 2 4 4 4 4 4 6 6 6 6 6 9 9 9 9 9 12 12
1 1 1 1 1 2 2 2 2 2 4 4 4 4 4 6 6 6 6 6 9 9 9 9 9 13 13
ans:13
f(4,90)=f(3,90)+f(3,90-25)+f(3,90-2*25)+f(3,90-3*25)
 */
func waysToChange(n int) int {
	var ans = make([]int, n+5)
	coin := []int{1,5,10,25}
	ans[0] =1
	for i := 0; i < 4; i++ {
		for j:=0; j<=n; j++ {
			if j-coin[i] >= 0 {
				ans[j] += ans[j-coin[i]]
			}
		}
		for j:=0; j<=n; j++ {
			fmt.Printf("%v ",ans[j])
		}
		fmt.Println()
	}
	return ans[n]

}



func main() {
	ans := waysToChange(5)
	fmt.Printf("ans:%v",ans)
}
