package four

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

type rng [2]int

func rangeContains(a, b rng) bool {
	return a[0] <= b[0] && a[1] >= b[1] || b[0] <= a[0] && b[1] >= a[1]
}

func rangesOverlap(a, b rng) bool {
	sections := map[int]bool{}

	for i := a[0]; i <= a[1]; i++ {
		sections[i] = true
	}

	for j := b[0]; j <= b[1]; j++ {
		if _, found := sections[j]; found {
			return true
		}
	}

	return false
}

type Pair struct {
	rngA rng
	rngB rng
}

func Run(input []byte) int {
	pairs := parse(input)

	overlapping := 0
	for _, p := range pairs {
		if rangesOverlap(p.rngA, p.rngB) {
			overlapping++
		}
	}

	return overlapping
}

func parse(input []byte) []Pair {

	scanner := bufio.NewScanner(bytes.NewReader(input))

	pairs := []Pair{}

	for scanner.Scan() {
		pair := Pair{}

		line := strings.TrimSpace(scanner.Text())

		splitLine := strings.Split(line, ",")

		pair.rngA = parseRange(splitLine[0])
		pair.rngB = parseRange(splitLine[1])

		pairs = append(pairs, pair)

	}

	return pairs
}

func parseRange(input string) rng {
	split := strings.Split(string(input), "-")

	lower, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatalf("failed to parse: %s", err)
	}

	upper, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatalf("failed to parse: %s", err)
	}

	return rng{lower, upper}
}
