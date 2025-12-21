package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	result := counterSort(slice, N)
	fmt.Println(result)
}

func counterSort(slice []int, N int) int {
	const k = 101
	counts := make([]int, k)
	for i := 0; i < N; i++ {
		x := slice[i]
		counts[x] += 1
	}

	minCount := N + 1
	var result int

	for i := 0; i < k; i++ {
		if counts[i] > 0 && counts[i] < minCount {
			minCount = counts[i]
			result = i
		}
	}
	return result
}
