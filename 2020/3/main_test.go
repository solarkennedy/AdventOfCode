package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testMap = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
)

func TestPartOne(t *testing.T) {
	expected := 7
	actual := partOne(inputToMap(testMap))
	assert.Equal(t, expected, actual)
}

func TestPartTwo(t *testing.T) {
	snowMap := inputToMap(testMap)
	assert.Equal(t, 2, runSlope(snowMap, 1, 1))
	assert.Equal(t, 3, runSlope(snowMap, 5, 1))

	expected := 336
	actual := partTwo(inputToMap(testMap))
	assert.Equal(t, expected, actual)
}
