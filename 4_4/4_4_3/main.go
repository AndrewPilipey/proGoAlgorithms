package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	sliceN := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&sliceN[i])
	}

	var x int
	_, _ = fmt.Scan(&x)

	fmt.Println(countOccurrences(sliceN, x))
}

func countOccurrences(sliceN []int, x int) int {
	leftBorder := leftBinarySearch(sliceN, x)
	rightBorder := rightBinarySearch(sliceN, x)

	if leftBorder == -1 || rightBorder == -1 {
		return 0
	}
	return rightBorder - leftBorder + 1
}

func leftBinarySearch(sliceN []int, x int) int {
	left, right := 0, len(sliceN)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if sliceN[mid] == x {
			result = mid
			right = mid - 1
		} else if sliceN[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func rightBinarySearch(sliceN []int, x int) int {
	left, right := 0, len(sliceN)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if sliceN[mid] == x {
			result = mid
			left = mid + 1
		} else if sliceN[mid] < x {
			left = mid + 1
		} else {

			right = mid - 1
		}
	}
	return result
}
