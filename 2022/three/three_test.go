package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
)

func TestParseLine(t *testing.T) {
	tcs := []struct {
		input    []byte
		expected Rucksack
	}{
		{
			[]byte("vJrwpWtwJgWrhcsFMMfFFhFp"),
			Rucksack{
				Comp1:    []byte("vJrwpWtwJgWr"),
				Comp2:    []byte("hcsFMMfFFhFp"),
				overlap:  'p',
				priority: 16,
			},
		},
		{
			[]byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
			Rucksack{
				Comp1:    []byte("jqHRNqRjqzjGDLGL"),
				Comp2:    []byte("rsFMfFZSrLrFZsSL"),
				overlap:  'L',
				priority: 38,
			},
		},
		{
			[]byte("PmmdzqPrVvPwwTWBwg"),
			Rucksack{
				Comp1:    []byte("PmmdzqPrV"),
				Comp2:    []byte("vPwwTWBwg"),
				overlap:  'P',
				priority: 42,
			},
		},
		{
			[]byte("baac"),
			Rucksack{
				Comp1:    []byte("ba"),
				Comp2:    []byte("ac"),
				priority: 1,
				overlap:  'a',
			},
		},
		{
			[]byte("zz"),
			Rucksack{
				Comp1:    []byte("z"),
				Comp2:    []byte("z"),
				priority: 26,
				overlap:  'z',
			},
		},
		{
			[]byte("AA"),
			Rucksack{
				Comp1:    []byte("A"),
				Comp2:    []byte("A"),
				priority: 27,
				overlap:  'A',
			},
		},
		{
			[]byte("ZZ"),
			Rucksack{
				Comp1:    []byte("Z"),
				Comp2:    []byte("Z"),
				priority: 52,
				overlap:  'Z',
			},
		},
	}

	for _, tc := range tcs {
		actual := parseLine(tc.input)
		require.Equal(t, len(actual.Comp1), len(actual.Comp2))
		assert.ElementsMatch(t, tc.expected.Comp1, actual.Comp1)
		assert.ElementsMatch(t, tc.expected.Comp2, actual.Comp2)
		assert.Equal(t, tc.expected.overlap, actual.overlap)
		assert.Equal(t, tc.expected.priority, actual.priority)
	}
}

func TestGetTotal(t *testing.T) {
	input := []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`)

	assert.Equal(t, 157, getTotalPart1(input))
}

func TestThreeWayContains(t *testing.T) {
	tcs := []struct {
		inputA, inputB, inputC string
		common                 byte
	}{
		{
			"ab", "ac", "ad",
			'a',
		},
		{
			"yyya", "zzza", "xaxx",
			'a',
		},
	}

	for _, tc := range tcs {
		actual := threeWayContains([]byte(tc.inputA), []byte(tc.inputB), []byte(tc.inputC))
		assert.Equal(t, tc.common, actual)
	}
}

func TestLookupAlphabet(t *testing.T) {
	assert.Equal(t, 1, slices.Index(lookup, 'a'))
	assert.Equal(t, 26, slices.Index(lookup, 'z'))
	assert.Equal(t, 27, slices.Index(lookup, 'A'))
	assert.Equal(t, 52, slices.Index(lookup, 'Z'))
}

func TestPart2(t *testing.T) {
	input := []byte(`ab
ac
ad
za
zb
zc
Za
Zb
Zc`)

	total := getTotalPart2(input)
	assert.Equal(t, 1+26+52, total)
}
