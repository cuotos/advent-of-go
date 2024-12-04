package utils

import (
	"bufio"
	"bytes"
	"os"
)

func GetFileAsBytes(filename string) []byte {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return fileContent
}

func ReadLines(fileContent []byte) []string {

	lines := []string{}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
