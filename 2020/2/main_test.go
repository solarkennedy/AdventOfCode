package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
)

func TestParseEntryLine(t *testing.T) {
	input := "1-3 a: abcde"
	p := passwordPolicy{
		first:  1,
		second: 3,
		letter: 'a',
	}
	expected := entry{
		policy:   p,
		password: "abcde",
	}
	actual := parseEntryLine(input)
	assert.Equal(t, expected, actual)
}

func TestParsePolicy(t *testing.T) {
	input := "1-3 a"
	expected := passwordPolicy{
		first:  1,
		second: 3,
		letter: 'a',
	}
	actual := parsePolicy(input)
	assert.Equal(t, expected, actual)
}

func TestPartOne(t *testing.T) {
	expected := 2
	actual := partOne(inputToEntries(input))
	assert.Equal(t, expected, actual)
}

func TestIsValidPasswordPolicyPart2Good(t *testing.T) {
	input := "1-3 a: abcde"
	entry := parseEntryLine(input)
	actual := isValidPasswordPart2(entry.policy, entry.password)
	assert.True(t, actual)
}

func TestIsValidPasswordPolicyPart2Bad(t *testing.T) {
	input := "1-3 b: cdefg"
	entry := parseEntryLine(input)
	actual := isValidPasswordPart2(entry.policy, entry.password)
	assert.False(t, actual)
}

func TestPartTwo(t *testing.T) {
	expected := 1
	actual := partTwo(inputToEntries(input))
	assert.Equal(t, expected, actual)
}
