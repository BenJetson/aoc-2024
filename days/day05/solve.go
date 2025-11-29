package day05

import (
	"fmt"
	"slices"
	"strings"

	"github.com/BenJetson/aoc-2024/aoc"
	"github.com/BenJetson/aoc-2024/utilities"
)

type PageNumber int

type Rule struct {
	Page  PageNumber
	After PageNumber
}

func lineToRule(line string) (Rule, error) {
	values, err := utilities.SliceStringsToIntsLike[PageNumber](
		strings.Split(line, "|"))
	if err != nil {
		return Rule{}, err
	} else if len(values) != 2 {
		return Rule{}, fmt.Errorf("unexpected rule values (%d)", len(values))
	}
	return Rule{
		Page:  values[0],
		After: values[1],
	}, nil
}

type UpdatedManual []PageNumber

func (u UpdatedManual) MiddlePage() PageNumber {
	if len(u)%2 == 0 {
		panic("unexpected manual with even page count")
	}
	idx := (len(u) / 2) // should guaranteed to be odd page count.
	return u[idx]
}

func lineToUpdatedManual(line string) (UpdatedManual, error) {
	pages, err := utilities.SliceStringsToIntsLike[PageNumber](
		strings.Split(line, ","))
	if err != nil {
		return nil, err
	}
	return UpdatedManual(pages), nil
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var rules []Rule
	var updates []UpdatedManual
	pastSeparator := false

	for idx, line := range input {
		switch {
		case line == "":
			pastSeparator = true
		case !pastSeparator:
			var r Rule
			r, err = lineToRule(line)
			if err != nil {
				err = fmt.Errorf("invalid rule on line %d: %w", idx+1, err)
				return
			}
			rules = append(rules, r)
		default:
			var u UpdatedManual
			u, err = lineToUpdatedManual(line)
			if err != nil {
				err = fmt.Errorf("invalid update on line %d: %w", idx+1, err)
				return
			}
			updates = append(updates, u)
		}
	}

	// XXX these linear scans are inefficient, but get the job done.
	var acc int
	for _, u := range updates {
		var invalid bool
		for idx, page := range u {
			for _, r := range rules {
				if page == r.Page {
					otherIdx := slices.Index(u, r.After)
					if otherIdx >= 0 && idx > otherIdx {
						invalid = true
					}
				}
			}

		}

		if !invalid {
			acc += int(u.MiddlePage())
		}
	}

	s.Part1.SaveIntAnswer(acc)

	return
}
