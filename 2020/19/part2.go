package main

import (
	"fmt"
	"regexp"
	"strings"
)

func modifyRulesForPart2(rules []rule) {
	rules[8] = rule{
		number:   8,
		ruleSet1: []int{42},
		ruleSet2: []int{42, 8},
		rawRule:  "8: 42 | 42 8",
	}
	rules[11] = rule{
		number:   8,
		ruleSet1: []int{42, 31},
		ruleSet2: []int{42, 11, 43},
		rawRule:  "11: 42 31 | 42 11 31",
	}
}

func specialRule8(rules []rule) string {
	return ruleToRegex(rules[42], rules) + `+`
}

func specialRule11(rules []rule) string {
	reg := ``
	r := "42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31"
	r = `42 31 |  42  42  31  31  |  42  42  42  31  31  31  |  42  42  42  42  31  31  31  31  |  42  42  42  42  42  31  31  31  31  31`
	for _, s := range strings.Split(r, " ") {
		if s == "42" {
			reg += ruleToRegex(rules[42], rules)
		} else if s == "31" {
			reg += ruleToRegex(rules[31], rules)
		} else {
			reg += s
		}
	}
	return reg
}

func ruleToRegexMulti(ruleSet []int, rules []rule) string {
	regex := ``
	for _, r := range ruleSet {
		regex += ruleToRegex(rules[r], rules)
	}
	return regex
}

func ruleToRegex(r rule, rules []rule) string {
	if r.mainRule != "" {
		return r.mainRule
	}
	if r.number == 8 {
		return `(?:` + specialRule8(rules) + `)`
	}
	if r.number == 11 {
		return `(?:` + specialRule11(rules) + `)`
	}
	if r.number == 0 {
		fmt.Printf("Computing mega main rule for 0\n")
	}
	reg := `(?:`
	reg += ruleToRegexMulti(r.ruleSet1, rules)
	if len(r.ruleSet2) > 0 {
		reg += `|` + ruleToRegexMulti(r.ruleSet2, rules)
	}
	reg += ")"
	fmt.Printf("Final regex for rule %d is `%s`\n", r.number, reg)
	return reg
}

func rulesToRegexp(rules []rule) *regexp.Regexp {
	totalRegex := `^`
	totalRegex += ruleToRegex(rules[0], rules)
	totalRegex += `$`
	fmt.Printf("Final regex for rule 0 is `%s`\n", totalRegex)
	return regexp.MustCompile(totalRegex)
}

func messageMatchesRegexRule0(message string, regex *regexp.Regexp) bool {
	return regex.MatchString(message)
}

func howManyMessagesMatchRegexRule0(regex *regexp.Regexp, messages []string) int {
	counter := 0
	for _, m := range messages {
		if messageMatchesRegexRule0(m, regex) {
			counter++
		}
	}
	return counter
}

func partTwo(input string) int {

	rules, messages := parseMessagesAndRules(input)

	s := strings.Split(input, "\n\n")
	out := parseVendor(s[0])
	fmt.Printf("vendor out: `%s`\n", out)
	unrolled := unloopWithDepth(s[0], 6)
	fmt.Printf("Unrolled : `%s`\n", unrolled)
	out2 := parseVendor(unrolled)
	fmt.Printf("vendor out2: `%s`\n", out2)
	regex2 := regexp.MustCompile(out2)
	v := howManyMessagesMatchRegexRule0(regex2, messages)
	fmt.Printf("Vendor answer is %d\n", v)

	modifyRulesForPart2(rules)
	regex := rulesToRegexp(rules)
	return howManyMessagesMatchRegexRule0(regex, messages)
}
