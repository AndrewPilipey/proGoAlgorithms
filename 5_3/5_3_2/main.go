package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	list := make([]int, N)
	for i := range list {
		_, _ = fmt.Scan(&list[i])
	}

	selectionSort(list)

	for j := range list {
		fmt.Print(list[j], " ")
	}

}

func selectionSort(list []int) {
	for i := len(list) - 1; i > 0; i-- {
		minIndex := i
		for j := 0; j < i; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}

}
