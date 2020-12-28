package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type rule struct {
	color  string
	amount int
}

func parseContents(contents string) []rule {
	colorsThatCanBeContained := []rule{}
	r := regexp.MustCompile("(?P<amount>[0-9]+) (?P<contents>.*) bag")
	for _, content := range strings.Split(contents, ",") {
		if content == "no other bags" {
			continue
		}
		match := r.FindStringSubmatch(content)
		if len(match) < 2 {
			panic(fmt.Errorf("Something went wrong matching this regex '%s' to the input '%s'\n", r.String(), content))
		}
		amount, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		r := rule{
			amount: amount,
			color:  match[2],
		}
		colorsThatCanBeContained = append(colorsThatCanBeContained, r)
	}
	return colorsThatCanBeContained
}

func parseRule(rule string) (string, []rule) {
	r := regexp.MustCompile("(?P<containingColor>[a-z ]*) bags contain (?P<contants>.*).")
	match := r.FindStringSubmatch(rule)
	if len(match) < 2 {
		panic(fmt.Errorf("Something went wrong matching this regex '%s' to the input '%s'\n", r.String(), rule))
	}
	containerColor := match[1]
	colorsThatCanBeContained := parseContents(match[2])
	return containerColor, colorsThatCanBeContained
}

func parseRules(input string) map[string][]string {
	rules := map[string][]string{}
	for _, rule := range strings.Split(input, "\n") {
		containerColor, colorRules := parseRule(rule)
		for _, colorThatCanBeContained := range colorRules {
			rules[colorThatCanBeContained.color] = append(rules[colorThatCanBeContained.color], containerColor)
		}
	}
	return rules
}

func parseRulesWithAmount(input string) map[string][]rule {
	rules := map[string][]rule{}
	for _, rule := range strings.Split(input, "\n") {
		containerColor, colorRules := parseRule(rule)
		rules[containerColor] = colorRules
	}
	return rules
}

func ColorsThatCanContainA(color string, rules map[string][]string) []string {
	colorsThatCanContain, ok := rules[color]
	if !ok {
		return []string{}
	}
	for _, c := range colorsThatCanContain {
		colorsThatCanContain = append(colorsThatCanContain, ColorsThatCanContainA(c, rules)...)
	}
	return utils.RemoveDuplicatesStrings(colorsThatCanContain)
}

func partOne(input string) int {
	rules := parseRules(input)
	c := ColorsThatCanContainA("shiny gold", rules)
	fmt.Println(c)
	return len(c)
}

func getAmountNeededToContain(color string, allRules map[string][]rule) int {
	amount := 0
	for _, rules := range allRules[color] {
		amount += rules.amount
		amount += rules.amount * getAmountNeededToContain(rules.color, allRules)
	}
	return amount
}

func partTwo(input string) int {
	rules := parseRulesWithAmount(input)
	return getAmountNeededToContain("shiny gold", rules)

}

func main() {
	input := utils.ReadInput()
	input = strings.TrimSpace(input)
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
