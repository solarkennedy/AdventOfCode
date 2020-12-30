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

func main() {
	input := utils.ReadInput()
	ints := utils.ConvertIntoInts(input)

	result := partOne(ints)
	fmt.Printf("Answer to part one: %d\n", result)
}
