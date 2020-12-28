package main

import (
	"testing"

	"github.com/solarkennedy/AdventOfCode/utils"
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

func BenchmarkGetPartTwo(b *testing.B) {
	input := utils.ReadInput()
	instructions := parseInstructions(input)
	b.Run("part2-serial__	", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = partTwo(instructions)
		}
	})
	b.Run("part2-parallel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = partTwoParallel(instructions)
		}
	})
}
