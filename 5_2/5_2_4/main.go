// В 39 школе в 7Б классе учатся N учеников.
// Недавно они писали полугодовые работы по информатике,
// математике и физике. Вам известны оценки каждого
// ученика по каждому из предметов.
// Выведите фамилии и имена учащихся в порядке
// невозрастания их среднего балла.

package main

import "fmt"

type StudentsAndScores struct {
	LastName  string
	FirstName string
	Score1    int
	Score2    int
	Score3    int
	AvgScore  int
}

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	var students []StudentsAndScores

	for i := 0; i < N; i++ {
		var lastName, firstName string
		var score1, score2, score3 int

		fmt.Scan(&lastName, &firstName, &score1, &score2, &score3)

		avgScore := (score1 + score2 + score3) / 2
		students = append(students, StudentsAndScores{
			LastName:  lastName,
			FirstName: firstName,
			Score1:    score1,
			Score2:    score2,
			Score3:    score3,
			AvgScore:  avgScore,
		})

	}

	n := len(students)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if students[j].AvgScore < students[j+1].AvgScore {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	for _, value := range students {
		fmt.Printf("%s %s\n", value.LastName, value.FirstName)
	}

}
