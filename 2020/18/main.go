package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func partOne(input string) int {
	counter := 0
	for _, line := range strings.Split(input, "\n") {
		tokens := tokenize(line)
		fmt.Printf("Evaluating %+v ...\n", tokens)
		answer := evaluate(tokens)
		fmt.Printf("Answer: %d!\n\n", answer)
		counter += answer
	}
	return counter
}

func isNumber(a byte) bool {
	r := regexp.MustCompile("[0-9]")
	return r.Match([]byte{a})
}

func nextRuneIsNumber(line string, i int) bool {
	if i+1 >= len(line) {
		return false
	} else {
		return isNumber(line[i+1])
	}
}

func tokenize(line string) []string {
	tokens := []string{}
	numberOfDigits := 0
	for i, potentialToken := range line {
		if potentialToken == '(' || potentialToken == ')' {
			tokens = append(tokens, string(potentialToken))
		} else if potentialToken == '*' || potentialToken == '+' {
			tokens = append(tokens, string(potentialToken))
		} else if potentialToken == ' ' {
			continue
		} else if nextRuneIsNumber(line, i) {
			numberOfDigits++
			continue
		} else if !nextRuneIsNumber(line, i) {
			tokens = append(tokens, line[i-numberOfDigits:i+1])
		}
	}
	return tokens
}

func findOpenParen(tokens []string) int {
	for i, token := range tokens {
		if token == "(" {
			return i
		}
	}
	return -1
}

func findCloseParen(tokens []string) int {
	depth := -1
	for i, token := range tokens {
		if token == "(" {
			depth++
		}
		if token == ")" {
			if depth == 0 {
				return i
			}
			depth--
		}
	}
	return -1
}

func removeElements(s []string, start int, end int) []string {
	return append(s[:start], s[end+1:]...)
}

func evaluate(tokens []string) int {
	fmt.Printf("  Evaluating '%+v'\n", tokens)
	if len(tokens) == 0 {
		return 0
	} else if len(tokens) == 1 {
		return utils.Atoi(tokens[0])
	} else if len(tokens) == 2 {
		panic(fmt.Errorf("  Somehow we got down to 2 tokens: %v", tokens))
	} else if len(tokens) == 3 {
		answer := performOperation(utils.Atoi(tokens[0]), tokens[1], utils.Atoi(tokens[2]))
		fmt.Printf("  Got %d!\n", answer)
		return answer
	}

	if tokens[0] == "(" {
		lhsEnd := findCloseParen(tokens)
		tokens[0] = strconv.Itoa(evaluate(tokens[1:lhsEnd]))
		tokens = removeElements(tokens, 1, lhsEnd)
	} else if tokens[2] == "(" {
		lhsEnd := findCloseParen(tokens)
		tokens[2] = strconv.Itoa(evaluate(tokens[3:lhsEnd]))
		tokens = removeElements(tokens, 3, lhsEnd)
	} else {
		tokens[0] = strconv.Itoa(evaluate(tokens[0:3]))
		tokens = removeElements(tokens, 1, 2)
	}

	return evaluate(tokens)
}

func performOperation(lhs int, op string, rhs int) int {
	if op == "*" {
		return lhs * rhs
	} else if op == "+" {
		return lhs + rhs
	} else {
		panic(fmt.Errorf("What operation is '%s'?", op))
	}
}

func findPlusLocation(tokens []string) int {
	for i, token := range tokens {
		if token == "+" {
			return i
		}
	}
	return -1
}

func evaluate2(tokens []string) int {
	fmt.Printf("  Evaluating '%+v'\n", tokens)
	if len(tokens) == 0 {
		return 0
	} else if len(tokens) == 1 {
		return utils.Atoi(tokens[0])
	} else if len(tokens) == 2 {
		panic(fmt.Errorf("  Somehow we got down to 2 tokens: %v", tokens))
	} else if len(tokens) == 3 {
		answer := performOperation(utils.Atoi(tokens[0]), tokens[1], utils.Atoi(tokens[2]))
		fmt.Printf("  Got %d!\n", answer)
		return answer
	}

	openParenLocation := findOpenParen(tokens)
	for openParenLocation != -1 {
		closeParenLocation := findCloseParen(tokens)
		tokens[openParenLocation] = strconv.Itoa(evaluate2(tokens[openParenLocation+1 : closeParenLocation]))
		tokens = removeElements(tokens, openParenLocation+1, closeParenLocation)
		fmt.Printf("  Tokens are now %+v\n", tokens)
		openParenLocation = findOpenParen(tokens)
	}

	plusLocation := findPlusLocation(tokens)
	for plusLocation != -1 {
		closePlusLocation := plusLocation + 2
		tokens[plusLocation-1] = strconv.Itoa(evaluate2(tokens[plusLocation-1 : closePlusLocation]))
		tokens = removeElements(tokens, plusLocation, closePlusLocation-1)
		fmt.Printf("  Tokens are now %+v\n", tokens)
		plusLocation = findPlusLocation(tokens)
	}

	if len(tokens) > 3 {
		tokens[0] = strconv.Itoa(evaluate2(tokens[0:3]))
		tokens = removeElements(tokens, 1, 2)
		fmt.Printf("  Tokens are now %+v\n", tokens)
	}

	return evaluate2(tokens)
}

func partTwo(input string) int {
	counter := 0
	for _, line := range strings.Split(input, "\n") {
		tokens := tokenize(line)
		fmt.Printf("Advanced Evaluating %+v ...\n", tokens)
		answer := evaluate2(tokens)
		fmt.Printf("Answer: %d!\n\n", answer)
		counter += answer
	}
	return counter
}

func main() {
	input := utils.ReadInput()
	input = strings.TrimSpace(input)
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
