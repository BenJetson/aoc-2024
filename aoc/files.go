package aoc

import (
	"fmt"
	"os"

	"github.com/BenJetson/aoc-2024/utilities"
)

func GetDayDirectory(day int) string {
	return fmt.Sprintf("days/day%02d", day)
}

// A ProblemSet describes which set of input and solutions to use.
type ProblemSet string

const (
	// ProblemSetMy represents the player's individual input and solution.
	ProblemSetMy ProblemSet = "my"
	// ProblemSetExample represents the example input and solution.
	ProblemSetExample ProblemSet = "example"
)

func GetInputFilename(day int, ps ProblemSet) string {
	return fmt.Sprintf("%s/%s_input.txt", GetDayDirectory(day), ps)
}

func GetInput(day int, ps ProblemSet) (Input, error) {
	inputFilename := GetInputFilename(day, ps)
	return utilities.ReadLinesFromFile(inputFilename)
}

func GetSolutionFilename(day int, ps ProblemSet) string {
	return fmt.Sprintf("%s/%s_solution.txt", GetDayDirectory(day), ps)
}

func scanAnswer(line, label string) (Answer, error) {
	var a Answer
	var foundLabel string
	var valueStr string

	n, err := fmt.Sscanf(line, "Part %s answer is: %s",
		&foundLabel, &valueStr)
	if err != nil {
		return a, fmt.Errorf("failed to scan answer: %w", err)
	} else if n != 2 {
		return a, fmt.Errorf("expect two scanned values, found: %d", n)
	}

	if foundLabel != label {
		return a, fmt.Errorf("found label '%s' does not match '%s'",
			foundLabel, label)
	}

	if valueStr != "blank" {
		a.Value = valueStr
		a.Valid = true
	}

	return a, nil
}

func GetSolution(day int, ps ProblemSet) (Solution, error) {
	var s Solution

	solutionFilename := GetSolutionFilename(day, ps)
	lines, err := utilities.ReadLinesFromFile(solutionFilename)
	if err != nil {
		return s, fmt.Errorf(
			"could not read solution file '%s': %w", solutionFilename, err)
	} else if len(lines) != 2 {
		return s, fmt.Errorf(
			"expected two lines from solution file '%s', found %d",
			solutionFilename, len(lines))
	}

	s.Part1, err = scanAnswer(lines[0], "one")
	if err != nil {
		return s, fmt.Errorf("could not get part one answer: %w", err)
	}
	s.Part2, err = scanAnswer(lines[1], "two")
	if err != nil {
		return s, fmt.Errorf("could not get part two answer: %w", err)
	}

	return s, nil
}

func WriteSolution(day int, ps ProblemSet, s Solution) error {
	solutionFilename := GetSolutionFilename(day, ps)
	solutionStr := s.String()

	err := os.WriteFile(solutionFilename, []byte(solutionStr), 0644)
	if err != nil {
		return fmt.Errorf("failed to write solution file '%s': %w",
			solutionFilename, err)
	}
	return nil
}
