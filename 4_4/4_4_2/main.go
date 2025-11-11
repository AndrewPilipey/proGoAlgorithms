package main

import "fmt"

func main() {
	var N, M int
	_, _ = fmt.Scan(&N, &M)

	sliceN := make([]int, N)
	sliceM := make([]int, M)

	sliceN = (fillSlices(sliceN))
	sliceM = (fillSlices(sliceM))

	for _, value := range sliceM { //value - x in rightBinarySearch
		result := rightBinarySearch(sliceN, value)
		fmt.Println(result)
	}
}

func fillSlices(tempSlice []int) []int {
	for i := 0; i < len(tempSlice); i++ {
		_, _ = fmt.Scan(&tempSlice[i])
	}
	return tempSlice
}

func rightBinarySearch(sliceN []int, x int) int {
	left := 0
	right := len(sliceN)

	for left+1 < right {
		mid := left + (right-left)/2
		if sliceN[mid] <= x {
			left = mid
		} else { //если больше
			right = mid
		}
	}
	if sliceN[left] == x {
		return left + 1
	}
	return 0
}
