package seven

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   []*Dir
	Files  []*File
}

func (d *Dir) Size() int {
	total := 0

	for _, f := range d.Files {
		total += f.Size
	}

	for _, d := range d.Dirs {
		total += d.Size()
	}

	return total
}

type File struct {
	Name string
	Size int

	Parent *Dir
}

func process(input []byte) *Dir {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	root := &Dir{
		Name: "/",
	}

	var curDir *Dir
	_ = curDir

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd ") {
			dirName := strings.TrimPrefix(line, "$ cd ")

			if dirName == ".." {
				curDir = curDir.Parent
				continue
			}
			if dirName == "/" {
				curDir = root
			} else {
				d := &Dir{
					Name:   dirName,
					Parent: curDir,
				}

				curDir.Dirs = append(curDir.Dirs, d)

				curDir = d
			}
		}

		if !strings.HasPrefix(line, "$") {
			// if the line starts with "dir" its a dir, else its a file
			if strings.HasPrefix(line, "dir") {
				// there is a dir
				continue
			} else {
				// this is a file
				split := strings.Split(line, " ")
				fileSize, _ := strconv.Atoi(split[0])

				f := &File{
					Name: split[1],
					Size: fileSize,
				}

				curDir.Files = append(curDir.Files, f)
			}

		}
	}

	return root
}

func (d *Dir) Tree(indent string) {
	fmt.Printf("%s- %s (dir, size: %d)\n", indent, d.Name, d.Size())
	for _, f := range d.Files {
		fmt.Printf("%s- %s (file, size: %d)\n", indent+"  ", f.Name, f.Size)
	}

	for _, d := range d.Dirs {
		d.Tree("  ")
	}
}

type Filter func(d *Dir) int

func ProcessTree(d *Dir, filters ...Filter) int {
	total := 0

	for _, d := range d.Dirs {
		if d.Size() <= 100000 {
			total += d.Size()
		}
		total += ProcessTree(d)
	}

	return total
}

func Run(input []byte) int {
	d := process(input)
	i := ProcessTree(d)

	return i
}
