package eight

import (
	"bufio"
	"bytes"
	"strconv"
)

type Forest [][]int

func parseForest(input []byte) Forest {
	f := Forest{}

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {

		row := []int{}

		line := scanner.Text()

		for _, tree := range line {
			treeInt, _ := strconv.Atoi(string(tree))
			row = append(row, treeInt)
		}
		f = append(f, row)
	}

	return f
}

func isTreeVisible(tree int, others [4][]int) bool {
	//if blocked == 4 then its not possible to see tree from any direction
	var blocked int

	for _, direction := range others {
		for _, other := range direction {
			if other >= tree {
				blocked++
				break
			}
		}
	}

	return blocked < 4
}

func Run(input []byte) int {
	f := parseForest(input)

	visibleTrees := 0

	for i, row := range f {
		for j, tree := range row {
			// if the tree is the top or bottom row, or on the ends then it will always be visible
			if i == 0 || i == len(f)-1 || j == 0 || j == len(row)-1 {
				visibleTrees++
				continue
			}

			otherTrees := getOtherTrees(f, i, j)

			if isTreeVisible(tree, otherTrees) {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func getOtherTrees(f Forest, treeRow int, treeColumn int) [4][]int {
	// 0 = above
	// 1 = below
	// 2 = left
	// 3 = right
	trees := [4][]int{}

	for rowIdx, row := range f {
		switch {
		case rowIdx < treeRow:
			trees[0] = append(trees[0], row[treeColumn])
		case rowIdx > treeRow:
			trees[1] = append(trees[1], row[treeColumn])
		case rowIdx == treeRow:
			trees[2] = row[:treeColumn]
			trees[3] = row[treeColumn+1:]
		}
	}

	return trees
}
