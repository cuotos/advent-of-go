package main

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

// A, X Rock
// B, Y Paper
// C, Z Scissors

// Score is shape YOU select
// 1 Rock
// 2 Paper
// 3 Scissors
// PLUS
// 0 Loose
// 3 Draw
// 6 Win

// Them Us
// Rock Rock         4
// Rock Paper        8
// Rock Scissors     3
// Paper Rock        1
// Paper Paper       5
// Paper Scissors    9
// Scissors Rock     7
// Scissors Paper    2
// Scissors Scissors 6

// X = lose
// Y = draw
// Z = win

type Shape int

func (s Shape) String() string {
	switch s {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}

	return ""
}

const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

func day2Pt1(input []byte) int {
	game := parseGame(input)

	totalScore := 0
	numberOfRounds := 0

	for _, round := range game {
		roundScore := PlayRound(round.Them, round.Us)
		totalScore += roundScore
		numberOfRounds++
	}

	log.Printf("played %d rounds", numberOfRounds)
	return totalScore
}

func day2Pt2(input []byte) int {
	return 0
}

func PlayRound(them, us Shape) (ourScore int) {

	switch {
	case them == Rock:
		switch us {
		case Rock:
			return 4
		case Paper:
			return 8
		case Scissors:
			return 3
		}
	case them == Paper:
		switch us {
		case Rock:
			return 1
		case Paper:
			return 5
		case Scissors:
			return 9
		}
	case them == Scissors:
		switch us {
		case Rock:
			return 7
		case Paper:
			return 2
		case Scissors:
			return 6
		}
	}

	return 0
}

type Round struct {
	Them Shape
	Us   Shape
}

func parseGame(input []byte) []Round {
	var game = []Round{}

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		splitLine := strings.Split(line, " ")

		round := Round{
			Them: stringToShape(splitLine[0]),
			Us:   stringToShape(splitLine[1]),
		}
		game = append(game, round)
	}

	return game
}

func stringToShape(input string) Shape {
	switch input {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	return Rock
}
