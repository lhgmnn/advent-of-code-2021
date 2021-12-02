package main

import (
	"fmt"
	"strconv"

	"github.com/lhgmnn/advent-of-code-2021/input"
)

func CountMeasurementIncreases(in *input.Input) int {
	counter := 0

	var current int64
	for v := 0; v < len(in.Values)-1; v++ {
		current, _ = strconv.ParseInt(in.Values[v], 10, 64)
		next, _ := strconv.ParseInt(in.Values[v+1], 10, 64)

		if next >= current {
			counter += 1
		}
	}

	return counter
}

func main() {
	in, err := input.NewFromFile("cmd/day1/input.txt")
	if err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}

	val := CountMeasurementIncreases(in)

	fmt.Printf("Count: %d\n", val)
}
