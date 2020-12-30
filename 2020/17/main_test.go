package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	input := `.#.
..#
###`
	initialGrid := parseGrid(input)
	actual := partOne(initialGrid)
	assert.Equal(t, 112, actual)
}

func Test_partTwo(t *testing.T) {
	input := `.#.
..#
###`
	initialGrid := parseGrid(input)
	initialGrid4 := place3to4(initialGrid)
	actual := partTwo(initialGrid4)
	assert.Equal(t, 848, actual)
}

func Test_parseAndRender(t *testing.T) {
	input := `.#.
..#
###`
	grid := parseGrid(input)
	actual := renderSlice(grid, 1, 0)
	fmt.Printf("+%v\n", grid)
	assert.Equal(t, "z=0\n"+input+"\n", actual)

}
func Test_isNeighbor(t *testing.T) {
	cube := vector3{x: 0, y: 0, z: 0}
	assert.True(t, cube.isNeighborTo(vector3{x: 1, y: 1, z: 1}))
	assert.True(t, cube.isNeighborTo(vector3{x: 0, y: 0, z: 1}))

	assert.False(t, cube.isNeighborTo(vector3{x: 3, y: 3, z: 3}))
	assert.False(t, cube.isNeighborTo(vector3{x: 2, y: 2, z: 2}))
	assert.False(t, cube.isNeighborTo(vector3{x: 0, y: 2, z: 2}))
}

func Test_isActive(t *testing.T) {
	grid := []vector3{{x: 1, y: 1, z: 1}}
	c1 := vector3{x: 1, y: 1, z: 1}
	assert.True(t, isActive(c1, grid))
	c2 := vector3{x: 0, y: 0, z: 1}
	assert.False(t, isActive(c2, grid))
}

func Test_countNeighbors(t *testing.T) {
	grid := []vector3{{x: 1, y: 1, z: 1}}
	actual := countNeighbors(vector3{x: 1, y: 1, z: 1}, grid)
	assert.Equal(t, 0, actual)

	grid = []vector3{
		{x: 1, y: 1, z: 1},
		{x: 1, y: 1, z: 0},
		{x: 0, y: 0, z: 0},
	}
	actual = countNeighbors(vector3{x: 0, y: 0, z: 0}, grid)
	assert.Equal(t, 2, actual)
}

func Test_problemSpace(t *testing.T) {
	actual := problemSpace(2)
	assert.Equal(t, 125, len(actual))
	first := vector3{x: -2, y: -2, z: -2}
	assert.Equal(t, first, actual[0])
}

func Test_cycle(t *testing.T) {
	input := `.#.
..#
###`
	grid := parseGrid(input)
	assert.Equal(t, 5, len(grid))

	grid = cycle(grid, 1)
	assert.Equal(t, 11, len(grid))
}
