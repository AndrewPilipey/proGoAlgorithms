package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&slice[i])
	}

	var x int
	_, _ = fmt.Scan(&x)

	foundX := findLastX(slice, x)

	if foundX != -1 {
		fmt.Println(foundX)
	}
}

func findLastX(slice []int, x int) int {
	lastX := -1
	for index, value := range slice {
		if value == x {
			lastX = index + 1
		}
	}
	return lastX
}
