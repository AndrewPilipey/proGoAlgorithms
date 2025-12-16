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

	isertionSort(slice)

}

func isertionSort(slice []int) {
	var buffer int
	for i := 1; i < len(slice); i++ {
		flag := false
		buffer = slice[i]
		j := i
		for j > 0 && slice[j-1] > buffer {
			slice[j] = slice[j-1]
			j--
			flag = true

		}
		slice[j] = buffer
		if flag {
			for k := range slice {
				fmt.Print(slice[k], " ")
			}
			fmt.Println()
		}
	}
}
