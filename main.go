package main

import (
	"GoPhercises/quiz"
	"flag"
)

func main() {
	path := flag.String("csv", "", "path to csv file")

	flag.Parse()
	quizData := quiz.ReadQuestions(*path)

	// fmt.Println("here is the quizdata: ", quizData)

	quiz.Play(quizData)
}
