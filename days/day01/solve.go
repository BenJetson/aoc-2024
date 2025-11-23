package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
)

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	left := make([]int, len(input))
	right := make([]int, len(input))
	for idx, line := range input {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			err = fmt.Errorf("expected 2 parts on line %d", idx+1)
			return
		} else if left[idx], err = strconv.Atoi(parts[0]); err != nil {
			err = fmt.Errorf("invalid left int on line %d: %w", idx+1, err)
		} else if right[idx], err = strconv.Atoi(parts[1]); err != nil {
			err = fmt.Errorf("invalid right int on line %d: %w", idx+1, err)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance int
	for idx := 0; idx < len(input); idx++ {
		minVal := min(left[idx], right[idx])
		maxVal := max(left[idx], right[idx])
		distance := maxVal - minVal
		totalDistance += distance
	}

	s.Part1.SaveIntAnswer(totalDistance)

	return
}
