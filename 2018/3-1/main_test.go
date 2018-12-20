package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClaimStringParsing(t *testing.T) {
	tcs := []struct {
		Input    string
		Expected Claim
	}{
		{
			"#217 @ 1,1: 10x10",
			Claim{217, 1, 1, 10, 10},
		},
		{
			"#123 @ 4,5: 67x89",
			Claim{123, 4, 5, 67, 89},
		},
	}

	for _, tc := range tcs {
		actual := parseLine(tc.Input)

		assert.Equal(t, tc.Expected, actual)
	}
}

// We have the correct answer for the input data. just make sure we haven't broken it
func TestFindOverLapsInRealInputDate(t *testing.T) {
	expected := 96569
	actual := findOverlaps("input")

	assert.Equal(t, expected, actual)
}
