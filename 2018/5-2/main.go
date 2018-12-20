package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
)

func main() {

	start := time.Now()

	bContent, err := ioutil.ReadFile("2018/5-1/input")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(bContent), "\n")

	var answers = make([]int, 26)

	wg := sync.WaitGroup{}

	index := 0

	for i:='a'; i<='z';i++{
		wg.Add(1)

		go func(i int, r rune, answers []int) {
			inputToProcess := remove(input, rune(r))
			answers[i] = process(inputToProcess)
			wg.Done()
		}(index, i, answers)

		index++
	}

	wg.Wait()

	sort.Slice(answers, func(i, j int) bool {
		return answers[i] < answers[j]
	})

	fmt.Println(answers[0])
	fmt.Println(time.Now().Sub(start))
}

func remove(input string, l rune) string {
	u := unicode.ToUpper(l)

	input = strings.Replace(input, string(l), "", -1)
	input = strings.Replace(input, string(u), "", -1)

	return input
}

func process(input string) int {
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
