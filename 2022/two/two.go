package two

import (
	"bufio"
	"bytes"
	"strings"
)

var Reference = map[Shape]map[Result]Shape{
	Rock: {
		Win:  Paper,
		Lose: Scissors,
	},
	Paper: {
		Win:  Scissors,
		Lose: Rock,
	},
	Scissors: {
		Win:  Rock,
		Lose: Paper,
	},
}

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

type Result int

func (s Result) String() string {
	switch s {
	case Lose:
		return "Lose"
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	}

	return ""
}

const (
	Lose Result = iota
	Draw
	Win
)

func Day2(input []byte) int {
	game := parseGame(input)

	totalScore := 0

	for _, round := range game {
		requiredShape := calculateRequiredShape(round)
		roundScore := PlayRound(round.Them, requiredShape)
		totalScore += roundScore
	}

	return totalScore
}

func calculateRequiredShape(round Round) Shape {
	if round.Result == Draw {
		return round.Them
	}

	return Reference[round.Them][round.Result]
}

func PlayRound(them, us Shape) int {

	if them == us {
		return 3 + int(us)
	}

	if Reference[them][Lose] == us {
		return 0 + int(us)
	}

	if Reference[them][Win] == us {
		return 6 + int(us)
	}

	return 0
}

type Round struct {
	Them   Shape
	Result Result
}

func parseGame(input []byte) []Round {
	var game = []Round{}

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		splitLine := strings.Split(line, " ")

		round := parseRound(splitLine)
		game = append(game, round)
	}

	return game
}

func parseRound(input []string) Round {
	round := Round{}

	switch input[0] {
	case "A":
		round.Them = Rock
	case "B":
		round.Them = Paper
	case "C":
		round.Them = Scissors
	}

	switch input[1] {
	case "X":
		round.Result = Lose
	case "Y":
		round.Result = Draw
	case "Z":
		round.Result = Win
	}

	return round
}
