package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	input = `1721
979
366
299
675
1456`
)

func Test_1(t *testing.T) {

	expected := 514579
	actual := one(inputToIntSlice(input))
	assert.Equal(t, expected, actual)
}

func Benchmark_1(b *testing.B) {
	inputSlice := inputToIntSlice(input)
	for i := 0; i < b.N; i++ {
		one(inputSlice)
	}
}

func Test_1PartTwo(t *testing.T) {
	expected := 241861950
	actual := onePartTwo(inputToIntSlice(input))
	assert.Equal(t, expected, actual)
}

func Benchmark_1Part2(b *testing.B) {
	inputSlice := inputToIntSlice(input)
	for i := 0; i < b.N; i++ {
		onePartTwo(inputSlice)
	}
}
