package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const basicTestInput = `
16
10
15
5
1
11
7
19
6
12
4`

const secondTestInput = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func Test_partOne(t *testing.T) {
	assert.Equal(t, 35, partOne(basicTestInput))
	assert.Equal(t, 220, partOne(secondTestInput))
}

func Test_removeAdapter(t *testing.T) {
	a := parseAdapters("1\n2\n3")
	assert.Equal(t, 3, len(a))
	b := removeAdapter(a, 0)
	assert.Equal(t, 3, len(a))
	assert.Equal(t, 2, len(b))
	c := removeAdapter(b, 0)
	assert.Equal(t, 3, len(a))
	assert.Equal(t, 2, len(b))
	assert.Equal(t, 1, len(c))
	d := removeAdapter(c, 0)
	assert.Equal(t, 3, len(a))
	assert.Equal(t, 2, len(b))
	assert.Equal(t, 1, len(c))
	assert.Equal(t, 0, len(d))
}

func Test_removeAdapterObject(t *testing.T) {
	a := parseAdapters("1\n2\n3")
	assert.Equal(t, 3, len(a))
	ch := AdapterChain{unused: a}
	b := removeAdapter(ch.unused, 0)
	assert.Equal(t, 3, len(ch.unused))
	assert.Equal(t, 1, ch.unused[0])
	assert.Equal(t, 2, len(b))
	assert.Equal(t, 2, b[0])

}

func Test_countJoltDifferences(t *testing.T) {
	a := parseAdapters("1\n4\n5\n6\n7\n10\n11\n12\n15\n16\n19")
	ones := countJoltDifferences(a, 1)
	assert.Equal(t, 7, ones)
	threes := countJoltDifferences(a, 3)
	assert.Equal(t, 5, threes)
}
