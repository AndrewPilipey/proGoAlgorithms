package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)

	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	bubbleSort(slice)
	for i := range slice {
		fmt.Print(slice[i], " ")
	}
}

func bubbleSort(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if slice[j+1] > slice[j] {
				slice[j+1], slice[j] = slice[j], slice[j+1]
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}
