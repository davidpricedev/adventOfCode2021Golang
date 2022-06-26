package day5

import (
	"aoc21go/utils"
	"fmt"
)

func Day5() {
	const inputFile = "day5/input.txt"
	fmt.Println("part1: ", part1(inputFile))
	fmt.Println("part2: ", part2(inputFile))
}

func part1(inputFile string) int {
	fileLines := utils.ReadFileToLines(inputFile)
	var lines []*Line
	for _, line := range fileLines {
		newLine := NewLineFromString(line)
		if newLine.IsCardinal() {
			lines = append(lines, newLine)
		}
	}

	return calculateAnswer(lines)
}

func part2(inputFile string) int {
	fileLines := utils.ReadFileToLines(inputFile)
	var lines []*Line
	for _, line := range fileLines {
		lines = append(lines, NewLineFromString(line))
	}

	return calculateAnswer(lines)
}

func calculateAnswer(lines []*Line) int {
	pointCounts := make(map[Point]int)
	countOfMult := 0
	for _, line := range lines {
		for _, point := range line.ToPoints() {
			pointCounts[*point] += 1
			if pointCounts[*point] == 2 {
				//fmt.Println("== Found a duplicate count == 2")
				//point.Inspect()
				//line.Inspect()
				countOfMult += 1
			}
		}
	}

	return countOfMult
}
