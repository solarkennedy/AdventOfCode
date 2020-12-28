package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type instruction struct {
	operation string
	argument  int
}

func parseInstruction(line string) instruction {
	parts := strings.Split(line, " ")
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return instruction{
		operation: parts[0],
		argument:  arg,
	}
}

func parseInstructions(input string) []instruction {
	input = strings.TrimSpace(input)
	instructions := []instruction{}
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

func partOne(instructions []instruction) int {
	instructionCounter := map[int]int{}
	acc := 0
	pc := 0
	for {
		if instructionCounter[pc] == 1 {
			fmt.Printf("We already ran line number %d. Breaking!\n", pc)
			break
		}
		instructionCounter[pc] += 1
		pc, acc = emulateInstruction(instructions[pc], acc, pc)
	}
	return acc
}

func emulateInstruction(in instruction, acc int, pc int) (int, int) {
	if in.operation == "acc" {
		acc += in.argument
		pc += 1
	} else if in.operation == "jmp" {
		pc += in.argument
	} else if in.operation == "nop" {
		pc += 1
	} else {
		panic(fmt.Errorf("Not programmed for '%s' operation on line %d", in.operation, pc))
	}
	return pc, acc
}

func main() {
	input := utils.ReadInput()
	instructions := parseInstructions(input)
	result := partOne(instructions)
	fmt.Printf("Answer to part one: %d\n", result)
}
