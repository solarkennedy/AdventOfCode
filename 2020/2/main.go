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

type passwordPolicy struct {
	first  int
	second int
	letter rune
}

type entry struct {
	policy   passwordPolicy
	password string
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func parseEntryLine(line string) entry {
	policyPart, password := splitEntryLine(line)
	p := parsePolicy(policyPart)
	e := entry{
		policy:   p,
		password: password,
	}
	return e
}

func splitEntryLine(line string) (string, string) {
	v := strings.Split(line, ": ")
	return v[0], v[1]
}

func parsePolicy(p string) passwordPolicy {
	v := strings.Split(p, " ")
	letter := rune(v[1][0])
	vLen := strings.Split(v[0], "-")
	first := parseInt(vLen[0])
	second := parseInt(vLen[1])
	return passwordPolicy{
		first:  first,
		second: second,
		letter: letter,
	}
}

func inputToEntries(input string) []entry {
	output := []entry{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, parseEntryLine(line))
	}
	return output
}

func isValidPasswordPart1(policy passwordPolicy, password string) bool {
	count := strings.Count(password, string(policy.letter))
	return count >= policy.first && count <= policy.second
}

func isValidPasswordPart2(policy passwordPolicy, password string) bool {
	firstInPlace := password[policy.first-1] == byte(policy.letter)
	secondInPlace := password[policy.second-1] == byte(policy.letter)
	// Equivilant to xor in the boolean case
	return firstInPlace != secondInPlace
}

func partOne(entries []entry) int {
	counter := 0
	for _, ent := range entries {
		if isValidPasswordPart1(ent.policy, ent.password) {
			counter += 1
		}
	}
	return counter
}

func partTwo(entries []entry) int {
	counter := 0
	for _, ent := range entries {
		if isValidPasswordPart2(ent.policy, ent.password) {
			counter += 1
		}
	}
	return counter
}

func main() {
	input := utils.ReadInput()
	entries := inputToEntries(input)
	result := partOne(entries)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(entries)
	fmt.Printf("Answer to part two: %d\n", result2)
}
