package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	list := make([]int, N)

	for i := range list {
		_, _ = fmt.Scan(&list[i])
	}

	insertionSort(list)

	for i := range list {
		fmt.Print(list[i], " ")
	}

}

func insertionSort(list []int) {
	var buffer int
	for i := 1; i < len(list); i++ {
		buffer = list[i]
		j := i
		for j > 0 && list[j-1] < buffer {
			list[j] = list[j-1]
			j--
		}
		list[j] = buffer
	}
}
