package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%dx%d", c.x, c.y)
}

type grid [][]coord

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func manhattanDist(a, b *coord) int {
	// for (p1, p2) and (q1, q2) the distance is (p1 - q1) + (p2 - q2)
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func parseCoord(input string) (*coord, error) {
	c := coord{}

	_, err := fmt.Sscanf(input, "%d, %d", &c.x, &c.y)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

// create a grid based on the largest size of X and Ys
func createGrid(stars []*coord) (int, int) {
	var maxX, maxY int

	for _, c := range stars {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	return maxX, maxY
}

func coordAtEdge(c *coord, w, h int) bool {
	return c.x == 0 || c.y == 0 || c.y == h || c.x == w
}

func findRemoteLocations(coords []*coord, minDist int) int {

	var totalMDistanceForGridToCoord = make(map[coord]int)

	w, h := createGrid(coords)

	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			curCoord := coord{x, y}
			for _, coord := range coords {
				totalMDistanceForGridToCoord[curCoord] += manhattanDist(&curCoord, coord)
			}
		}
	}

	total := 0

	for _, coordTotal := range totalMDistanceForGridToCoord {
		if coordTotal < minDist {
			total++
		}
	}

	return total
}

func main() {
	f, err := os.Open("2018/6-1/input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var coords []*coord
	s := bufio.NewScanner(f)
	for s.Scan() {
		c, err := parseCoord(s.Text())
		if err != nil {
			log.Fatal("cannot parse coord:", err)
		}
		coords = append(coords, c)
	}

	fmt.Println(findRemoteLocations(coords, 10000))

}

func printGrid(g grid) {
	for _, row := range g {
		fmt.Println(row)
	}
}
