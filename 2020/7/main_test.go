package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRule(t *testing.T) {
	var actual1 string
	var actual2 []rule
	var expectedRules []rule

	actual1, actual2 = parseRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	assert.Equal(t, "light red", actual1)
	expectedRules = []rule{{amount: 1, color: "bright white"}, {amount: 2, color: "muted yellow"}}
	assert.Equal(t, expectedRules, actual2)

	actual1, actual2 = parseRule("bright white bags contain 1 shiny gold bag.")
	assert.Equal(t, "bright white", actual1)
	expectedRules = []rule{{amount: 1, color: "shiny gold"}}
	assert.Equal(t, expectedRules, actual2)

	actual1, actual2 = parseRule("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags, 2 pale bronze bags.")
	assert.Equal(t, "muted yellow", actual1)
	expectedRules = []rule{{amount: 2, color: "shiny gold"}, {amount: 9, color: "faded blue"}, {amount: 2, color: "pale bronze"}}
	assert.Equal(t, expectedRules, actual2)

	actual1, actual2 = parseRule("dotted black bags contain no other bags.")
	assert.Equal(t, "dotted black", actual1)
	expectedRules = []rule{}
	assert.Equal(t, expectedRules, actual2)

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

func Test_partTwo(t *testing.T) {
	input := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
	assert.Equal(t, 126, partTwo(input))
}
