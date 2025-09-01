package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&slice[i])
	}

	var x int
	_, _ = fmt.Scan(&x)

	fmt.Println(findNearest(slice, x))
}

func findNearest(slice []int, x int) int {
	diffs := make([]int, len(slice))

	for i, value := range slice {
		diffs[i] = int(math.Abs(float64(value - x)))
	}

	minIndex := 0
	for i := 1; i < len(diffs); i++ {
		if diffs[i] < diffs[minIndex] {
			minIndex = i
		}
		if diffs[i] == diffs[minIndex] && slice[i] > slice[minIndex] {
			minIndex = i
		}
	}

	return slice[minIndex]
}
