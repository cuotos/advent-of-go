package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsRange(t *testing.T) {
	tcs := []struct {
		RangeA     rng
		RangeB     rng
		BContainsA bool
	}{
		{
			rng{1, 10},
			rng{4, 6},
			true,
		},
		{
			rng{1, 5},
			rng{4, 7},
			false,
		},
		{
			rng{5, 6},
			rng{3, 9},
			true,
		},
	}

	for _, tc := range tcs {
		actual := rangeContains(tc.RangeA, tc.RangeB)
		assert.Equal(t, tc.BContainsA, actual)
	}
}

func TestRangeOverlap(t *testing.T) {
	tcs := []struct {
		RangeA     rng
		RangeB     rng
		BContainsA bool
	}{
		{
			rng{1, 10},
			rng{4, 20},
			true,
		},
		{
			rng{1, 3},
			rng{4, 7},
			false,
		},
		{
			rng{50, 600},
			rng{1, 100},
			true,
		},
		{
			rng{1, 4},
			rng{4, 7},
			true,
		},
		{
			rng{10, 12},
			rng{4, 11},
			true,
		},
	}

	for _, tc := range tcs {
		actual := rangesOverlap(tc.RangeA, tc.RangeB)
		assert.Equal(t, tc.BContainsA, actual)
	}
}

func TestParseRangeLine(t *testing.T) {
	tcs := []struct {
		input    string
		expected rng
	}{
		{"62-77", rng{62, 77}},
		{"1-5", rng{1, 5}},
	}

	for _, tc := range tcs {

		actual := parseRange(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestParse(t *testing.T) {
	tcs := []struct {
		input    []byte
		expected []Pair
	}{
		{
			[]byte("8-18,10-19\n12-69,8-15"),
			[]Pair{
				{
					rngA: rng{8, 18},
					rngB: rng{10, 19},
				},
				{
					rngA: rng{12, 69},
					rngB: rng{8, 15},
				},
			},
		},
	}

	for _, tc := range tcs {
		actual := parse(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestXxx(t *testing.T) {
	input := []byte(`8-18,10-19
	1-100,10-15
	12-69,8-15
	62-77,36-50
	26-27,26-91
	16-23,24-63
	17-43,18-44
5-10,6-9
	29-68,29-70
	15-90,28-91
	8-39,10-40`)

	actual := Run(input)
	assert.Equal(t, 2, actual)
}
