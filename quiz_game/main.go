package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	value := make([]problem, len(lines))

	for i, line := range lines {
		value[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return value
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		showExit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}

	_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		showExit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You scored: %d out of %d\n", correct, len(problems))
}

func showExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
