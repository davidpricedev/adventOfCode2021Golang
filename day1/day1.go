package day1

import (
	"aoc21go/utils"
	"fmt"
	"strconv"
)

func loadData(inputFile string) []int {
	return utils.ReadFileToIntArray(inputFile)
}

// Count number of increasing pairs
func countIncreases(data []int) int {
	var increases int = 0
	for i, _ := range data[1:] {
		if data[i] < data[i+1] {
			increases = increases + 1
		}
	}
	return increases
}

func window(data []int, windowSize int) [][]int {
	var output [][]int
	for i, _ := range data[(windowSize - 1):] {
		output = append(output, data[i:i+windowSize])
	}
	return output
}

func sumWindow(data []int, windowSize int) []int {
	var output []int
	for i, _ := range data[(windowSize - 1):] {
		output = append(output, utils.Sum(data[i:i+windowSize]...))
	}
	return output
}

func part1(data []int) int {
	return countIncreases(data)
}

func part2(data []int) int {
	sumWindowed := sumWindow(data, 3)
	return countIncreases(sumWindowed)
	return 0
}

func Day1() {
	const inputFile = "day1/input.txt"
	data := loadData(inputFile)
	fmt.Println("part1: " + strconv.Itoa(part1(data)))
	fmt.Println("part2: " + strconv.Itoa(part2(data)))
}
