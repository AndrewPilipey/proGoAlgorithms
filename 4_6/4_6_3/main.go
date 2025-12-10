package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	sliceN := make([]int, N)
	for i := range sliceN {
		_, _ = fmt.Scan(&sliceN[i])
	}

	var x int
	_, _ = fmt.Scan(&x)

	fmt.Println(JumpSearch(sliceN, x))
}

func JumpSearch(sliceN []int, x int) int {
	counter := 0
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
		return 0 //break?
	}

	for i := start; i <= len(sliceN)-1; i++ {
		if sliceN[i] == x {
			counter++
			continue
		}
	}
	return counter
}
