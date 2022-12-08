package six

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMarker(t *testing.T) {
	tcs := []struct {
		input          string
		expectedMarker int
		sizeOfMarker   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 4},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 4},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 4},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 4},
	}

	for _, tc := range tcs {
		actual := findMarker([]byte(tc.input), tc.sizeOfMarker)
		assert.Equal(t, tc.expectedMarker, actual)
	}
}

func TestFourWayCompare(t *testing.T) {
	tcs := []struct {
		Input  []rune
		Expect bool
	}{
		{[]rune{'A', 'B', 'C', 'D'}, true},
		{[]rune{'A', 'A', 'C', 'D'}, false},
		{[]rune{'A', 'B', 'C', 'C'}, false},
		{[]rune{'A', 'B', 'B', 'C'}, false},
		{[]rune{'A', 'B', 'C', 'A'}, false},
	}

	for _, tc := range tcs {
		allDifferent := allElementsDifferent(tc.Input)
		assert.Equal(t, tc.Expect, allDifferent)
	}
}

func TestDetectAllDifferent(t *testing.T) {
	tcs := []struct {
		Input        []rune
		AllDifferent bool
	}{
		{[]rune{'A', 'B'}, true},
		{[]rune{'A', 'A'}, false},
		{[]rune{'A', 'B', 'C', 'D'}, true},
		{[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'B'}, false},
	}

	for _, tc := range tcs {
		actual := allElementsDifferent(tc.Input)
		assert.Equal(t, tc.AllDifferent, actual)
	}

}
