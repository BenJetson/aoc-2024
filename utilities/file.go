package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
)

const linefeed = '\n'

// CreateEmptyFile creates an empty file, then immediately closes its file
// handle.
func CreateEmptyFile(filename string) error {
	if f, err := os.Create(filename); err != nil {
		return fmt.Errorf("could not create file: %w", err)
	} else if err = f.Close(); err != nil {
		return fmt.Errorf("could not close file: %w", err)
	}
	return nil
}

// ReadLinesFromFile reads the entire file specified, and returns a slice
// of strings
func ReadLinesFromFile(filename string) ([]string, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	var line string

	for _, b := range raw {
		if b == linefeed {
			lines = append(lines, line)
			line = ""

			continue
		}

		line += string(b)
	}

	return lines, nil
}

// ReadIntegersFromFile reads a file and attempts to parse each line as an
// integer, returning the resultant slice of integers.
func ReadIntegersFromFile(filename string) ([]int, error) {
	lines, err := ReadLinesFromFile(filename)
	if err != nil {
		return nil, err
	}

	nums, err := SliceStringsToInts(lines)
	if err != nil {
		return nil, err
	}

	return nums, nil
}
