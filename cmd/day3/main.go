package main

import (
	"fmt"
	"strconv"

	"github.com/lhgmnn/advent-of-code-2021/input"
)

type Counter struct {
	Zeroes int
	Ones   int
}

// O(nxm) - probably the dumbest way, just iterate and count.
func CountCharacters(in *input.Input, c []Counter) {
	for _, value := range in.Values {
		for pos, char := range value {

			switch char {
			case '0':
				c[pos].Zeroes += 1
			case '1':
				c[pos].Ones += 1
			}
		}
	}
}

func ExtractGamma(c []Counter) int64 {
	b := make([]byte, len(c))

	for pos, v := range c {
		if v.Zeroes > v.Ones {
			b[pos] = '0'
		} else {
			b[pos] = '1'
		}
	}

	v, _ := strconv.ParseInt(string(b), 2, 64)
	return v
}

func ExtractEpsilon(c []Counter) int64 {
	b := make([]byte, len(c))

	for pos, v := range c {
		if v.Zeroes < v.Ones {
			b[pos] = '0'
		} else {
			b[pos] = '1'
		}
	}

	v, _ := strconv.ParseInt(string(b), 2, 64)
	return v
}

func main() {
	in, err := input.NewFromFile("cmd/day3/input.txt")
	if err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}

	counters := make([]Counter, len(in.Values[0]))
	CountCharacters(in, counters)

	epsilon := ExtractEpsilon(counters)
	gamma := ExtractGamma(counters)

	powerConsumption := epsilon * gamma
	fmt.Printf("Power consumption: %d\n", powerConsumption)

}
