// Напишите программу, которая сортирует массив по
// неубыванию с помощью сортировки вставками.

package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	insertionSort(slice)

	for i := range slice {
		fmt.Print(slice[i], " ")
	}

}

func insertionSort(slice []int) {
	var buffer int
	for i := 1; i < len(slice); i++ {
		buffer = slice[i]
		j := i
		for j > 0 && slice[j-1] > buffer {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = buffer
	}
}
