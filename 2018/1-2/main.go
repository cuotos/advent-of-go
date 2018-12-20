package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	var deltas = []int{}

	fileContent, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		deltas = append(deltas, i)
	}

	var seen = make(map[int]struct{})
	var runningTotal int

	for {
		for _, i := range deltas {
			runningTotal += i

			if _, ok := seen[runningTotal]; ok {
				fmt.Println(runningTotal)
				fmt.Println(time.Now().Sub(start))
				os.Exit(0)
			} else {
				seen[runningTotal] = struct{}{}
			}
		}
	}
}
