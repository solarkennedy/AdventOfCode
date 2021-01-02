package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type rule struct {
	name string
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
	minA := utils.Atoi(match[2])
	maxA := utils.Atoi(match[3])
	minB := utils.Atoi(match[4])
	maxB := utils.Atoi(match[5])
	return rule{
		name: match[1],
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
	input = strings.TrimSpace(input)
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

func partTwo(input string) int {
	rules, myTicket, nearbyTickets := parseInput(input)
	validTickets := filterOutInvalidTickets(nearbyTickets, rules)

	fieldMappingPossibilities := map[string][]int{}
	for _, rule := range rules {
		possibleFields := figureOutPotentialFieldsForRule(rule, validTickets)
		fieldMappingPossibilities[rule.name] = possibleFields
	}
	fieldMapping := deduceFieldMappingFromPossibilities(fieldMappingPossibilities)

	printTicket(myTicket, fieldMapping)
	destinationFields := getDestinationFields(myTicket, fieldMapping)
	return multiplyFields(destinationFields, myTicket.fields)
}

func filterOutInvalidTickets(tickets []ticket, rules []rule) []ticket {
	validTickets := []ticket{}
	for _, ticket := range tickets {
		if calculateSumOfInvalidFields(ticket, rules) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func getAllValuesForField(f int, tickets []ticket) []int {
	values := []int{}
	for n, t := range tickets {
		if len(t.fields) <= f {
			panic(fmt.Errorf("This ticket #%d: %+v doesn't have field number %d", n, t, f))
		}
		values = append(values, t.fields[f])
	}
	fmt.Printf("All values for field %d are: %+v\n", f, values)
	return values
}

func isRuleValidForAllInputValues(r rule, values []int) bool {
	for _, v := range values {
		if !fieldObeysRule(v, r) {
			return false
		}
	}
	return true
}

func figureOutPotentialFieldsForRule(r rule, tickets []ticket) []int {
	numberOfPotentialFields := len(tickets[0].fields)
	possibleFields := []int{}
	for f := 0; f < numberOfPotentialFields; f++ {
		inputValues := getAllValuesForField(f, tickets)
		if isRuleValidForAllInputValues(r, inputValues) {
			fmt.Printf("Rule for %s looks valid for field number %d\n", r.name, f)
			possibleFields = append(possibleFields, f)
		}
	}
	if len(possibleFields) == 0 {
		panic(fmt.Errorf("Rule for %s doesn't look valid for any field?", r.name))
	}
	return possibleFields
}

func intSliceContains(haystack []int, needle int) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func removeFromSlice(s int, slice []int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func findPosition(value int, slice []int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func removePossibilityFromEverything(p int, possibilities map[string][]int) {
	for k, v := range possibilities {
		if intSliceContains(v, p) {
			position := findPosition(p, possibilities[k])
			possibilities[k] = removeFromSlice(position, possibilities[k])
		}
	}
}

func deduceFieldMappingFromPossibilities(possibilities map[string][]int) map[string]int {
	numberOfFields := len(possibilities)
	fieldMapping := map[string]int{}
	counter := 0
	for len(fieldMapping) < numberOfFields {

		for f, p := range possibilities {
			if len(p) == 1 {
				fmt.Printf("It looks like the only possible place for %s to go is field number %d!\n", f, p[0])
				panicIfFieldMappingIsAlreadyTaken(fieldMapping, f)
				fieldMapping[f] = p[0]
				removePossibilityFromEverything(p[0], possibilities)
				delete(possibilities, f)
			}
		}

		counter += 1
		if counter > 1000 {
			fmt.Printf("We tried this a thousand times and couldn't figure out the whole mapping. Here is what is left: %+v\n", possibilities)
			return fieldMapping
		}
	}
	return fieldMapping
}

func panicIfFieldMappingIsAlreadyTaken(fieldMapping map[string]int, f string) {
	value, ok := fieldMapping[f]
	if ok {
		panic(fmt.Errorf("But field %d is already assigned to %s?!", value, f))
	}
}

func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func printTicket(t ticket, mapping map[string]int) {
	rMapping := reverseMap(mapping)
	for i, f := range t.fields {
		fieldName := rMapping[i]
		fmt.Printf("%s: %d, ", fieldName, f)
	}
	fmt.Println()
}

func multiplyFields(fields []int, myValues []int) int {
	product := 1
	for _, i := range fields {
		product *= myValues[i]
	}
	return product
}

func getDestinationFields(myTicket ticket, fieldMapping map[string]int) []int {
	fields := []int{}
	for k, v := range fieldMapping {
		if strings.HasPrefix(k, "departure") {
			fields = append(fields, v)
		}
	}
	return fields
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
