package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)


func main() {

	start := time.Now()

	lines := []string{}

	fileContent, err := ioutil.ReadFile("2-2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	for _, line := range lines {
		for _, compareLine := range lines {
			same, diff := compareLines(line, compareLine)
			if diff == 1 {
				fmt.Println(same)
				goto exit
			}
		}
	}

	exit:
	fmt.Println(time.Now().Sub(start))
}

func compareLines(s1, s2 string) (string, int) {

	differences := 0
	same := []byte{}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			differences++
		} else {
			same = append(same, s1[i])
		}
	}

	return string(same), differences
}
