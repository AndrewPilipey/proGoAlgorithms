package main

import "fmt"

func main() {
	var N, x int

	_, _ = fmt.Scan(&N)
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&slice[i])
	}

	_, _ = fmt.Scan(&x)

	if linearSearch(slice, x) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func linearSearch(slice []int, x int) bool {
	for _, value := range slice {
		if value == x {
			return true
		}
	}
	return false
}
