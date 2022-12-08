package seven

import (
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

func TestGetSizeOfDir(t *testing.T) {
	f1 := &File{Name: "f1", size: 10}
	f2 := &File{Name: "f2", size: 20}

	f3 := &File{Name: "f3", size: 1}
	f4 := &File{Name: "f4", size: 2}

	f5 := &File{Name: "f5", size: 1}
	f6 := &File{Name: "f6", size: 2}

	d3 := &Dir{
		Name:  "dir3",
		Files: []*File{f5, f6},
	}

	d2 := &Dir{
		Name:  "dir2",
		Files: []*File{f3, f4},
		Dirs:  []*Dir{d3},
	}

	d1 := &Dir{
		Name:  "/",
		Files: []*File{f1, f2},
		Dirs:  []*Dir{d2},
	}

	expectedTree := `- / (dir, size=36)
  - f1 (file, size=10)
  - f2 (file, size=20)
  - dir2 (dir, size=6)
    - f3 (file, size=1)
    - f4 (file, size=2)
    - dir3 (dir, size=3)
      - f5 (file, size=1)
      - f6 (file, size=2)
`

	assert.Equal(t, 36, d1.Size())
	assert.Equal(t, expectedTree, d1.Tree())
}

func TestParseLine(t *testing.T) {

}

func TestSandbox(t *testing.T) {
	// rootDir := &Dir{}
	// s := State{
	// 	CurrentDir: rootDir,
	// 	RootDir:    rootDir,
	// }

}
