package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	grades := make([]int, N)

	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&grades[i])
	}

	minGrade := minGradeFinding(grades)
	maxGrade := maxGradeFinding(grades)

	replacedGrades := make([]int, len(grades))
	replacedGrades = replacingGrades(grades, minGrade, maxGrade)

	//и тут вывожу:
	for _, replacedGrades := range replacedGrades {
		fmt.Print(replacedGrades, " ")
	}
}

func minGradeFinding(grades []int) int {
	minGrade := 5

	for _, value := range grades {
		if value < minGrade {
			minGrade = value
		}
	}
	return minGrade
}

func maxGradeFinding(grades []int) int {
	maxGrade := 0

	for _, value := range grades {
		if value > maxGrade {
			maxGrade = value
		}
	}
	return maxGrade
}

func replacingGrades(grades []int, minGrade int, maxGrade int) []int {

	for i, _ := range grades {
		if grades[i] == minGrade {
			grades[i] = maxGrade
		}
	}
	return grades
}
