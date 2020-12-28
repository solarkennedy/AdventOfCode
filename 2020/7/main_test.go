package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRule(t *testing.T) {
	var actual1 string
	var actual2 []string

	actual1, actual2 = parseRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	assert.Equal(t, "light red", actual1)
	assert.Equal(t, []string{"bright white", "muted yellow"}, actual2)

	actual1, actual2 = parseRule("bright white bags contain 1 shiny gold bag.")
	assert.Equal(t, "bright white", actual1)
	assert.Equal(t, []string{"shiny gold"}, actual2)

	actual1, actual2 = parseRule("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags, 2 pale bronze bags.")
	assert.Equal(t, "muted yellow", actual1)
	assert.Equal(t, []string{"shiny gold", "faded blue", "pale bronze"}, actual2)

	actual1, actual2 = parseRule("dotted black bags contain no other bags.")
	assert.Equal(t, "dotted black", actual1)
	assert.Equal(t, []string{}, actual2)
}

func Test_partOne(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	assert.Equal(t, 4, partOne(input))
}
