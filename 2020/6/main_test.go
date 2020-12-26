package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAnswers(t *testing.T) {
	input := `ab
ac`
	assert.Equal(t, 3, len(countAnswers(input)))
}

func TestPartOne(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	actual := partOne(input)
	assert.Equal(t, 11, actual)
}

func Test_countCommonYesAnswers(t *testing.T) {
	input := `ab
ac`
	answers := countAnswers(input)
	assert.Equal(t, 1, countCommonYesAnswers(answers, 2))
}

func TestPartTwo(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	actual := partTwo(input)
	assert.Equal(t, 6, actual)
}
