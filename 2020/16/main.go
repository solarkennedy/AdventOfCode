package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type rule struct {
	minA int
	maxA int
	minB int
	maxB int
}

type ticket struct {
	fields []int
}

func parseInput(input string) ([]rule, ticket, []ticket) {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	myTicket := parseMyTicket(parts[1])
	nearbyTickets := parseNearbyTickets(parts[2])
	return rules, myTicket, nearbyTickets
}

func parseRule(input string) rule {
	reg := regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)`)
	match := reg.FindStringSubmatch(input)
	minA, _ := strconv.Atoi(match[2])
	maxA, _ := strconv.Atoi(match[3])
	minB, _ := strconv.Atoi(match[4])
	maxB, _ := strconv.Atoi(match[5])
	return rule{
		minA: minA,
		maxA: maxA,
		minB: minB,
		maxB: maxB,
	}
}

func parseRules(input string) []rule {
	rules := []rule{}
	for _, r := range strings.Split(input, "\n") {
		rules = append(rules, parseRule(r))
	}
	return rules
}

func parseMyTicket(input string) ticket {
	parts := strings.Split(input, "\n")
	return parseTicket(parts[1])
}

func parseTicket(input string) ticket {
	t := ticket{}
	for _, f := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(f)
		t.fields = append(t.fields, i)
	}
	return t
}

func parseNearbyTickets(input string) []ticket {
	nearbyTickets := []ticket{}
	parts := strings.Split(input, "\n")
	for _, ticket := range parts[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(ticket))
	}
	return nearbyTickets
}

func fieldObeysRule(f int, r rule) bool {
	return (r.minA <= f && f <= r.maxA) || (r.minB <= f && f <= r.maxB)
}

func isInvalidForAnyRule(field int, rules []rule) bool {
	for _, rule := range rules {
		if fieldObeysRule(field, rule) {
			//fmt.Printf("field %d volates rule %+v\n", field, rule)
			return false
		}
	}
	return true
}

func calculateSumOfInvalidFields(t ticket, rules []rule) int {
	counter := 0
	for _, field := range t.fields {
		if isInvalidForAnyRule(field, rules) {
			//fmt.Printf("Ticket %+v field %d violates rules %+v\n", t, field, rules)
			counter += field
		}
	}
	return counter
}

func partOne(input string) int {
	counter := 0
	rules, _, nearbyTickets := parseInput(input)
	for _, ticket := range nearbyTickets {
		counter += calculateSumOfInvalidFields(ticket, rules)
	}
	return counter
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)
}
