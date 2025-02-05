package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadQuestions(path string) [][]string {
	if path == "" {
		path = "./quiz/problems.csv"
	}
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}
	return data
}

func Play(data [][]string) {
	var finalScore int
	for _, question := range data {
		var answer string
		fmt.Print("Question: ", question[0]+" = ")
		fmt.Scan(&answer)
		if answer == question[1] {
			finalScore = handleScoring("Correct", finalScore)
			fmt.Println("Correct. Current Score: ", finalScore)
		} else {
			finalScore = handleScoring("Incorrect", finalScore)
			fmt.Println("Incorrect. Current Score: ", finalScore)
		}
	}
}

func handleScoring(answerStatus string, currentScore int) (finalScore int) {
	if answerStatus == "Correct" {
		return currentScore + 1
	} else if answerStatus == "Incorrect" {
		return currentScore
	}
	return -1
}
