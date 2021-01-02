package main

import (
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func csvToSlice(input string) []int {
	input = strings.TrimSpace(input)
	s := []int{}
	for _, c := range strings.Split(input, ",") {
		s = append(s, utils.Atoi(c))
	}
	return s
}

func findLastTimeSpoken(thing int, history []int) int {
	for i := len(history) - 1; i >= 0; i-- {
		if history[i] == thing {
			return i
		}
	}
	return -1
}

func turn(history []int) int {
	lastSpoken := history[len(history)-1]
	lastTimeSpoken := findLastTimeSpoken(lastSpoken, history[0:len(history)-1])
	if lastTimeSpoken == -1 {
		return 0
	} else {
		return len(history) - lastTimeSpoken - 1
	}
}

func partOne(input string) int {
	history := csvToSlice(input)
	for len(history) != 2020 {
		history = append(history, turn(history))
	}
	return history[len(history)-1]
}

func partTwo(input string) int {
	return 0
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
