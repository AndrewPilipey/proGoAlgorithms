package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	_, _ = fmt.Scanln(&N)

	slice := make([]string, N)

	scanner := bufio.NewScanner(os.Stdin)
	for i := range slice {
		_ = scanner.Scan()
		slice[i] = scanner.Text()
	}

	counterSort(slice, N)
}

func counterSort(slice []string, N int) {
	k := 26
	counts := make([]int, k)
	for i := 0; i < N; i++ {
		words := strings.Fields(slice[i])
		for _, word := range words {
			if len(word) > 0 {
				firstLetter := word[0]
				index := int(firstLetter - 'a')
				counts[index]++
			}
		}
	}

	for i := 0; i < 26; i++ {
		if counts[i] > 0 {
			letter := string('a' + i)
			fmt.Printf("%s %d\n", letter, counts[i])

		}
	}

}
