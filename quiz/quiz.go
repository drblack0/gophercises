package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadQuestions() {
	file, err := os.Open("./quiz/problems.csv")

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

	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Println()
	}

}
