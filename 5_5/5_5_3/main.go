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
	k := 10_001
	counts := make([]int, k)
	for i := 0; i < N; i++ {
		x := slice[i]
		counts[x] += 1
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	for i := k - 1; i >= 0; i-- {
		if counts[i] == maxCount {
			return i
		}
	}
	return -1
}
