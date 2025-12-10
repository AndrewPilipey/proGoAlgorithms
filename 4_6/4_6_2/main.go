package main

import (
	"fmt"
	"math"
)

func fillSlices(slice []int) []int {
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}
	return slice
}

func main() {
	var N, M int
	_, _ = fmt.Scan(&N, &M)

	sliceN := make([]int, N)
	sliceN = fillSlices(sliceN)

	sliceM := make([]int, M)
	sliceM = fillSlices(sliceM)

	for _, value := range sliceM {
		fmt.Println(rightJumpSearch(sliceN, value))
	}
}

func rightJumpSearch(sliceN []int, x int) int {
	b := int(math.Sqrt(float64(len(sliceN))))
	start := 0
	end := b - 1

	for sliceN[end] <= x && end < len(sliceN)-1 && sliceN[end+1] <= x {
		start = int(math.Min(float64(len(sliceN)-1), float64(start+b)))
		end = int(math.Min(float64(len(sliceN)-1), float64(end+b)))
	}

	if x > sliceN[end] {
		return 0
	}

	for i := end; i >= start; i-- {
		if sliceN[i] == x {
			return i + 1
		}
	}
	return 0
}
