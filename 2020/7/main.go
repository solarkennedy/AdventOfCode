package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

func parseContents(contents string) []string {
	colorsThatCanBeContained := []string{}
	r := regexp.MustCompile("(?P<ammount>[0-9]+) (?P<contents>.*) bag")
	for _, content := range strings.Split(contents, ",") {
		if content == "no other bags" {
			continue
		}
		match := r.FindStringSubmatch(content)
		if len(match) < 2 {
			panic(fmt.Errorf("Something went wrong matching this regex '%s' to the input '%s'\n", r.String(), content))
		}
		colorsThatCanBeContained = append(colorsThatCanBeContained, match[2])
	}
	return colorsThatCanBeContained
}

func parseRule(rule string) (string, []string) {
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
		containerColor, colorsThatCanBeContained := parseRule(rule)
		for _, colorThatCanBeContained := range colorsThatCanBeContained {
			rules[colorThatCanBeContained] = append(rules[colorThatCanBeContained], containerColor)
		}
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

func partTwo(input string) int {
	return 0
}

func main() {
	input := utils.ReadInput()
	input = strings.TrimSpace(input)
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
