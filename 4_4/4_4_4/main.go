// Даны два массива.
// Для каждого элемента второго массива определите,
// сколько раз он встречается в первом массиве.
package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	sliceN := make([]int, N)
	sliceN = fillSlices(sliceN)

	var M int
	_, _ = fmt.Scan(&M)

	sliceM := make([]int, M)
	sliceM = fillSlices(sliceM)

	for _, value := range sliceM {
		fmt.Print(countOccurrences(sliceN, value), " ")
	}
}

func fillSlices(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	return slice
}

func countOccurrences(sliceN []int, x int) int {
	rightBorder := rightBinarySearch(sliceN, x)
	leftBorder := leftBinarySearch(sliceN, x)

	if rightBorder == -1 || leftBorder == -1 {
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
