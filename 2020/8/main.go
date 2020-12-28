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

func partOne(instructions []instruction) (int, error) {
	instructionCounter := map[int]int{}
	acc := 0
	pc := 0
	for pc != len(instructions) {
		if instructionCounter[pc] == 1 {
			return acc, fmt.Errorf("We already ran line number %d. Breaking!\n", pc)
		}
		instructionCounter[pc] += 1
		pc, acc = emulateInstruction(instructions[pc], acc, pc)
	}
	return acc, nil
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

func swapLine(lineNo int, instructions []instruction) ([]instruction, error) {
	newInstructions := make([]instruction, len(instructions))
	copy(newInstructions, instructions)
	if instructions[lineNo].operation == "jmp" {
		newInstructions[lineNo].operation = "nop"
	} else if instructions[lineNo].operation == "nop" {
		newInstructions[lineNo].operation = "jmp"
	} else {
		return nil, fmt.Errorf("Not swapping instruction %+v on line %d", instructions[lineNo], lineNo)
	}
	return newInstructions, nil
}

func partTwo(instructions []instruction) (int, error) {
	for lineNo := range instructions {
		newInstructions, err := swapLine(lineNo, instructions)
		if err != nil {
			continue
		}
		acc, err := partOne(newInstructions)
		if err == nil {
			return acc, nil
		}

	}
	return 0, fmt.Errorf("No amount of swapping resulted in a terminal program")
}

func main() {
	input := utils.ReadInput()
	instructions := parseInstructions(input)
	var result int
	var err error
	result, err = partOne(instructions)
	if err == nil {
		panic("Part one program terminated without loop? Unexpected")
	}
	fmt.Printf("Answer to part one: %d\n", result)
	result, err = partTwo(instructions)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer to part one: %d\n", result)
}
