package main

import (
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type instruction struct {
	mask         string
	location     int
	preMaskValue int
}

func parseMemLocation(input string) int {
	first := strings.Split(input, "[")[1]
	second := strings.Split(first, "]")[0]
	return utils.Atoi(second)

}

func parseInstructions(input string) []instruction {
	ins := []instruction{}
	var mask string
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		tokens := strings.Split(s, " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
		} else {
			ins = append(ins, instruction{
				mask:         mask,
				location:     parseMemLocation(tokens[0]),
				preMaskValue: utils.Atoi(tokens[1]),
			})
		}
	}
	return ins
}

func bitSet(value int, place int) int {
	realMask := 1 << place
	return value | realMask
}

func bitClear(value int, place int) int {
	realMask := 1 << place
	return value &^ realMask
}

func computePostMaskValue(preValue int, mask string) int {
	value := preValue
	fmt.Printf("Value: %036b (%d)\n", value, value)
	fmt.Printf("Mask : %s\n", mask)
	for c := range mask {
		if mask[c] == 'X' {
			continue
		} else if mask[c] == '1' {
			value = bitSet(value, 35-c)
		} else if mask[c] == '0' {
			value = bitClear(value, 35-c)
		} else {
			panic(mask[c])
		}
	}
	fmt.Printf("Res  : %036b (%d)\n", value, value)
	return value
}

func computePostMaskValues(ins []instruction) int {
	m := map[int]int{}
	for i := range ins {
		fmt.Printf("Ins %d (%d):\n", i, ins[i].location)
		m[ins[i].location] = computePostMaskValue(ins[i].preMaskValue, ins[i].mask)
	}
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func partOne(input string) int {
	ins := parseInstructions(input)
	return computePostMaskValues(ins)
}

func partTwo(input string) int {
	return 42
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
