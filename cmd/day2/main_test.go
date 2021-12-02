package main

import (
	"testing"

	"github.com/lhgmnn/advent-of-code-2021/input"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDepthAndPosition(t *testing.T) {
	in := &input.Input{
		Values: []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		},
	}

	d, p := CalculateDepthAndPosition(in)

	assert.Equal(t, 10, d)
	assert.Equal(t, 15, p)
}

func TestParseCommand(t *testing.T) {
	c, v := ParseCommand("forward 5")
	assert.Equal(t, "forward", c)
	assert.Equal(t, 5, v)
}
