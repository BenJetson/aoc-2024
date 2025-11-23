package days

import (
	"github.com/BenJetson/aoc-2024/aoc"
	// BEGIN DAY IMPORTS
	"github.com/BenJetson/aoc-2024/days/day01"
	"github.com/BenJetson/aoc-2024/days/day02"
	// END DAY IMPORTS
)

var Solvers = map[int]aoc.Solver{
	// BEGIN DAY SOLVERS
	// 1: day01.SolvePuzzle,
	1: day01.SolvePuzzle,
	2: day02.SolvePuzzle,
	// END DAY SOLVERS
}
