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

func TestManipulateState(t *testing.T) {
	s := InitialState()

	assert.Equal(t, 8, len(s.Rows[0]))
	assert.Equal(t, 3, len(s.Rows[4-1]))
	assert.Equal(t, []rune{'N', 'B', 'S'}, s.Rows[3])
	assert.Equal(t, []rune{'R', 'N', 'F', 'V', 'L', 'J', 'S', 'M'}, s.Rows[0])

	move := move{
		count: 2,
		src:   1,
		dest:  4,
	}

	s.move(move)

	assert.Equal(t, []rune{'R', 'N', 'F', 'V', 'L', 'J'}, s.Rows[1-1])
	assert.Equal(t, []rune{'N', 'B', 'S', 'M', 'S'}, s.Rows[4-1])
}

func TestManipulateState2(t *testing.T) {
	s := InitialState()

	move1 := move{
		count: 3,
		src:   4,
		dest:  9,
	}
	s.move(move1)

	assert.Equal(t, []rune{}, s.Rows[4-1])
	assert.Equal(t, []rune{'L', 'M', 'H', 'Z', 'N', 'F', 'S', 'B', 'N'}, s.Rows[9-1])
}

func TestPrintState(t *testing.T) {
	s := InitialState()
	assert.Equal(t, "MHGSNWWVF", s.printState())
}

// func TestRun(t *testing.T) {
// 	actual := Run([]byte(`move 1 from 1 to 5
// move 1 from 2 to 3`))
// 	assert.Equal(t, "SWHSMWWVF", actual)
// }

func TestDemoData(t *testing.T) {
	s := &State{
		Rows: [9][]rune{
			{'Z', 'N'},
			{'M', 'C', 'D'},
			{'P'},
			{}, {}, {}, {}, {}, {},
		},
	}
	actual := Run([]byte(`move 1 from 2 to 1
	move 3 from 1 to 3
	move 2 from 2 to 1
	move 1 from 1 to 2`),
		s)

	assert.Equal(t, "CMZ", actual)
}
