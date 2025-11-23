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

func (r Report) Without(badIdx int) Report {
	r2 := make(Report, 0, len(r)-1)
	for idx, level := range r {
		if idx != badIdx {
			r2 = append(r2, level)
		}
	}
	return r2
}

func (r Report) IsSafeWithDampener() bool {
	if r.IsSafe() {
		return true // already safe; nothing to do.
	}

	for idx := range r {
		if r.Without(idx).IsSafe() {
			return true
		}
	}

	return false
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var numSafe, numSafeWithDampener int
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

		if report.IsSafeWithDampener() {
			numSafeWithDampener++
		}
	}

	s.Part1.SaveIntAnswer(numSafe)
	s.Part2.SaveIntAnswer(numSafeWithDampener)

	return
}
