package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	assert.Equal(t, 42, partOne("1,3,2"))
}

func Test_partTwo(t *testing.T) {
	assert.Equal(t, 43, partTwo("1,2,3"))
}

