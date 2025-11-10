package main

import "fmt"

func main() {
	var N, M int
	_, _ = fmt.Scan(&N, &M)

	slice1 := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&slice1[i])
	}

	slice2 := make([]int, M)
	for i := 0; i < M; i++ {
		_, _ = fmt.Scan(&slice2[i])
	}

	for _, value := range slice2 {
		result := leftBinarySearch(slice1, value)
		fmt.Println(result)
	}

}

func leftBinarySearch(slice1 []int, value int) int {
	left := -1
	right := len(slice1) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if slice1[mid] < value {
			left = mid //+ 1
		} else {
			right = mid // номер найденной позиции. Нумерация с 1го.
		}

	}
	if slice1[right] == value {
		return right + 1
	}
	return 0
}
