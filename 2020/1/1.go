package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/solarkennedy/AdventOfCode/utils"
)

const (
	debug = false
)

func inputToIntSlice(input string) []int {
	output := []int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		output = append(output, i)

	}
	return output
}

func one(inputSlice []int) int {
	for i, n := range inputSlice {
		// Addition is commutative, so we don't need to accidentally-quadratic this thing
		for j, m := range inputSlice[i:] {
			if n+m == 2020 {
				if debug {
					fmt.Println("Found two numbers that sum to 2020:")
					fmt.Printf("%d (line %d)\n", n, i+1)
					fmt.Printf("%d (line %d)\n", m, j+1)
				}
				return m * n
			}
		}
	}
	panic("Couldn't find the answer")
}

func onePartTwo(inputSlice []int) int {
	for i, n := range inputSlice {
		for j, m := range inputSlice[i:] {
			if n+m >= 2020 {
				continue
			}
			for k, o := range inputSlice[j:] {
				if n+m+o == 2020 {
					if debug {
						fmt.Println("Found three numbers that sum to 2020:")
						fmt.Printf("%d (line %d)\n", n, i+1)
						fmt.Printf("%d (line %d)\n", m, j+1)
						fmt.Printf("%d (line %d)\n", o, k+1)
					}
					return m * n * o
				}
			}
		}
	}
	panic("Couldn't find the answer")
}

func main() {
	input := utils.ReadInput()
	inputSlice := inputToIntSlice(input)
	result := one(inputSlice)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := onePartTwo(inputSlice)
	fmt.Printf("Answer to part two: %d\n", result2)
}
