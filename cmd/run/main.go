package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BenJetson/aoc-2024/aoc"
	"github.com/BenJetson/aoc-2024/solver"
)

var dayFlag = flag.Int("day", 0, "day of the advent calendar, 1-25")
var exampleFlag = flag.Bool("example", false,
	"optionally use example input")
var inputFileFlag = flag.String("input", "",
	"optional filename to use for input, other than the default")

func main() {
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		log.Fatal("invalid or missing AoC day number")
	}

	if *exampleFlag {
		*inputFileFlag = aoc.GetExampleFilename(*dayFlag)
	} else if len(*inputFileFlag) < 1 {
		*inputFileFlag = aoc.GetInputFilename(*dayFlag)
	}

	solution, err := solver.RunForDayWithInput(*dayFlag, *inputFileFlag)
	if err != nil {
		log.Fatalf("error while solving puzzle: %v\n", err)
	}

	fmt.Print(solution.String())
}
