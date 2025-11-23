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

// IsSafe determines if a given report is safe.
// When a report is safe, the bool value is true (ignore the int).
// When a report is unsafe, the bool value is false and the int reports the
// index of the level that caused the report to be unsafe.
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

	// fmt.Println(r)

	for idx := 0; idx < len(r); idx++ {
		// r2 := slices.Delete(r, idx, idx+1)
		// r2 := append(r[:idx], r[idx+1:]...)
		r2 := r.Without(idx)

		// fmt.Println("index", idx,
		// 	"original", r,
		// 	"modified", r2,
		// 	"removed_value", r[idx],
		// 	"len_original", len(r),
		// 	"len_modified", len(r2),
		// )

		if r2.IsSafe() {
			// fmt.Println("--- ok ---")
			return true
		}
	}

	// fmt.Println("--- bad ---")
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
