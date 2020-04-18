package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "mini-applications/quiz-app/problems.csv", "a csv file of format 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit("File not found "+*csvFile, 1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Unable to parse the csv file", 1)
	}

	problems := parseCsv(lines)
	correct := startTest(problems)
	fmt.Println("You scored ", correct, "/", len(problems))
}

func startTest(p []problems) int {
	correct := 0
	for i, problem := range p {
		fmt.Printf("Problem %d: %s =\n", i+1, problem.q)
		var input string
		fmt.Scanln(&input) //this will trim the spaces. No need to explicit trim
		if input == problem.a {
			correct++
		}
	}
	return correct
}

func parseCsv(lines [][]string) []problems {
	prob := make([]problems, len(lines))
	for key, value := range lines {
		prob[key] = problems{
			q: value[0],
			a: strings.TrimSpace(value[1]),
		}
	}
	return prob
}

type problems struct {
	q string
	a string
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
