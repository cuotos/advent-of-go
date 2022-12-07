package main

import (
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

func TestStringToShape(t *testing.T) {

	tcs := []struct {
		Input    string
		Expected Shape
	}{
		{"A", Rock},
		{"B", Paper},
		{"C", Scissors},
		{"X", Rock},
		{"Y", Paper},
		{"Z", Scissors},
	}

	for _, tc := range tcs {
		actual := stringToShape(tc.Input)
		assert.Equalf(t, tc.Expected, actual, "expected %s but got %s", tc.Expected, actual)
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
		{Rock, Rock},
		{Rock, Paper},
		{Rock, Scissors},
		{Paper, Rock},
		{Paper, Paper},
		{Paper, Scissors},
		{Scissors, Rock},
		{Scissors, Paper},
		{Scissors, Scissors},
	}

	actual := parseGame(input)

	assert.ElementsMatch(t, expected, actual)
}

func TestRoundScores(t *testing.T) {
	tcs := []struct {
		Round
		roundScore int
		subTotal   int
	}{
		{Round{Rock, Rock}, 4, 4},
		{Round{Paper, Paper}, 5, 9},
		{Round{Paper, Paper}, 5, 14},
		{Round{Scissors, Rock}, 7, 21},
		{Round{Paper, Rock}, 1, 22},
		{Round{Scissors, Scissors}, 6, 28},
		{Round{Scissors, Scissors}, 6, 34},
		{Round{Rock, Scissors}, 3, 37},
		{Round{Rock, Scissors}, 3, 40},
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

	game := parseGame(input)

	total := 0
	for _, round := range game {
		roundScore := PlayRound(round.Them, round.Us)
		total += roundScore
	}

	assert.Equal(t, 40, total)
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
