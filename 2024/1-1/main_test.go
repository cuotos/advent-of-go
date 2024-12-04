package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	expected := 11

	actual := calculateOffset(list1, list2)

	assert.Equal(t, expected, actual)
}

func TestXxx(t *testing.T) {
	main()
}
