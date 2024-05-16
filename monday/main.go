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

	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("time", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *filename))
	}

	defer file.close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	if *shuffle {
		shuffleQuestions(lines)
	}

	correct := 0
	total := len(lines)

	fmt.Println("Press Enter to start the quiz...")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, line := range lines {
		question := line[0]
		answer := strings.TrimSpace(line[1])
		fmt.Printf("%s = ", question)
		answeCh := make(chan string)

		go func ()  {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			answeCh <- strings.TrimSpace(userAnswer)
		}()

		select {
		case <- timer.C:
			fmt.Println("\nTime's up!")
			printResult(correct, total)
			return
		case userAnswer := <- answer
		}
	}

}