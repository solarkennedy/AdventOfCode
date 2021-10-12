package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partTwo(t *testing.T) {
	assert.Equal(t, 8, partTwo(basicTestInput))
	assert.Equal(t, 19208, partTwo(secondTestInput))
}
