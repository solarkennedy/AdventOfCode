package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testBatch = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
)

func TestParsePassport(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`
	passport := parsePassport(input)
	assert.NotNil(t, passport)
	assert.Equal(t, "2020", passport.Eyr)
}

func TestAttributeKVToPassport(t *testing.T) {
	p := passport{}
	attributeKVToPassport(&p, "eyr", "42")
	assert.Equal(t, p.Eyr, "42")
}

func TestIsValidPassportGood(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`
	p := parsePassport(input)
	isValid, err := isValidPassport(p)
	// "all eight fields are present"
	assert.True(t, isValid)
	assert.Nil(t, err)
}

func TestIsValidPassportBad(t *testing.T) {
	input := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`
	p := parsePassport(input)
	isValid, err := isValidPassport(p)
	// "it is missing hgt"
	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "missing Hgt")
}
func TestPartOne(t *testing.T) {
	passports := parseBatchInput(testBatch)
	actual := partOne(passports)
	assert.Equal(t, 2, actual)
}
