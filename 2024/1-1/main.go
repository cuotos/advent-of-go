package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/cuotos/advent-of-code/utils"
)

//go:embed input.txt
var input []byte

func calculateOffset(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		i1 := l1[i]
		i2 := l2[i]

		if i1 > i2 {
			sum += i1 - i2
		} else {
			sum += i2 - i1
		}
	}
	return sum
}

func main() {
	lines := utils.ReadLines(input)

	l1 := []int{}
	l2 := []int{}
	for _, l := range lines {
		s1 := strings.Fields(l)[0]
		s2 := strings.Fields(l)[1]

		i1, _ := strconv.Atoi(s1)
		i2, _ := strconv.Atoi(s2)

		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	fmt.Println(calculateOffset(l1, l2))
}
