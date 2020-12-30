package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	input := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

	actual := partOne(input)
	assert.Equal(t, 71, actual)
}

func Test_parseRule(t *testing.T) {
	actual := parseRule(`class: 1-3 or 5-7`)
	expected := rule{
		minA: 1,
		maxA: 3,
		minB: 5,
		maxB: 7,
	}
	assert.Equal(t, expected, actual)
}

func Test_parseMyTicket(t *testing.T) {
	input := `your ticket:
7,1,14`
	actual := parseTicket(input)
	expected := ticket{
		fields: []int{7, 1, 14},
	}
	assert.Equal(t, expected, actual)
}
