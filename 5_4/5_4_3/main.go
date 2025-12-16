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

	var insertedNum int
	_, _ = fmt.Scan(&insertedNum)

	var position int
	_, _ = fmt.Scan(&position)

	result := insertionNum(slice, insertedNum, position)

	for i := range result {
		fmt.Print(result[i], " ")
	}
}

func insertionNum(slice []int, insertedNum int, position int) []int {
	newSlice := make([]int, len(slice)+1)

	copy(newSlice[:position-1], slice[:position-1])

	newSlice[position-1] = insertedNum

	copy(newSlice[position:], slice[position-1:])

	return newSlice
}
