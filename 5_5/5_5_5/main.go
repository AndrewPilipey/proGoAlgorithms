package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	slice := make([]int, N)
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}

	sums := make([]int, N)
	for i, num := range slice {
		sums[i] = sumfDigits(num)
	}

	result := counterSort(sums)
	fmt.Println(result)
}

func sumfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func counterSort(sums []int) int {
	maxSum := 36 // это макс, ибо 9+9+9+9
	countsSum := make([]int, maxSum+1)
	for _, sum := range sums {
		countsSum[sum]++
	}

	maxCount := 0
	for _, count := range countsSum {
		if count > maxCount {
			maxCount = count
		}
	}

	for i := 0; i <= maxSum; i++ {
		if countsSum[i] == maxCount {
			return i
		}
	}

	return -1

}
