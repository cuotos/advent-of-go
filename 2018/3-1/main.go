package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

type coord struct {
	x, y int
}

type claim struct {
	id int
	x, y int
	w, h int
}

type fabric struct {
	claims map[coord]int
}

func (f *fabric) claim(c claim){

	for x:=0; x<c.w; x++{
		for y:=0; y<c.h; y++{
			f.claims[coord{c.x+x, c.y+y}]++
		}
	}
}

func newFabric() *fabric {
	return &fabric{
		make(map[coord]int),
	}
}

func parseLine(input string) claim {
	var c claim

	_, err := fmt.Sscanf(input, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
	if err != nil {
		panic(err)
	}

	return c
}

func (f fabric) overlaps() int {
	overlaps := 0
	for _, c := range f.claims {
		if c > 1 {
			overlaps++
		}
	}

	return overlaps
}

func main() {

	f := newFabric()

	content, err := ioutil.ReadFile("2018/3-1/input")
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(content)

	s := bufio.NewScanner(r)

	for s.Scan() {
		c := parseLine(s.Text())
		f.claim(c)
	}

	fmt.Println(f.overlaps())
}
