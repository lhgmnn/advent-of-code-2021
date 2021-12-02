package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Input struct {
	Values []string
}

func NewFromFile(filename string) (*Input, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	i := Input{}

	defer file.Close()

	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err.Error())
	}

	r := bufio.NewReader(file)
	for {
		value, err := r.ReadString('\n')
		if len(value) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		i.Values = append(i.Values, strings.TrimSuffix(value, "\n"))
	}

	return &i, nil
}
