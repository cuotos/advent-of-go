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
	claims map[coord][]int
	ids map[int]struct{}
}

func (f *fabric) claim(c claim){

	for x:=0; x<c.w; x++{
		for y:=0; y<c.h; y++{

			xy := coord{c.x + x, c.y + y}

			f.ids[c.id] = struct{}{}

			f.claims[xy] = append(f.claims[xy], c.id)
		}
	}
}

func newFabric() *fabric {
	return &fabric{
		make(map[coord][]int),
		make(map[int]struct{}),
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

//func (f fabric) collectOverlapIds() {
//
//	for _, c := range f.claims {
//		if len(c) > 1 {
//		}
//	}
//}

func main() {

	f := newFabric()

	var allClaims []claim

	content, err := ioutil.ReadFile("2018/3-2/input")
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(content)

	s := bufio.NewScanner(r)

	for s.Scan() {
		c := parseLine(s.Text())
		f.claim(c)
		allClaims = append(allClaims, c)
	}

	seen := make(map[int]struct{})

	for _, cl := range f.claims {
		if len(cl) > 1 {
			for _, clid := range cl {
				seen[clid] = struct{}{}
			}
		}
	}

	for id, _ := range f.ids {
		if _, ok := seen[id]; !ok {
			fmt.Println(id)
		}
	}

}
