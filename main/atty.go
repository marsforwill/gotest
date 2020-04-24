package main

import "fmt"

func waysToChange(n int) int {
	if n < 0 {
		return 0
	}
	if n== 0 {
		return 1
	}
	if n > 25 {
		return waysToChange(n-25)
	}
	if n > 10 {
		return waysToChange(n-10)
	}
	if n > 5 {
		return waysToChange(n-5)
	}
	if n < 5 {
		return 1
	}
	return 1
}
func main() {
	ans := waysToChange(6)
	fmt.Printf("ans:%v",ans)
}
