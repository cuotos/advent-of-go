package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestThreeOne(t *testing.T) {
	fmt.Println(parseLine(`00100`))
}

func parseLine(line string) int64 {
	i, _ := strconv.ParseInt(line, 2, 64)
	return i
}

var sampleInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
