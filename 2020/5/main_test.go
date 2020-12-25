package main

import (
	"testing"

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
	actual := partOne(input)
	assert.Equal(t, 820, actual)
}
