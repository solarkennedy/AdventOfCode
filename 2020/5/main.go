package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func getMax(ids []int) int {
	max := ids[0]
	for _, id := range ids {
		if id > max {
			max = id
		}
	}
	return max
}

// seatToID takes a boarding pass string, like BBFFBBFRLL
// and turns it into an ID, per the rule that the ID
// is 'multiply the row by 8, then add the column'
// But it is way simpler than that, because the string
// is simply an integer in binary form, just in a funny
// representation.
func seatToID(seat string) int {
	r := strings.NewReplacer(
		"B", "1",
		"F", "0",
		"L", "0",
		"R", "1")
	seatInBinary := "0b" + r.Replace(seat)
	d, err := strconv.ParseInt(seatInBinary, 0, 64)
	if err != nil {
		panic(err)
	}
	return int(d)
}

func seatsToIDs(input string) []int {
	ids := []int{}
	for _, seat := range strings.Split(input, "\n") {
		if seat == "" {
			continue
		}
		ids = append(ids, seatToID(seat))
	}
	return ids
}

func partOne(input string) int {
	ids := seatsToIDs(input)
	return getMax(ids)
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
