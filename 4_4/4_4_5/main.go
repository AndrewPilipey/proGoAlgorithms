package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	list := make([]int, 7)

	for i := 0; i < len(list); i++ {
		fmt.Scan(&list[i])
	}

	fmt.Println(binarySearch(list, x))
}

func binarySearch(list []int, x int) int {
	left := -1
	right := len(list) - 1
	for left+1 < right {
		m := left + (right-left)/2
		if list[m] < x {
			left = m
		} else {
			right = m
		}
	}
	if list[right] == x {
		return right
	}
	return -1
}
