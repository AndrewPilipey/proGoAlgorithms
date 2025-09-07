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

	index := findingX(slice, x)

	if index != -1 {
		fmt.Println(index + 1)
	}

}

func findingX(slice []int, x int) int {
	for index, value := range slice {
		if value == x {
			return index
		}
	}
	return -1
}
