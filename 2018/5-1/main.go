package main

import (
	"io/ioutil"
	"strings"
	"unicode"
)

func doit() int {

	bContent, err := ioutil.ReadFile("2018/5-1/input")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(bContent), "\n")

	i := 0
	for {
		if i+1 >= len(input) {
			break
		}

		current := rune(input[i])
		var invert rune

		if unicode.IsLower(current) {
			invert = unicode.ToUpper(current)
		} else {
			invert = unicode.ToLower(current)
		}

		if rune(input[i+1]) != invert {
			i++
			continue
		}

		input = input[:i] + input[i+2:]

		i = 0
	}

	return len(input)
}

func main() {
	doit()
}