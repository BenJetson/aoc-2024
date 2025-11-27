package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
)

// var doInstrPattern = regexp.MustCompile(`do\(\)`)
// var dontInstrPattern = regexp.MustCompile(`don't\(\)`)
var modifierInstrPattern = regexp.MustCompile(`do(n')?t\(\)`)
var mulInstrPattern = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	memory := strings.Join(input, "")
	matches := mulInstrPattern.FindAllStringSubmatch(memory, -1)

	var acc int
	var a, b int
	for _, match := range matches {
		if a, err = strconv.Atoi(match[1]); err != nil {
			return
		} else if b, err = strconv.Atoi(match[2]); err != nil {
			return
		}

		acc += a * b
	}

	s.Part1.SaveIntAnswer(acc)

	modifierMatches := modifierInstrPattern.FindAllStringIndex(memory, -1)
	modifierIndices := make([]int, 1, len(modifierMatches)+2)
	modifierIndices[0] = 0
	for _, matchIndices := range modifierMatches {
		modifierIndices = append(modifierIndices, matchIndices[0])
	}
	modifierIndices = append(modifierIndices, len(memory))

	var memoryRanges []string
	for idx := 0; idx < len(modifierIndices)-1; idx++ {
		a, b = modifierIndices[idx], modifierIndices[idx+1]
		memoryRanges = append(memoryRanges, memory[a:b])
	}

	var doAcc int
	canDo := true

	for _, memoryRange := range memoryRanges {
		if strings.HasPrefix(memoryRange, "do()") {
			canDo = true
		} else if strings.HasPrefix(memoryRange, "don't()") {
			canDo = false
		}

		if canDo {
			matches := mulInstrPattern.FindAllStringSubmatch(memoryRange, -1)
			for _, match := range matches {
				if a, err = strconv.Atoi(match[1]); err != nil {
					return
				} else if b, err = strconv.Atoi(match[2]); err != nil {
					return
				}

				doAcc += a * b
			}
		}
	}

	s.Part2.SaveIntAnswer(doAcc)

	return
}
