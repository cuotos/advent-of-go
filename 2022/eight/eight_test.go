package eight

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var demoInput = []byte(`30373
25512
65332
33549
35390`)

func TestParseForest(t *testing.T) {
	f := parseForest(demoInput)
	fmt.Println(f)
}

func TestIsTreeVisible(t *testing.T) {
	tcs := []struct {
		tree    int
		others  [4][]int
		visible bool
	}{
		{5, [4][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}, true},
		{5, [4][]int{{1, 6}, {1, 2}, {1, 2}, {1, 2}}, true},
		{1, [4][]int{{2, 2}, {2, 2}, {2, 2}, {2, 2}}, false},
	}

	for _, tc := range tcs {
		actual := isTreeVisible(tc.tree, tc.others)
		assert.Equal(t, tc.visible, actual)
	}
}

func TestDemo(t *testing.T) {
	actual := Run(demoInput)
	assert.Equal(t, 21, actual)
}
