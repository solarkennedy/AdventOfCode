package main

import (
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func countAnswers(answers string) map[string]int {
	yesResponses := map[string]int{}
	for _, response := range strings.Split(answers, "\n") {
		for _, answeredYesTo := range strings.Split(response, "") {
			yesResponses[answeredYesTo] += 1
		}
	}
	return yesResponses
}

func partOne(input string) int {
	counter := 0
	for _, answers := range strings.Split(input, "\n\n") {
		counter += len(countAnswers(answers))
	}
	return counter
}

func countCommonYesAnswers(answers map[string]int, numberOfPeople int) int {
	counter := 0
	for _, numberAnsweredYesTo := range answers {
		if numberAnsweredYesTo == numberOfPeople {
			counter += 1
		}
	}
	return counter

}

func countPeople(answers string) int {
	return len(strings.Split(answers, "\n"))
}

func partTwo(input string) int {
	counter := 0
	for _, answers := range strings.Split(input, "\n\n") {
		yesAnswers := countAnswers(answers)
		numberOfPeople := countPeople(answers)
		counter += countCommonYesAnswers(yesAnswers, numberOfPeople)
	}
	return counter
}

func main() {
	input := utils.ReadInput()
	input = strings.TrimSpace(input)
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
