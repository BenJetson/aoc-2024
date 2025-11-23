package day02

import (
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
	"github.com/BenJetson/aoc-2024/utilities"
)

type Report []int

type Mode int

const (
	ModeNoChange Mode = iota
	ModeDecreasing
	ModeIncreasing
)

func GetMode(a, b int) Mode {
	switch {
	case a > b:
		return ModeDecreasing
	case a < b:
		return ModeIncreasing
	default:
		return ModeNoChange
	}
}

func (r Report) IsSafe() bool {
	var modeForReport Mode
	for idx := 0; idx < len(r)-1; idx++ {
		a, b := r[idx], r[idx+1]
		m := GetMode(a, b)
		if m == ModeNoChange {
			return false
		} else if idx == 0 {
			modeForReport = m
		} else if m != modeForReport {
			return false
		}

		diff := max(a, b) - min(a, b)
		if diff > 3 {
			return false
		}
	}
	return true
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var numSafe int
	for _, line := range input {
		var levels []int
		levels, err = utilities.SliceStringsToInts(strings.Split(line, " "))
		if err != nil {
			return
		}
		report := Report(levels)

		if report.IsSafe() {
			numSafe++
		}
	}

	s.Part1.SaveIntAnswer(numSafe)

	return
}
