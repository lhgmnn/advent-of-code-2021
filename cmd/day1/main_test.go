package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhgmnn/advent-of-code-2021/input"
)

func TestCountMeasurementIncreasesExample(t *testing.T) {
	in := &input.Input{
		Values: []string{
			"199",
			"200",
			"208",
			"210",
			"200",
			"207",
			"240",
			"269",
			"260",
			"263",
		},
	}

	assert.Equal(t, 7, CountMeasurementIncreases(in))
}

func TestCountMeasurementIncreasesCornerCases(t *testing.T) {
	assert.Equal(t, 0, CountMeasurementIncreases(&input.Input{
		Values: []string{
			"199",
		},
	}))

	assert.Equal(t, 1, CountMeasurementIncreases(&input.Input{
		Values: []string{
			"199",
			"200",
		},
	}))
}
