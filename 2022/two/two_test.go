package two

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	tcs := []struct {
		Them          Shape
		Us            Shape
		ExpectedScore int
	}{
		{Rock, Rock, 4},
		{Rock, Paper, 8},
		{Rock, Scissors, 3},
		{Paper, Rock, 1},
		{Paper, Paper, 5},
		{Paper, Scissors, 9},
		{Scissors, Rock, 7},
		{Scissors, Paper, 2},
		{Scissors, Scissors, 6},
	}

	for _, tc := range tcs {
		actual := PlayRound(tc.Them, tc.Us)
		assert.Equal(t, tc.ExpectedScore, actual)
	}
}

func TestRoundScores(t *testing.T) {
	tcs := []struct {
		Them       Shape
		Us         Shape
		roundScore int
		subTotal   int
	}{
		{Rock, Rock, 4, 4},
		{Paper, Paper, 5, 9},
		{Paper, Paper, 5, 14},
		{Scissors, Rock, 7, 21},
		{Paper, Rock, 1, 22},
		{Scissors, Scissors, 6, 28},
		{Scissors, Scissors, 6, 34},
		{Rock, Scissors, 3, 37},
		{Rock, Scissors, 3, 40},
	}

	totalScore := 0

	for _, r := range tcs {
		actualRoundScore := PlayRound(r.Them, r.Us)
		totalScore += actualRoundScore
		assert.Equal(t, r.roundScore, actualRoundScore)
		assert.Equal(t, r.subTotal, totalScore)
	}
}

func TestStrategyAndParse(t *testing.T) {
	input := []byte(`A X
	B Y
	B Y
	C X
	B X
	C Z
	C Z
	A Z
	A Z`)

	requiredShape := []Shape{
		Scissors,
		Paper,
		Paper,
		Paper,
		Rock,
		Rock,
		Rock,
		Paper,
		Paper,
	}

	game := parseGame(input)

	total := 0
	for i, round := range game {
		roundScore := PlayRound(round.Them, requiredShape[i])
		total += roundScore
	}

	assert.Equal(t, 46, total)
}

func TestShapeString(t *testing.T) {
	tcs := []struct {
		input  Shape
		expect string
	}{
		{Rock, "Rock"},
		{Paper, "Paper"},
		{Scissors, "Scissors"},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.expect, tc.input.String())
	}
}

func TestParseGame(t *testing.T) {
	input := []byte(`A X
A Y
A Z
B X
B Y
B Z
C X 
C Y
C Z`)

	expected := []Round{
		{Rock, Lose},
		{Rock, Draw},
		{Rock, Win},
		{Paper, Lose},
		{Paper, Draw},
		{Paper, Win},
		{Scissors, Lose},
		{Scissors, Draw},
		{Scissors, Win},
	}

	actual := parseGame(input)

	assert.ElementsMatch(t, expected, actual)
}

func TestGetRequiredShape(t *testing.T) {

	tcs := []struct {
		Them          Shape
		Result        Result
		ExpectedShape Shape
	}{
		{Paper, Lose, Rock},
		{Paper, Draw, Paper},
		{Paper, Win, Scissors},
		{Rock, Lose, Scissors},
		{Rock, Draw, Rock},
		{Rock, Win, Paper},
		{Scissors, Lose, Paper},
		{Scissors, Draw, Scissors},
		{Scissors, Win, Rock},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s-%s-%s", tc.Them, tc.Result, tc.ExpectedShape), func(t *testing.T) {
			r := Round{tc.Them, tc.Result}
			actual := calculateRequiredShape(r)
			assert.Equalf(t, tc.ExpectedShape, actual, "expected %s but got %s", tc.ExpectedShape, actual)
		})
	}
}
