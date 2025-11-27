package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
)

var doPrefixInstrPattern = regexp.MustCompile(`^do\(\)`)
var dontPrefixInstrPattern = regexp.MustCompile(`^don't\(\)`)
var mulPrefixInstrPattern = regexp.MustCompile(`^mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
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

	var doAcc int
	canDo := true

	for len(memory) > 0 {
		match := doPrefixInstrPattern.FindStringIndex(memory)
		if match != nil {
			canDo = true
			memory = memory[match[1]:]
			continue
		}

		match = dontPrefixInstrPattern.FindStringIndex(memory)
		if match != nil {
			canDo = false
			memory = memory[match[1]:]
			continue
		}

		match = mulPrefixInstrPattern.FindStringSubmatchIndex(memory)
		if match != nil {
			if canDo {
				a, err = strconv.Atoi(memory[match[2]:match[3]])
				if err != nil {
					return
				}

				b, err = strconv.Atoi(memory[match[4]:match[5]])
				if err != nil {
					return
				}

				doAcc += a * b
			}

			memory = memory[match[1]:]
			continue
		}

		memory = memory[1:]
	}

	s.Part2.SaveIntAnswer(doAcc)

	return
}
