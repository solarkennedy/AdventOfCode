package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	input := `
	2 * 3 + (4 * 5)
    5 + (8 * 3 + 9 + 3 * 4 * 3)
    5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
    ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`
	actual := partOne(input)
	assert.Equal(t, 26335, actual)
}

func Test_tokenize(t *testing.T) {
	var line string
	var expected []string

	line = "2 * 3"
	expected = []string{"2", "*", "3"}
	assert.Equal(t, expected, tokenize(line))

	line = "2 * 55"
	expected = []string{"2", "*", "55"}
	assert.Equal(t, expected, tokenize(line))
}

func Test_isNumber(t *testing.T) {
	assert.False(t, isNumber('a'))
	assert.False(t, isNumber('('))
	assert.True(t, isNumber('0'))
}

func Test_nextRuneIsNumber(t *testing.T) {
	assert.True(t, nextRuneIsNumber("2 * 55", 4))
	assert.False(t, nextRuneIsNumber("2 * 55", 5))
	assert.False(t, nextRuneIsNumber("(2 * 55)", 6))
}

func Test_findCloseParen(t *testing.T) {
	line := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	tokens := tokenize(line)
	expected := 20
	assert.Equal(t, expected, findCloseParen(tokens))
}

func Test_findCloseParen2(t *testing.T) {
	line := "(2 + 4 * 9) * (6 + 9 * 8 + 6)"
	tokens := tokenize(line)
	expected := 6
	assert.Equal(t, expected, findCloseParen(tokens))
}

func Test_evaulateLine1(t *testing.T) {
	line := "2 * 3"
	assert.Equal(t, 6, evaluate(tokenize(line)))
}
func Test_evaulateLine2(t *testing.T) {
	line := "(4 * 5)"
	assert.Equal(t, 20, evaluate(tokenize(line)))
}

func Test_evaulateLineLeftToRightSimple(t *testing.T) {
	line := "1 + 2 * 3"
	assert.Equal(t, 9, evaluate(tokenize(line)))
}
func Test_evaulateLine3(t *testing.T) {
	line := "2 * 3 + (4 * 5)"
	assert.Equal(t, 26, evaluate(tokenize(line)))
}
func Test_evaulateLine4(t *testing.T) {
	line := "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	assert.Equal(t, 437, evaluate(tokenize(line)))
}
func Test_evaulateLine5(t *testing.T) {
	line := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	assert.Equal(t, 12240, evaluate(tokenize(line)))
}
func Test_evaulateLine6(t *testing.T) {
	line := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	assert.Equal(t, 13632, evaluate(tokenize(line)))
}

func Test_partTwoEvaluation(t *testing.T) {
	var line string
	var expected int

	line = "1 + (2 * 3) + (4 * (5 + 6))"
	expected = 51
	assert.Equal(t, expected, evaluate2(tokenize(line)))

	line = "2 * 3 + (4 * 5)"
	expected = 46
	assert.Equal(t, expected, evaluate2(tokenize(line)))

	line = "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	expected = 1445
	assert.Equal(t, expected, evaluate2(tokenize(line)))

	line = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	expected = 669060
	assert.Equal(t, expected, evaluate2(tokenize(line)))

	line = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 "
	expected = 23340
	assert.Equal(t, expected, evaluate2(tokenize(line)))
}
