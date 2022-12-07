package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	tcs := []struct {
		input  string
		expect move
	}{
		{"move 1 from 7 to 6", move{1, 7, 6}},
		{"move 10 from 17 to 6", move{10, 17, 6}},
	}

	for _, tc := range tcs {
		actual := parseLine([]byte(tc.input))

		assert.Equal(t, tc.expect.count, actual.count)
		assert.Equal(t, tc.expect.src, actual.src)
		assert.Equal(t, tc.expect.dest, actual.dest)
	}
}
