package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type location struct {
	x int
	y int
}

func inputToMap(input string) [][]bool {
	output := [][]bool{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		row := []bool{}
		for _, i := range line {
			row = append(row, i == '#')
		}
		output = append(output, row)

	}
	return output
}

func runSlope(snowMap [][]bool, right int, down int) int {
	counter := 0
	location := location{0, 0}
	width := len(snowMap[0])
	length := len(snowMap)
	for location.y < length {
		if snowMap[location.y][location.x] {
			counter = counter + 1
		}
		location.x = (location.x + right) % width
		location.y = location.y + down
	}
	return counter
}

func partOne(snowMap [][]bool) int {
	return runSlope(snowMap, 3, 1)
}

func partTwo(snowMap [][]bool) int {
	product := 1
	product *= runSlope(snowMap, 1, 1)
	product *= runSlope(snowMap, 3, 1)
	product *= runSlope(snowMap, 5, 1)
	product *= runSlope(snowMap, 7, 1)
	product *= runSlope(snowMap, 1, 2)
	return product
}

func main() {
	input := utils.ReadInput()
	snowMap := inputToMap(input)
	result := partOne(snowMap)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(snowMap)
	fmt.Printf("Answer to part two: %d\n", result2)
}
