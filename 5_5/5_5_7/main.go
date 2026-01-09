package main

import (
	"fmt"
	"sort"
)

type studentsAndScore struct {
	lastName  string
	firstName string
	score     int
}

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	students := make([]studentsAndScore, N)

	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&students[i].lastName, &students[i].firstName,
			&students[i].score)
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].score > students[j].score
	})

	for i := 0; i < 3 && i < len(students); i++ {
		fmt.Println(students[i].lastName, students[i].firstName)
	}

}
