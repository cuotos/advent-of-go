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

func maxFinite(stars []*coord) int {
	w, h := createGrid(stars)

	var numberOfCellsPerStar = make(map[*coord]int)
	var infStar = make(map[*coord]struct{})

	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			var (
				gc  = &coord{x, y}
				min = -1
				cc  *coord
			)

			for _, c := range stars {
				if dist := manhattanDist(gc, c); dist < min || min == -1 {
					min = dist
					cc = c
				} else if dist == min {
					cc = &coord{-1,-1}
				}
			}

			// if this coord is on an edge, its star will be an "infinite" one and should be noted for later
			if coordAtEdge(gc, w, h) {
				infStar[cc] = struct{}{}
			}

			numberOfCellsPerStar[cc]++
		}
	}

	max := 0

	for c, l := range numberOfCellsPerStar {
		if _, found := infStar[c]; l > max && !found {
			max = l
		}
	}

	return max
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

	fmt.Println("mine:", maxFinite(coords))
}

func printGrid(g grid) {
	for _, row := range g {
		fmt.Println(row)
	}
}
