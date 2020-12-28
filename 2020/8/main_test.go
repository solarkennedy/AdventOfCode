package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	instructions := parseInstructions(input)
	actual, _ := partOne(instructions)
	assert.Equal(t, 5, actual)
}

func Test_partTwo(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	instructions := parseInstructions(input)
	actual, _ := partTwo(instructions)
	assert.Equal(t, 8, actual)
}
