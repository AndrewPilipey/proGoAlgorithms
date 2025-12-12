package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	list := make([]int, N)
	for input := range list {
		_, _ = fmt.Scan(&list[input])
	}

	selectionSort(list)

	for output := range list {
		fmt.Print(list[output], " ")
	}

}

func selectionSort(list []int) {
	for i := len(list) - 1; i > 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			list[maxIndex], list[i] = list[i], list[maxIndex]
		}
	}

}
