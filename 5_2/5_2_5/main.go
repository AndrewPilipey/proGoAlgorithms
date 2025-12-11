package main

import (
	"fmt"
	"sort"
)

type studentAndScore struct {
	LastName  string
	FirstName string
	Score     int
}

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	var students []studentAndScore

	for i := 0; i < N; i++ {
		var lastName, firstName string
		var score int

		_, _ = fmt.Scan(&lastName, &firstName, &score)

		students = append(students, studentAndScore{
			LastName:  lastName,
			FirstName: firstName,
			Score:     score,
		})

	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Score > students[j].Score
	})

	for i := 0; i < 3 && i < len(students); i++ {
		fmt.Printf("%s %s\n", students[i].LastName, students[i].FirstName)
	}

}
