// Дан массив целых чисел. Определите, сколько
// обменов сделает алгоритм пузырьковой
// сортировки по возрастанию для данного массива.
package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	list := make([]int, N)
	for i := range list {
		_, _ = fmt.Scan(&list[i])
	}

	counter := 0
	counter = bubbleSort(list, counter)
	fmt.Println(counter)
}

func bubbleSort(list []int, counter int) int {

	for i := len(list) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
				counter++
			}
		}
		if flag == false {
			return counter
		}
	}
	return counter
}
