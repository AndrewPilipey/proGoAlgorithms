// Напишите программу, которая сортирует массив по
// невозрастанию с помощью сортировки подсчетом.

package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := (make([]int, N))
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	counterSort(slice)

}

func counterSort(slice []int) {
	k := 100000
	counts := make([]int, k)
	var n int
	n = len(slice)
	for i := 0; i < n; i++ {
		var x int
		x = slice[i]
		counts[x] += 1
	}
	for i := k - 1; i >= 0; i-- {
		for j := 0; j < counts[i]; j++ {
			fmt.Print(i, " ")
		}
	}
}
