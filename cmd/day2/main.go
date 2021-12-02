package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lhgmnn/advent-of-code-2021/input"
)

func ParseCommand(line string) (string, int) {
	command := strings.Split(line, " ")
	value, _ := strconv.Atoi(command[1])

	return command[0], value
}

func CalculateDepthAndPosition(in *input.Input) (int, int) {
	position := 0
	depth := 0

	for _, v := range in.Values {
		command, value := ParseCommand(v)
		switch command {
		case "forward":
			position += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	return depth, position
}

func main() {
	in, err := input.NewFromFile("cmd/day2/input.txt")
	if err != nil {
		fmt.Printf("error %s", err.Error())
		return
	}

	depth, position := CalculateDepthAndPosition(in)

	fmt.Printf("Multiplication value: %d\n", depth*position)
}
