package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		number := i
		if number%3 == 0 {
			sum = sum + number
		}
	}
	fmt.Println(sum)
	fmt.Printf("Программа отработала за %d наносекунд", time.Since(start).Nanoseconds())
}
