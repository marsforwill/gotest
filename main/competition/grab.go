package main

import (
	"fmt"
	"sync"
)

func calculate(in ...[]int) []int {
	result := make([]int, len(in))
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i, s := range in {
		go func(index int, s []int) {
			sum := 0
			for _, value := range s {
				sum += value
			}
			result[index] = sum
			fmt.Println(index)
			//  result = append(result, sum)
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	return result
}

func main() {
	a := []int{1, 1, 1}
	b := []int{2, 2, 2}
	c := []int{3, 3, 3}
	fmt.Println(calculate(a, b, c))
}
