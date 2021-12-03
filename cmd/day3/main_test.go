package main

import (
	"testing"

	"github.com/lhgmnn/advent-of-code-2021/input"
	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	in := &input.Input{
		Values: []string{
			"00100",
			"11110",
			"10110",
			"10111",
		},
	}

	counters := make([]Counter, len(in.Values[0]))
	CountCharacters(in, counters)

	assert.Equal(t, counters[0].Zeroes, 1)
	assert.Equal(t, counters[0].Ones, 3)

	assert.Equal(t, counters[1].Zeroes, 3)
	assert.Equal(t, counters[1].Ones, 1)

	assert.Equal(t, counters[2].Zeroes, 0)
	assert.Equal(t, counters[2].Ones, 4)

	assert.Equal(t, counters[3].Zeroes, 1)
	assert.Equal(t, counters[3].Ones, 3)

	assert.Equal(t, counters[4].Zeroes, 3)
	assert.Equal(t, counters[4].Ones, 1)
}

func TestExtractGamma(t *testing.T) {
	counters := []Counter{
		{Ones: 7, Zeroes: 5},
		{Ones: 5, Zeroes: 7},
		{Ones: 8, Zeroes: 4},
		{Ones: 7, Zeroes: 5},
		{Ones: 5, Zeroes: 7},
	}

	assert.Equal(t, int64(22), ExtractGamma(counters))
}

func TestExtractEpsilon(t *testing.T) {
	counters := []Counter{
		{Ones: 7, Zeroes: 5},
		{Ones: 5, Zeroes: 7},
		{Ones: 8, Zeroes: 4},
		{Ones: 7, Zeroes: 5},
		{Ones: 5, Zeroes: 7},
	}

	assert.Equal(t, int64(9), ExtractEpsilon(counters))
}
