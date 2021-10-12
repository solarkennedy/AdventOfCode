package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	assert.Equal(t, 2, partOne(`0: 4 1 5
	1: 2 3 | 3 2
	2: 4 4 | 5 5
	3: 4 5 | 5 4
	4: "a"
	5: "b"

	ababbb
	bababa
	abbbab
	aaabbb
	aaaabbb`))
}


func Test_messageMatchesRule0(t *testing.T) {
	r := parseRules(`0: 4 1 5
	1: 2 3 | 3 2
	2: 4 4 | 5 5
	3: 4 5 | 5 4
	4: "a"
	5: "b"`)
	assert.True(t, messageMatchesRule0(`ababbb`, r))
	assert.True(t, messageMatchesRule0(`abbbab`, r))
	assert.False(t, messageMatchesRule0(`bababa`, r))
	assert.False(t, messageMatchesRule0(`aaabbb`, r))
	assert.False(t, messageMatchesRule0(`aaaabbb`, r))
}

func Test_parseRule(t *testing.T) {
	r := parseRule("0: 4 1 5")
	assert.Equal(t, []int{4, 1, 5}, r.ruleSet1)
	r = parseRule("1: 2 3 | 3 2")
	assert.Equal(t, []int{2, 3}, r.ruleSet1)
	assert.Equal(t, []int{3, 2}, r.ruleSet2)
	r = parseRule(`5: "b"`)
	assert.Equal(t, "b", r.mainRule)
}

func Test_parseRulesMixedInputs(t *testing.T) {
	input := `0: 2 3 | 3 2
	4: 4 4 | 5 5
	3: 4 5 | 5 4
	2: "a"
	1: "b"`
	actualRules := parseRules(input)
	assert.Equal(t, 1, actualRules[1].number)
}
