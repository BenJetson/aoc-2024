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
	Page   PageNumber
	Before PageNumber
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
		Page:   values[0],
		Before: values[1],
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

func (u UpdatedManual) IsValid(rules []Rule) bool {
	for _, r := range rules {
		idx := slices.Index(u, r.Page)
		otherIdx := slices.Index(u, r.Before)
		if idx < 0 || otherIdx < 0 {
			continue
		} else if otherIdx < idx {
			return false
		}
	}
	return true
}

func (u UpdatedManual) Correct(rules []Rule) (m UpdatedManual) {
	m = make(UpdatedManual, 0, len(u))
	m = append(m, u...)

	for !m.IsValid(rules) {
		for _, r := range rules {
			idx := slices.Index(m, r.Page)
			otherIdx := slices.Index(m, r.Before)
			if idx < 0 || otherIdx < 0 {
				continue
			} else if otherIdx > idx {
				continue
			}

			m = slices.Delete(m, otherIdx, otherIdx+1)
			m = slices.Insert(m, idx, r.Before)
		}
	}

	return m
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
	// Using a map for rules would be better.
	var acc, acc2 int
	for _, u := range updates {
		if u.IsValid(rules) {
			acc += int(u.MiddlePage())
		} else {
			m := u.Correct(rules)
			acc2 += int(m.MiddlePage())
		}
	}

	s.Part1.SaveIntAnswer(acc)
	s.Part2.SaveIntAnswer(acc2)

	return
}
