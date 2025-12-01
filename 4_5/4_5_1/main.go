package main

import (
	"fmt"
	"math"
)

func main() {
	var N, K int
	_, _ = fmt.Scan(&N, &K)

	sliceN := make([]int, N)
	for i := 0; i < len(sliceN); i++ {
		_, _ = fmt.Scan(&sliceN[i])
	}

	sliceK := make([]int, K)
	for i := 0; i < len(sliceK); i++ {
		_, _ = fmt.Scan(&sliceK[i])
	}

	for _, value := range sliceK {
		fmt.Println(jumpSearch(sliceN, value))
	}
}

func jumpSearch(sliceN []int, x int) string {
	result := "NO"
	b := int(math.Sqrt(float64(len(sliceN))))
	start := 0
	end := b - 1
	for sliceN[end] < x {
		if end == len(sliceN)-1 {
			break
		}
		start = int(math.Min(float64(len(sliceN)-1), float64(start+b)))
		end = int(math.Min(float64(len(sliceN)-1), float64(end+b)))
	}
	if x > sliceN[end] {
		return result
	}
	for i := end; i >= start; i-- {
		if sliceN[i] == x {
			result = "YES"
			return result
		}
	}
	return result
}
