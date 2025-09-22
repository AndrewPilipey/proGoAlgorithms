package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) bool {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sliceN := make([]int, n)
	sliceK := make([]int, k)

	for i := 0; i < n; i++ {
		fmt.Scan(&sliceN[i])
	}

	for i := 0; i < k; i++ {
		fmt.Scan(&sliceK[i])
	}

	for _, num := range sliceK {
		if binarySearch(sliceN, num) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
