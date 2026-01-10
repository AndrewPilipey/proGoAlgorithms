package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	freq := make([]int, 9)

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		length := len(strconv.Itoa(num))
		freq[length]++
	}

	minFreq := N + 1
	for i := 1; i <= 8; i++ {
		if freq[i] > 0 && freq[i] < minFreq {
			minFreq = freq[i]
		}
	}

	for i := 1; i <= 8; i++ {
		if freq[i] == minFreq {
			fmt.Println(i, minFreq)
			return
		}
	}
}
