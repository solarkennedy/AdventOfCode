package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const basicTestInput = `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const part2Input = `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func Test_partOne(t *testing.T) {
	assert.Equal(t, 165, partOne(basicTestInput))
}

func Test_computePostMaskValue(t *testing.T) {
	assert.Equal(t, 73, computePostMaskValue(11, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))
	assert.Equal(t, 101, computePostMaskValue(101, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))
	assert.Equal(t, 64, computePostMaskValue(0, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))
}

func Test_partTwo(t *testing.T) {
	assert.Equal(t, 208, partTwo(part2Input))
}

func Test_parseSingleInstructionAddressDecode(t *testing.T) {
	ins := parseInstructions(`mask = 000000000000000000000000000000X1001X
	mem[42] = 100`)
	actual := parseSingleInstructionAddressDecode(ins[0])
	assert.Equal(t, "000000000000000000000000000000X1101X", actual)
}
