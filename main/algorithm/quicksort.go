package main

import "fmt"

func quickSort(a *[]int, low int, high int) {
	if low >= high {
		return
	}
	keyPos := partitions(a, low, high)
	quickSort(a, low, keyPos-1)
	quickSort(a, keyPos+1, high)
}

func partitions(a *[]int, low int, high int) int {
	key := (*a)[low]
	for low < high {
		// 从后往前比key大需交换的数
		for high > low && (*a)[high] >= key {
			high--
		}
		//换到前面
		(*a)[low] = (*a)[high]
		// 从前往后比key小的数
		for low < high && (*a)[low] <= key {
			low++
		} // 换到后面
		(*a)[high] = (*a)[low]
	}
	(*a)[low] = key
	return low
}

func main() {
	a := []int{3, 5, 8, 6, 2}
	quickSort(&a, 0, 4)
	fmt.Println(a)
}
