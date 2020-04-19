package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "mini-applications/quiz-app/problems.csv", "a csv file of format 'question, answer'")
	timer := flag.String("timeout", "2s", "time per question in seconds")
	flag.Parse()

	duration, err := time.ParseDuration(*timer)

	if err != nil {
		exit("Timeout value is not in right format, use values like 2s, 4m, 1h", 1)
	}

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
	correct := startTest(problems, duration)
	fmt.Println("You scored ", correct, "/", len(problems))
}

func startTest(p []problems, timer time.Duration) int {
	ch := make(chan string, 1)

	correct := 0
	for i, problem := range p {
		fmt.Printf("Problem %d: %s =\n", i+1, problem.q)
		go read(ch)

		select {
		case input := <-ch:
			{
				if input == problem.a {
					correct++
				}
			}
		case <-time.After(timer):
			{
				close(ch)
				fmt.Println("Too late. Exiting!")
				return correct
			}
		}
	}
	return correct
}

func read(c chan string) {
	var input string
	_, _ = fmt.Scanln(&input) //this will trim the spaces. No need to explicit trim
	c <- input
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
