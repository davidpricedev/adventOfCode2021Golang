package day8

import (
	"aoc21go/utils"
	"fmt"
)

func Day8() {
	const inputFile = "day8/input.txt"
	fmt.Println("part1: ", part1(inputFile))
	//fmt.Println("part2: ", part2(inputFile))
}

func loadData(filename string) []Sample {
	lines := utils.ReadFileToLines(filename)
	var ios []Sample
	for _, line := range lines {
		ios = append(ios, *NewSampleFromString(line))
	}

	return ios
}

func part1(filename string) int {
	samples := loadData(filename)
	count := 0
	for _, sample := range samples {
		for _, x := range sample.Outputs {
			if len(x) == 2 || len(x) == 3 || len(x) == 4 || len(x) == 7 {
				count += 1
			}
		}
	}

	return count
}
