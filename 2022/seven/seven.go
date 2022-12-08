package seven

import (
	"bytes"
	"fmt"
)

type State struct {
	CurrentDir *Dir
	RootDir    *Dir
}

type Object interface {
	Size() int
}

type File struct {
	Name string
	size int
}

func (f *File) String() string {
	return fmt.Sprintf("file: %s, size: %d", f.Name, f.size)
}

func (f *File) Size() int {
	return f.size
}

type Dir struct {
	Name  string
	Dirs  []*Dir
	Files []*File
}

func (d *Dir) String() string {
	return fmt.Sprintf("dir: %s, size: %d", d.Name, d.Size())
}

func (d *Dir) Size() int {
	total := 0

	for _, f := range d.Files {
		total += f.size
	}

	for _, d := range d.Dirs {
		total += d.Size()
	}
	return total
}

func (d *Dir) LS() string {
	output := ""

	for _, d := range d.Dirs {
		output += fmt.Sprintln(d.String())
	}

	for _, f := range d.Files {
		output += fmt.Sprintln(f.String())
	}

	return output
}

func (d *Dir) Tree() string {
	output := printDir(d, "")
	return output
}

func printDir(d *Dir, indent string) string {
	output := fmt.Sprintf("%s- %s (dir, size=%d)\n", indent, d.Name, d.Size())

	for _, f := range d.Files {
		output += fmt.Sprintf("%s- %s (file, size=%d)\n", indent+"  ", f.Name, f.size)
	}
	for _, d := range d.Dirs {
		output += printDir(d, indent+"  ")
	}
	return output
}

func ParseLine(input []byte) {
	if bytes.HasPrefix(input, []byte("$ ")) {
		ParseInstruction(input)
	} else {
		ParseInfo(input)
	}
}

func ParseInstruction(input []byte) {}
func ParseInfo(input []byte)        {}

func Run(input []byte) int {
	return 0
}
