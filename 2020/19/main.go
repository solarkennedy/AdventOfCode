package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type rule struct {
	number   int
	mainRule string
	ruleSet1 []int
	ruleSet2 []int
	rawRule  string
}

func getMainRule(input string) string {
	if strings.Contains(input, `"`) {
		return strings.Trim(input, `"`)
	}
	return ""
}

func getRuleSet1(input string) []int {
	r1 := []int{}
	if !strings.Contains(input, `"`) {
		lhs := strings.Split(input, `|`)
		els := strings.Split(strings.TrimSpace(lhs[0]), ` `)
		return spaceSepratedStringsToArray(els)
	}
	return r1
}

func getRuleSet2(input string) []int {
	r2 := []int{}
	if strings.Contains(input, `|`) {
		lhs := strings.Split(input, `|`)
		els := strings.Split(strings.TrimSpace(lhs[1]), ` `)
		r2 = spaceSepratedStringsToArray(els)
	}
	return r2
}

func spaceSepratedStringsToArray(els []string) []int {
	r := []int{}
	for _, e := range els {
		a, _ := strconv.Atoi(e)
		r = append(r, a)
	}
	return r
}

func parseRule(input string) rule {
	s := strings.Split(input, `: `)
	ruleNumber, err := strconv.Atoi(strings.TrimSpace(s[0]))
	if err != nil {
		panic(fmt.Sprintf("Error parsing rule '%s'", input))
	}
	r := s[1]
	return rule{
		number:   ruleNumber,
		mainRule: getMainRule(r),
		ruleSet1: getRuleSet1(r),
		ruleSet2: getRuleSet2(r),
		rawRule:  strings.TrimSpace(input),
	}
}

func parseRules(input string) []rule {
	splitStrings := strings.Split(strings.TrimSpace(input), "\n")
	rules := make([]rule, 1000)
	for _, r := range splitStrings {
		parsedRule := parseRule(r)
		rules[parsedRule.number] = parsedRule
	}
	return rules
}

func parseMessages(input string) []string {
	messages := []string{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		messages = append(messages, strings.TrimSpace(s))
	}
	return messages
}

func parseMessagesAndRules(input string) ([]rule, []string) {
	s := strings.Split(input, "\n\n")
	rules := parseRules(s[0])
	messages := parseMessages(s[1])
	return rules, messages
}

func messageMatchesMainRule(m string, mainRule string) int {
	if m[0] == byte(mainRule[0]) {
		return 1
	}
	return 0
}

func messageMatchesMultiRule(message string, rules []rule, ruleNumbers []int) int {
	counter := 0
	for _, ri := range ruleNumbers {
		remaining := message[counter:]
		if remaining == "" {
			return 0
		}
		matched := messageMatchesRule(remaining, rules, ri)
		if matched == 0 {
			return 0
		}
		counter += matched
	}
	return counter
}

func messageMatchesRule(message string, rules []rule, ruleNumber int) int {
	r := rules[ruleNumber]
	if r.mainRule != "" {
		return messageMatchesMainRule(message, r.mainRule)
	}
	if len(r.ruleSet1) != 0 {
		matched := messageMatchesMultiRule(message, rules, r.ruleSet1)
		if matched != 0 {
			return matched
		}
	}
	if len(r.ruleSet2) != 0 {
		matched := messageMatchesMultiRule(message, rules, r.ruleSet2)
		if matched != 0 {
			return matched
		}
	}
	return 0
}

func messageMatchesRule0(message string, rules []rule) bool {
	matched := messageMatchesRule(message, rules, 0)
	return matched == len(message)
}

func howManyMessagesMatchRule0(rules []rule, messages []string) int {
	count := 0
	for _, message := range messages {
		passed := ""
		if messageMatchesRule0(message, rules) {
			count++
			passed = "passed"
		}
		fmt.Printf("message %d: %s %s\n", count, message, passed)
	}
	return count
}

func partOne(input string) int {
	rules, messages := parseMessagesAndRules(input)
	return howManyMessagesMatchRule0(rules, messages)
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
