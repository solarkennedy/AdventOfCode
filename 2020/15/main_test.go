package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	assert.Equal(t, 1, partOne("1,3,2"))
	assert.Equal(t, 10, partOne("2,1,3"))
	assert.Equal(t, 27, partOne("1,2,3"))
	assert.Equal(t, 78, partOne("2,3,1"))
	assert.Equal(t, 438, partOne("3,2,1"))
	assert.Equal(t, 1836, partOne("3,1,2"))
}

func Test_findLastTimeSpoken(t *testing.T) {
	assert.Equal(t, 0, findLastTimeSpoken(0, []int{0, 3, 6}))
}

func Test_turn(t *testing.T) {
	assert.Equal(t, 0, turn([]int{0, 3, 6, -1}, 4))
	assert.Equal(t, 3, turn([]int{0, 3, 6, 0, -1}, 5))
	assert.Equal(t, 3, turn([]int{0, 3, 6, 0, 3, -1}, 6))
	assert.Equal(t, 1, turn([]int{0, 3, 6, 0, 3, 3, -1}, 7))
	assert.Equal(t, 0, turn([]int{0, 3, 6, 0, 3, 3, 1, -1}, 8))
	assert.Equal(t, 4, turn([]int{0, 3, 6, 0, 3, 3, 1, 0, -1}, 9))
	assert.Equal(t, 0, turn([]int{0, 3, 6, 0, 3, 3, 1, 0, 4, -1}, 10))
}

func Test_turnExample(t *testing.T) {
	assert.Equal(t, 436, partOne("0,3,6"))
}
