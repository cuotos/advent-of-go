package main

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
)

func getInventory(input []byte) []int {

	inventory := []int{}

	scanner := bufio.NewScanner(bytes.NewReader(input))

	currentTotal := 0
	for scanner.Scan() {

		lineText := scanner.Text()
		if lineText == "" {
			inventory = append(inventory, currentTotal)
			currentTotal = 0
			continue
		}

		caloriesInt, err := strconv.Atoi(lineText)
		if err != nil {
			panic(err)
		}
		currentTotal += caloriesInt
	}

	// add the last one to the slice
	inventory = append(inventory, currentTotal)

	return inventory
}

func day1Pt1(input []byte) int {
	inv := getInventory(input)

	sort.IntSlice(inv).Sort()

	return inv[len(inv)-1]
}

func day1Pt2(input []byte) int {

	inv := getInventory(input)
	sort.IntSlice(inv).Sort()

	l := len(inv)

	return inv[l-1] + inv[l-2] + inv[l-3]
}
