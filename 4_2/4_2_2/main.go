package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&slice[i])
	}

	var key int
	_, _ = fmt.Scan(&key)
	fmt.Println(linearSearch(slice, key))
}

func linearSearch(slice []int, key int) int {
	counter := 0
	for _, value := range slice {
		if value == key {
			counter++
		}
	}
	return counter
}
