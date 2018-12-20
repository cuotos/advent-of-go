package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	var total int

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

		total += i
	}

	fmt.Println(total)
}
