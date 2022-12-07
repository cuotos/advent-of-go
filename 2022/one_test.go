package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOne(t *testing.T) {
	testInput := []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

1`)

	// part1
	output := day1Pt1(testInput)
	assert.Equal(t, 24000, output)

	// part2
	output = day1Pt2(testInput)
	assert.Equal(t, 45000, output)
}
