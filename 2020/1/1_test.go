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
	actual := one(input)
	assert.Equal(t, expected, actual)
}

func Benchmark_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		one(input)
	}
}
