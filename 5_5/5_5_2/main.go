package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	result := counterSort(slice)
	fmt.Println(result)

}

func counterSort(slice []int) int {
	counter := 0
	k := 1000
	counts := make([]int, k+1)
	var n int
	n = len(slice)
	for i := 0; i < n; i++ {
		x := slice[i]
		counts[x] += 1
	}

	for i := range counts {
		if counts[i] == 2 {
			counter++
		}
	}
	return counter
}
