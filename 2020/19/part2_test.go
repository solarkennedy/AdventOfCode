package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	regex := rulesToRegexp(rules)
	assert.True(t, messageMatchesRegexRule0("babbbbaabbbbbabbbbbbaabaaabaaa", regex))
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

func Test_ruleToRegex(t *testing.T) {
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
	24: 14 1`
	rules := parseRules(input)
	modifyRulesForPart2(rules)
	assert.Equal(t, ruleToRegex(rules[14], rules), `(?:b)`)
	assert.Equal(t, ruleToRegex(rules[24], rules), `(?:(?:b)(?:a))`)
	assert.Equal(t, ruleToRegex(rules[0], rules), `(?:(?:b)(?:a))`)

}
