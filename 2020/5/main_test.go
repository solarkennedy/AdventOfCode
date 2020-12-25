package main

import (
	"testing"

	"github.com/solarkennedy/AdventOfCode/utils"
	"github.com/stretchr/testify/assert"
)

func TestSeatToID(t *testing.T) {
	assert.Equal(t, 567, seatToID("BFFFBBFRRR"))
	assert.Equal(t, 119, seatToID("FFFBBBFRRR"))
	assert.Equal(t, 820, seatToID("BBFFBBFRLL"))
}

func TestPartOne(t *testing.T) {
	input := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
	ids := seatsToIDs(input)
	actual := partOne(ids)
	assert.Equal(t, 820, actual)
}

func BenchmarkGetMax(b *testing.B) {
	input := utils.ReadInput()
	ids := seatsToIDs(input)
	b.Run("max", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getMax(ids)
		}
	})
}

func BenchmarkGetPartTwo(b *testing.B) {
	input := utils.ReadInput()
	ids := seatsToIDs(input)
	b.Run("part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			partTwo(ids)
		}
	})
}
