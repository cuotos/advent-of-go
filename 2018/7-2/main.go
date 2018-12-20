package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

type step struct {
	stepId    string
	blockedBy string
}

func getSecondsPerLetter(input rune) int {
	return int(unicode.ToUpper(input)) - 4
}

func doit(input io.Reader) string {

	var (
		allSteps  = make(map[string]struct{})
		stepOrder = []string{}
	)

	var steps = make(map[string][]string)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var (
			stepString string
			br         string
		)

		_, err := fmt.Sscanf(scanner.Text(), "Step %v must be finished before step %v can begin.", &br, &stepString)
		if err != nil {
			panic(err)
		}

		allSteps[br] = struct{}{}
		allSteps[stepString] = struct{}{}

		steps[stepString] = append(steps[stepString], br)
	}

	for {
		readyToRun := []string{}

		for step := range allSteps {
			if len(steps[step]) == 0 {
				readyToRun = append(readyToRun, step)
			}
		}

		sort.Strings(readyToRun)

		if len(readyToRun) > 0 {
			stepOrder = append(stepOrder, readyToRun[0])
		} else {
			break
		}

		delete(allSteps, readyToRun[0])

		for step, blockers := range steps {
			for i, b := range blockers {
				if b == readyToRun[0] {
					steps[step] = append(steps[step][:i], steps[step][i+1:]...)
				}
			}
		}

	}

	return strings.Join(stepOrder, "")
}

func main() {

	f, err := os.Open("2018/7-1/exampleinput")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileContent, _ := ioutil.ReadAll(f)
	fmt.Println(doit(bytes.NewReader(fileContent)))
}
