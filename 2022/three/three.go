package three

import (
	"bufio"
	"bytes"

	"golang.org/x/exp/slices"
)

var lookup = []byte(".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Rucksack struct {
	Input    []byte
	Comp1    []byte
	Comp2    []byte
	overlap  byte
	priority int
}

func Run(input []byte) int {
	return getTotalPart1(input)
}

func parseLine(input []byte) Rucksack {
	r := Rucksack{}

	r.Input = input
	r.Comp1 = r.Input[:len(r.Input)/2]
	r.Comp2 = r.Input[len(r.Input)/2:]

	for _, c := range r.Comp1 {
		if slices.Contains(r.Comp2, c) {
			r.overlap = c
			break
		}
	}

	r.priority = slices.Index(lookup, r.overlap)

	return r
}

func getTotalPart1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	total := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		r := parseLine(line)

		total += r.priority
	}
	return total
}

func getTotalPart2(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	total := 0
	for {
		if !scanner.Scan() {
			break
		}

		inputA := scanner.Bytes()
		scanner.Scan()
		inputB := scanner.Bytes()
		scanner.Scan()
		inputC := scanner.Bytes()

		overlap := threeWayContains(inputA, inputB, inputC)

		total += slices.Index(lookup, overlap)
	}

	return total
}

func threeWayContains(a, b, c []byte) byte {
	for _, i := range a {
		if slices.Contains(b, i) && slices.Contains(c, i) {
			return i
		}
	}
	return '.'
}
