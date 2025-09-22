package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	question := int(math.Ceil(math.Log2(float64(N))))

	fmt.Println(question)

}
