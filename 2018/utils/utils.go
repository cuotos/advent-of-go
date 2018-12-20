package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
)

func ReadFileLines(filename string) []string {
	fileContent, err := ioutil.ReadFile(filename)

	lines := []string{}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
