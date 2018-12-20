package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {

	tcs := []struct{
		coord1 coord
		coord2 coord
		expected int
	}{
		{coord{1,1},coord{1,6},5},
		{coord{1,1},coord{8,3},9},
	}

	for _, tc := range tcs {
		actual := findManhattenDistance(tc.coord1, tc.coord2)
		assert.Equal(t, tc.expected, actual)
	}
}
