package main

import (
	"flag"
	"fmt"
)

func main() {
	// quiz.ReadQuestions()

	boolPtr := flag.Bool("boolean", true, "a boolean")

	flag.Parse()

	fmt.Printf("here is the bool flag: %v", *boolPtr)
}
