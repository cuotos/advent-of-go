package seven

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// - / (dir)
//   - a (dir)
//     - e (dir)
//       - i (file, size=584)
//     - f (file, size=29116)
//     - g (file, size=2557)
//     - h.lst (file, size=62596)
//   - b.txt (file, size=14848514)
//   - c.dat (file, size=8504156)
//   - d (dir)
//     - j (file, size=4060174)
//     - d.log (file, size=8033020)
//     - d.ext (file, size=5626152)
//     - k (file, size=7214296)

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDemo(t *testing.T) {
	root := process([]byte(testInput))
	total := ProcessTree(root)

	assert.Equal(t, 95437, total)
}

func TestProcess(t *testing.T) {

	input := []byte(`$ cd /
$ ls
dir a
1 b.txt
$ cd a
$ ls
100 c.txt`)
	_ = input

	root := process([]byte(testInput))
	total := ProcessTree(root)
	fmt.Println("total", total)
}

func TestDemoPt2(t *testing.T) {
	i := Run([]byte(testInput))
	fmt.Println(i)
}

func PrintEach(s []*Dir) {
	for _, i := range s {
		fmt.Println(i.Name, i.Size())
	}
}
