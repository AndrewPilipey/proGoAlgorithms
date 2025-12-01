package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	reader := bufio.NewReader(os.Stdin)

	// Пропускаем оставшийся символ новой строки после fmt.Scan
	_, _ = reader.ReadByte() // читаем и игнорируем '\n'

	sliceN := fillSlices(reader, N)
	sliceK := fillSlices(reader, K)

	results := make([]int, K)
	for i, value := range sliceK {
		results[i] = leftBinarySearch(sliceN, value)
	}
	for _, res := range results {
		fmt.Println(res)
	}
}

func fillSlices(reader *bufio.Reader, expectedCount int) []int {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Ошибка чтения строки: %v", err))
	}

	tokens := strings.Fields(line)

	if len(tokens) != expectedCount {
		panic(fmt.Sprintf(
			"Ожидалось %d чисел, но получено %d. Строка: '%s'",
			expectedCount, len(tokens), line,
		))
	}

	slice := make([]int, expectedCount)
	for i, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			panic(fmt.Sprintf("Не удалось преобразовать '%s' в число: %v", token, err))
		}
		slice[i] = num
	}
	return slice
}

func leftBinarySearch(sliceN []int, value int) int {
	left, right := 0, len(sliceN)-1

	for left <= right {
		mid := left + (right-left)/2
		if sliceN[mid] == value {
			return sliceN[mid]
		} else if sliceN[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == 0 {
		return sliceN[0]
	}
	if right == len(sliceN)-1 {
		return sliceN[len(sliceN)-1]
	}

	leftVal := sliceN[left-1]
	rightVal := sliceN[left]

	if abs(value-leftVal) <= abs(value-rightVal) {
		return leftVal
	}
	return rightVal
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
