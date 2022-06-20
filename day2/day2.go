package day2

import (
	"aoc21go/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(data []string) int {
	pos, dep := 0, 0
	for _, line := range data {
		value, _ := strconv.Atoi(strings.Split(line, " ")[1])
		if strings.HasPrefix(line, "up") {
			dep -= value
		} else if strings.HasPrefix(line, "down") {
			dep += value
		} else if strings.HasPrefix(line, "forward") {
			pos += value
		} else {
			fmt.Println("unexpected command: ", line)
		}

		//fmt.Println("position: ", pos, " depth: ", dep)
	}

	return pos * dep
}

func part2(data []string) int {
	pos, dep, aim := 0, 0, 0
	for _, line := range data {
		value, _ := strconv.Atoi(strings.Split(line, " ")[1])
		if strings.HasPrefix(line, "up") {
			aim -= value
		} else if strings.HasPrefix(line, "down") {
			aim += value
		} else if strings.HasPrefix(line, "forward") {
			pos += value
			dep += value * aim
		} else {
			fmt.Println("unexpected command: ", line)
		}

		//fmt.Println("aim: ", aim, "position: ", pos, " depth: ", dep)
	}

	return pos * dep
}

func Day2() {
	const inputFile = "day2/input.txt"
	data := utils.ReadFileToLines(inputFile)
	fmt.Println("part1: " + strconv.Itoa(part1(data)))
	fmt.Println("part2: " + strconv.Itoa(part2(data)))
}
