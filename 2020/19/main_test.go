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

func Test_partTwo(t *testing.T) {
	input := `42: 9 14 | 10 1
	9: 14 27 | 1 26
	10: 23 14 | 28 1
	1: "a"
	11: 42 31
	5: 1 14 | 15 1
	19: 14 1 | 14 14
	12: 24 14 | 19 1
	16: 15 1 | 14 14
	31: 14 17 | 1 13
	6: 14 14 | 1 14
	2: 1 24 | 14 4
	0: 8 11
	13: 14 3 | 1 12
	15: 1 | 14
	17: 14 2 | 1 7
	23: 25 1 | 22 14
	28: 16 1
	4: 1 1
	20: 14 14 | 1 15
	3: 5 14 | 16 1
	27: 1 6 | 14 18
	14: "b"
	21: 14 1 | 1 14
	25: 1 1 | 1 14
	22: 14 14
	8: 42
	26: 14 22 | 1 20
	18: 15 15
	7: 14 5 | 1 21
	24: 14 1

	abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
	bbabbbbaabaabba
	babbbbaabbbbbabbbbbbaabaaabaaa
	aaabbbbbbaaaabaababaabababbabaaabbababababaaa
	bbbbbbbaaaabbbbaaabbabaaa
	bbbababbbbaaaaaaaabbababaaababaabab
	ababaaaaaabaaab
	ababaaaaabbbaba
	baabbaaaabbaaaababbaababb
	abbbbabbbbaaaababbbbbbaaaababb
	aaaaabbaabaaaaababaa
	aaaabbaaaabbaaa
	aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
	babaaabbbaaabaababbaabababaaab
	aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`
	assert.Equal(t, 12, partTwo(input))
}

func Test_partTwoSingleMessage(t *testing.T) {
	input := `0: 8 11
	1: "a"
	2: 1 24 | 14 4
	3: 5 14 | 16 1
	4: 1 1
	5: 1 14 | 15 1
	6: 14 14 | 1 14
	7: 14 5 | 1 21
	8: 42
	9: 14 27 | 1 26
	10: 23 14 | 28 1
	11: 42 31
	12: 24 14 | 19 1
	13: 14 3 | 1 12
	14: "b"
	15: 1 | 14
	16: 15 1 | 14 14
	17: 14 2 | 1 7
	18: 15 15
	19: 14 1 | 14 14
	20: 14 14 | 1 15
	21: 14 1 | 1 14
	22: 14 14
	23: 25 1 | 22 14
	24: 14 1
	25: 1 1 | 1 14
	26: 14 22 | 1 20
	27: 1 6 | 14 18
	28: 16 1
	31: 14 17 | 1 13
	42: 9 14 | 10 1`
	rules := parseRules(input)
	modifyRulesForPart2(rules)
	assert.True(t, messageMatchesRule0("babbbbaabbbbbabbbbbbaabaaabaaa", rules))
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
