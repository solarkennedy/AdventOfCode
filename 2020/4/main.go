package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
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
	return strings.Split(input, "\n\n")
}

func parseKV(kv string) (string, string) {
	v := strings.Split(kv, ":")
	if len(v) != 2 {
		fmt.Printf("Couldn't parse the the : kv: '%s'", kv)
		panic(v)
	}
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
	p := passport{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for _, kv := range strings.Split(line, " ") {
			if kv == "" || kv == "\t" {
				continue
			}
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

// isValidPassportPartOne validates by verifying that all fields
// are set, *except* `cid`, which is OK to be unset
// because we are dirty
func isValidPassportPartOne(p passport) (bool, error) {
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
		isValid, err := isValidPassportPartOne(passport)
		if isValid {
			counter += 1
		} else {
			fmt.Printf("Invalid passport: %s\n", err)
		}
	}
	return counter
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func validateByr(p passport) (bool, error) {
	y, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false, fmt.Errorf("Failed to parse byr: %w", err)
	}
	if y < 1920 || y > 2002 {
		return false, fmt.Errorf("byr '%d' out of range, must be between 1920 and 2002 inclusive", y)
	}
	return true, nil
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func validateIyr(p passport) (bool, error) {
	y, err := strconv.Atoi(p.Iyr)
	if err != nil {
		return false, fmt.Errorf("Failed to parse Iyr: %w", err)
	}
	if y < 2010 || y > 2020 {
		return false, fmt.Errorf("iyr '%d' out of range, must be between 2010 and 2020 inclusive", y)
	}
	return true, nil
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validateEyr(p passport) (bool, error) {
	y, err := strconv.Atoi(p.Eyr)
	if err != nil {
		return false, fmt.Errorf("Failed to parse Eyr: %w", err)
	}
	if y < 2020 || y > 2030 {
		return false, fmt.Errorf("eyr '%d' out of range, must be between 2020 and 2030 inclusive", y)
	}
	return true, nil
}

// hgt (Height) - a number followed by either cm or in:
// 	If cm, the number must be at least 150 and at most 193.
// 	If in, the number must be at least 59 and at most 76.
func validateHgt(p passport) (bool, error) {
	r := regexp.MustCompile(`(?P<Value>\d+)(?P<Unit>in|cm)`)
	//r := regexp.MustCompile(`(?P<Value>.*)`)
	match := r.FindStringSubmatch(p.Hgt)
	if len(match) != 3 {
		return false, fmt.Errorf("Input hgt: '%s' didn't match our regex '%s'", p.Hgt, r.String())
	}
	unit := match[2]
	value, err := strconv.Atoi(match[1])
	if err != nil {
		return false, fmt.Errorf("Could not parse value of height: '%s' for '%s'", err, p.Hgt)
	}
	if unit == "cm" {
		if value < 150 || value > 193 {
			return false, fmt.Errorf("Height must be between 150-193 cm: '%s'", p.Hgt)
		}
	} else if unit == "in" {
		if value < 59 || value > 76 {
			return false, fmt.Errorf("Height must be between 59-76 in: '%s'", p.Hgt)
		}
	} else {
		return false, fmt.Errorf("Unknown hight unit '%s' on '%s'", unit, p.Hgt)
	}
	return true, nil
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validateHcl(p passport) (bool, error) {
	r := regexp.MustCompile(`#[0-9a-f]{6}`)
	match := r.Match([]byte(p.Hcl))
	if !match {
		return false, fmt.Errorf("hcl '%s' didn't match our regex '%s'", p.Hcl, r.String())
	}
	return true, nil
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validateEcl(p passport) (bool, error) {
	r := regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	if !r.Match([]byte(p.Ecl)) {
		return false, fmt.Errorf("ecl '%s' didn't match our regex '%s'", p.Ecl, r.String())
	}
	return true, nil
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func validatePid(p passport) (bool, error) {
	r := regexp.MustCompile(`^\d{9}$`)
	match := r.MatchString(p.Pid)
	if !match {
		return false, fmt.Errorf("pid '%s' didn't match our regex '%s'", p.Pid, r.String())
	}
	return true, nil
}

// cid (Country ID) - ignored, missing or not.
func validateCid(p passport) (bool, error) {
	return true, nil
}

// isValidPassportPartTwo does the stricter validation per the rules:
func isValidPassportPartTwo(p passport) (bool, error) {
	if ok, err := validateByr(p); !ok {
		return false, err
	}
	if ok, err := validateIyr(p); !ok {
		return false, err
	}
	if ok, err := validateEyr(p); !ok {
		return false, err
	}
	if ok, err := validateHgt(p); !ok {
		return false, err
	}
	if ok, err := validateHcl(p); !ok {
		return false, err
	}
	if ok, err := validateEcl(p); !ok {
		return false, err
	}
	if ok, err := validatePid(p); !ok {
		return false, err
	}
	if ok, err := validateCid(p); !ok {
		return false, err
	}
	return true, nil
}

func partTwo(passports []passport) int {
	counter := 0
	for _, passport := range passports {
		isValid, err := isValidPassportPartTwo(passport)
		if isValid {
			counter += 1
		} else {
			fmt.Printf("Invalid passport: %s\n", err)
		}
	}
	return counter
}

func main() {
	input := utils.ReadInput()
	passports := parseBatchInput(input)
	result := partOne(passports)
	fmt.Printf("Answer to part one: %d\n", result)
	result2 := partTwo(passports)
	fmt.Printf("Answer to part two: %d\n", result2)
}
