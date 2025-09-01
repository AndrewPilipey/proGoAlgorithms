package main

import (
	"fmt"
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

	indexes := findIndex(slice, x)
	for _, index := range indexes {
		fmt.Print(index, " ")
	}

}

func findIndex(slice []int, x int) []int {
	var indexes []int
	for i, value := range slice {
		if x == value {
			indexes = append(indexes, i+1)
		}
	}
	return indexes
}
