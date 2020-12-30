package main

import (
	"fmt"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func isValid(theInt int, lookBackInts []int) bool {
	for i := range lookBackInts {
		for j := range lookBackInts[:i] {
			if lookBackInts[i]+lookBackInts[j] == theInt {
				return true
			}
		}
	}
	return false
}

func findFirstInvalidNumber(ints []int, preamble int) int {
	lookBack := preamble
	for i, theInt := range ints {
		if i <= preamble {
			continue
		}
		if !isValid(theInt, ints[i-lookBack:i]) {
			return theInt
		}
	}
	return -1
}

func partOne(ints []int) int {
	return findFirstInvalidNumber(ints, 25)
}

func rangeSumsTo(n int, r []int) bool {
	sum := 0
	for i := range r {
		sum += r[i]
	}
	return sum == n
}

func findContinguousBlockThatAddsTo(n int, ints []int) []int {
	for start := range ints {
		for end := start; end <= len(ints); end++ {
			if rangeSumsTo(n, ints[start:end]) {
				return ints[start:end]
			}
		}
	}
	return []int{}
}

func partTwo(ints []int, firstInvalidNumber int) int {
	block := findContinguousBlockThatAddsTo(firstInvalidNumber, ints)
	min := utils.MinInt(block...)
	max := utils.MaxInt(block...)
	return min + max
}

func main() {
	input := utils.ReadInput()
	ints := utils.ConvertIntoInts(input)

	result := partOne(ints)
	fmt.Printf("Answer to part one: %d\n", result)

	result = partTwo(ints, result)
	fmt.Printf("Answer to part two: %d\n", result)
}
