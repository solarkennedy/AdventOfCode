package main

import (
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func csvToSlice(input string, n int) ([]int, int) {
	input = strings.TrimSpace(input)
	s := make([]int, n)
	starters := strings.Split(input, ",")
	for i, c := range starters {
		s[i] = utils.Atoi(c)
	}
	return s, len(starters)
}

func findLastTimeSpoken(thing int, history []int) int {
	for i := len(history) - 1; i >= 0; i-- {
		if history[i] == thing {
			return i
		}
	}
	return -1
}

func turn(history []int, turnNumber int) int {
	lastSpoken := history[turnNumber-2]
	lastTimeSpoken := findLastTimeSpoken(lastSpoken, history[0:turnNumber-2])
	if lastTimeSpoken == -1 {
		return 0
	} else {
		return (turnNumber - 2) - lastTimeSpoken
	}
}

func whatIsSaidOnTurn(n int, input string) int {
	history, turnNumber := csvToSlice(input, n)
	for turnNumber < n {
		history[turnNumber] = turn(history, turnNumber+1)
		turnNumber++
		if turnNumber%1000 == 0 {
			perc := float64(turnNumber*100) / float64(n)
			fmt.Printf("Completed: %.2f\n", perc)
		}
	}
	return history[n-1]
}

func partOne(input string) int {
	return whatIsSaidOnTurn(2020, input)
}

func partTwo(input string) int {
	return whatIsSaidOnTurn(30000000, input)
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
