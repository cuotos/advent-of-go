package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	tcs := []struct {
		input    string
		expected string
	}{
		{
			`Step A must be finished before step B can begin.
Step B must be finished before step C can begin.`,
			`ABC`},
		{`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`,
			`CABDFE`},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.expected, doit(strings.NewReader(tc.input)))
	}
}
