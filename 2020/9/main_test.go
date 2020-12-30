package main

import (
	"testing"

	"github.com/solarkennedy/AdventOfCode/utils"
	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	ints := utils.ConvertIntoInts(input)
	actual := findFirstInvalidNumber(ints, 5)
	assert.Equal(t, 127, actual)
}

func Test_isValid(t *testing.T) {
	preamble := make([]int, 25)
	for i := range preamble {
		preamble[i] = i + 1
	}

	assert.True(t, isValid(26, preamble))
	assert.True(t, isValid(49, preamble))
	assert.False(t, isValid(100, preamble))
	assert.False(t, isValid(50, preamble))
}

func Test_rangeSumsTo(t *testing.T) {
	r := []int{15, 25, 47, 40}
	assert.True(t, rangeSumsTo(127, r))
}

func Test_partTwo(t *testing.T) {
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	ints := utils.ConvertIntoInts(input)
	invalid := findFirstInvalidNumber(ints, 5)

	actualBlock := findContinguousBlockThatAddsTo(invalid, ints)
	expectedBlock := []int{15, 25, 47, 40}
	assert.Equal(t, expectedBlock, actualBlock)

	actual := partTwo(ints, invalid)
	assert.Equal(t, 62, actual)
}
