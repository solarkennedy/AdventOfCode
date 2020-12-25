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

func TestIsValidPassportPartOneGood(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`
	p := parsePassport(input)
	isValid, err := isValidPassportPartOne(p)
	// "all eight fields are present"
	assert.True(t, isValid)
	assert.Nil(t, err)
}

func TestIsValidPassportPartOneBad(t *testing.T) {
	input := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`
	p := parsePassport(input)
	isValid, err := isValidPassportPartOne(p)
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

func TestPartTwoInvalidExamples(t *testing.T) {
	invalidPassports := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
`
	passports := parseBatchInput(invalidPassports)
	assert.Equal(t, 4, len(passports))
	actual := partTwo(passports)
	assert.Equal(t, 0, actual)
}

func TestPartTwoValidExamples(t *testing.T) {
	validPassports := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	passports := parseBatchInput(validPassports)
	assert.Equal(t, 4, len(passports))
	actual := partTwo(passports)
	assert.Equal(t, 4, actual)
}

func TestPartTwoValidations(t *testing.T) {
	var ok bool

	ok, _ = validateByr(passport{Byr: "2000"})
	assert.True(t, ok)
	ok, _ = validateByr(passport{Byr: "1900"})
	assert.False(t, ok)

	ok, _ = validateIyr(passport{Iyr: "2019"})
	assert.True(t, ok)
	ok, _ = validateIyr(passport{Iyr: "1990"})
	assert.False(t, ok)

	ok, _ = validateEyr(passport{Eyr: "2022"})
	assert.True(t, ok)
	ok, _ = validateEyr(passport{Eyr: "1990"})
	assert.False(t, ok)

	ok, _ = validateHgt(passport{Hgt: "149cm"})
	assert.False(t, ok)
	ok, _ = validateHgt(passport{Hgt: "58in"})
	assert.False(t, ok)
	ok, _ = validateHgt(passport{Hgt: "foo"})
	assert.False(t, ok)

	ok, _ = validateHcl(passport{Hcl: "#ffffff"})
	assert.True(t, ok)
	ok, _ = validateHcl(passport{Hcl: "white"})
	assert.False(t, ok)

	ok, _ = validateEcl(passport{Ecl: "blu"})
	assert.True(t, ok)
	ok, _ = validateEcl(passport{Ecl: "white"})
	assert.False(t, ok)

	ok, _ = validatePid(passport{Pid: "000111222"})
	assert.True(t, ok)
	ok, _ = validatePid(passport{Pid: "123"})
	assert.False(t, ok)

}
