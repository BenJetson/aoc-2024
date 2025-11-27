package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
)

var mulInstrPattern = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	memory := strings.Join(input, "")
	matches := mulInstrPattern.FindAllStringSubmatch(memory, -1)

	var acc int
	for _, match := range matches {
		var a, b int
		if a, err = strconv.Atoi(match[1]); err != nil {
			return
		} else if b, err = strconv.Atoi(match[2]); err != nil {
			return
		}

		acc += a * b
	}

	s.Part1.SaveIntAnswer(acc)
	return
}
