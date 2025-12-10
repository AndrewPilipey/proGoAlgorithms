package main

import (
	"fmt"
	"math"
)

func main() {
	x := 4
	slice := []int{3, 3, 4, 5, 7, 7, 8}

	fmt.Println(jumpSearch(slice, x))
}

func jumpSearch(list []int, x int) int {
	B := int(math.Sqrt(float64(len(list))))
	start := 0
	end := B - 1
	for list[end] < x {
		if end == len(list)-1 {
			break
		}
		start = int(math.Min(float64(len(list)-1), float64(start+B)))
		end = int(math.Min(float64(len(list)-1), float64(end+B)))
	}
	if x > list[end] {
		return -1
	}
	for i := start; i <= end; i++ {
		if list[i] == x {
			return i
		}
	}
	return -1
}
