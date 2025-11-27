package day04

import (
	"github.com/BenJetson/aoc-2024/aoc"
)

type Direction int

const (
	DirectionNorth Direction = iota
	DirectionNorthEast
	DirectionEast
	DirectionSouthEast
	DirectionSouth
	DirectionSouthWest
	DirectionWest
	DirectionNorthWest
)

func (dir Direction) String() string {
	switch dir {
	case DirectionNorth:
		return "N_"
	case DirectionNorthEast:
		return "NE"
	case DirectionEast:
		return "E_"
	case DirectionSouthEast:
		return "SE"
	case DirectionSouth:
		return "S_"
	case DirectionSouthWest:
		return "SW"
	case DirectionWest:
		return "W_"
	case DirectionNorthWest:
		return "NW"
	}
	return "__"
}

var AllDirections = []Direction{
	DirectionNorth,
	DirectionNorthEast,
	DirectionEast,
	DirectionSouthEast,
	DirectionSouth,
	DirectionSouthWest,
	DirectionWest,
	DirectionNorthWest,
}

type Grid [][]rune

func (g Grid) ColCount() int {
	return len(g[0])
}

func (g Grid) RowCount() int {
	return len(g)
}

func (g Grid) InRange(xPos, yPos int) bool {
	return xPos > -1 && xPos < g.ColCount() && yPos > -1 && yPos < g.RowCount()
}

func (g Grid) GetString(
	xPos, yPos int,
	dir Direction,
	length int,
) (out string, ok bool) {
	// defer func(xPosOg, yPosOg int, s *string) {
	// 	fmt.Printf("X %03d Y %03d D %s: %s (%t)\n",
	// 		xPosOg, yPosOg, dir.String(), *s, ok)
	// }(xPos, yPos, &out)
	if length < 1 {
		return "", false
	}

	var xMod int
	switch dir {
	case DirectionNorthEast, DirectionEast, DirectionSouthEast:
		xMod = 1
	case DirectionNorthWest, DirectionWest, DirectionSouthWest:
		xMod = -1
	}

	var yMod int
	switch dir {
	case DirectionNorthWest, DirectionNorth, DirectionNorthEast:
		yMod = -1
	case DirectionSouthWest, DirectionSouth, DirectionSouthEast:
		yMod = 1
	}

	chars := make([]rune, length)
	for idx := 0; idx < length; idx++ {
		if idx != 0 {
			xPos += xMod
			yPos += yMod
		}

		// fmt.Printf("check X %03d Y %03d I %t\n", xPos, yPos, g.InRange(xPos, yPos))

		if !g.InRange(xPos, yPos) {
			return
		}

		chars[idx] = g[yPos][xPos]
	}

	out = string(chars)
	ok = true
	return
}

func (g Grid) CheckForMatch(
	xPos, yPos int,
	dir Direction,
	target string,
) bool {
	s, ok := g.GetString(xPos, yPos, dir, len(target))
	if !ok {
		return false
	}

	return s == target
}

func (g Grid) GetMatchCount(xPos, yPos int, target string) (count int) {
	for _, dir := range AllDirections {
		if g.CheckForMatch(xPos, yPos, dir, target) {
			count++
		}
	}
	return
}

func (g Grid) IsMasX(xPos, yPos int) (ok bool) {
	var diag1 string
	if diag1, ok = g.GetString(xPos-1, yPos-1, DirectionSouthEast, 3); !ok {
		return
	}

	var diag2 string
	if diag2, ok = g.GetString(xPos+1, yPos-1, DirectionSouthWest, 3); !ok {
		return
	}

	return (diag1 == "MAS" || diag1 == "SAM") &&
		(diag2 == "MAS" || diag2 == "SAM")
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var g Grid
	for _, line := range input {
		g = append(g, []rune(line))
	}

	const target = "XMAS"
	var count int
	for yPos := 0; yPos < g.RowCount(); yPos++ {
		for xPos := 0; xPos < g.ColCount(); xPos++ {
			if g[yPos][xPos] == 'X' {
				count += g.GetMatchCount(xPos, yPos, target)
			}
		}
	}

	s.Part1.SaveIntAnswer(count)

	count = 0
	for yPos := 0; yPos < g.RowCount(); yPos++ {
		for xPos := 0; xPos < g.ColCount(); xPos++ {
			if g[yPos][xPos] == 'A' && g.IsMasX(xPos, yPos) {
				count++
			}
		}
	}

	s.Part2.SaveIntAnswer(count)

	return
}
