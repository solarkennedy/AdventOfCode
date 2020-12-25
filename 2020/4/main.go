package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func splitPassportInputs(input string) []string {
	chunks := []string{}
	for _, chunk := range strings.Split(input, "\n\n") {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func parseKV(kv string) (string, string) {
	v := strings.Split(kv, ":")
	return v[0], v[1]
}

// attributeKVToPassport dynamically takes a passport pointer and a k,v pair
// then assigns that to the struct. For example, if k:foo and v:bar
// it would run the equivilant go code to p.foo = "bar"
// This helps prevent hard-coding struct names as the format changes,
// but requires reflect magic.
// Everything I needed to write this function came from
// https://blog.golang.org/laws-of-reflection#TOC_10.
func attributeKVToPassport(p *passport, k string, v string) {
	rv := reflect.ValueOf(p).Elem()
	typeOfS := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		if k == strings.ToLower(fieldName) {
			fv := rv.Field(i)
			fv.SetString(v)
		}
	}
}

func parsePassport(input string) passport {
	fmt.Println(input)
	p := passport{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for _, kv := range strings.Split(line, " ") {
			k, v := parseKV(kv)
			attributeKVToPassport(&p, k, v)
		}
	}
	return p
}

func parseBatchInput(input string) []passport {
	passports := []passport{}
	for _, passportStringInput := range splitPassportInputs(input) {
		passport := parsePassport(passportStringInput)
		passports = append(passports, passport)
	}
	return passports
}

// isValidPassport validates by verifying that all fields
// are set, *except* `cid`, which is OK to be unset
// because we are dirty
func isValidPassport(p passport) (bool, error) {
	rv := reflect.ValueOf(&p).Elem()
	typeOfS := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		fv := rv.Field(i)
		if fv.String() == "" && fieldName != "Cid" {
			return false, fmt.Errorf("Passport is missing %s: %+v", fieldName, p)
		}
	}
	return true, nil
}

func partOne(passports []passport) int {
	counter := 0
	for _, passport := range passports {
		isValid, err := isValidPassport(passport)
		if isValid {
			counter += 1
		} else {
			fmt.Printf("Invalid passport: %s\n", err)
		}
	}
	return counter
}

func partTwo(passports []passport) int {
	return 0
}

func main() {
	input := utils.ReadInput()
	passports := parseBatchInput(input)
	result := partOne(passports)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(passports)
	fmt.Printf("Answer to part two: %d\n", result2)
}
