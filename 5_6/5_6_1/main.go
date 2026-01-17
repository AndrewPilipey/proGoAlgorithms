package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	counterSort(slice, N)
}

func counterSort(slice []int, N int) {
	min, max := slice[0], slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	rangeSize := max - min + 1
	counts := make([]int, rangeSize)

	for _, x := range slice {
		counts[x-min]++
	}
	for i := 0; i < rangeSize; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Print(i+min, " ")
		}
	}
}
