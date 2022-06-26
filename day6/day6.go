package day6

import (
	"aoc21go/utils"
	"fmt"
	"strings"
)

func Day6() {
	const inputFile = "day6/input.txt"
	fmt.Println("part1: ", part1(inputFile))
	fmt.Println("part2: ", part2(inputFile))
}

func part1(filename string) int {
	initialAges := utils.AtoiMap(strings.Split(strings.TrimSpace(utils.ReadFile(filename)), ","))
	ageMap := make(map[int]int)
	for _, age := range initialAges {
		ageMap[age] += 1
	}

	ageMap = simulateForDays(80, ageMap)
	return countPop(ageMap)
}

func part2(filename string) int {
	initialAges := utils.AtoiMap(strings.Split(strings.TrimSpace(utils.ReadFile(filename)), ","))
	ageMap := make(map[int]int)
	for _, age := range initialAges {
		ageMap[age] += 1
	}

	ageMap = simulateForDays(256, ageMap)
	return countPop(ageMap)
}

func simulateForDays(days int, ageMap map[int]int) map[int]int {
	for i := 0; i < days; i++ {
		//fmt.Println("Day =", i, ", population =", countPop(ageMap), ": ")
		//fmt.Println(ageMap)
		ageMap = runDay(ageMap)
	}

	return ageMap
}

// Have to do += instead of just = to ensure the new resets don't overwrite
// those who are just aging into slot 6
func runDay(ageMap map[int]int) map[int]int {
	newAgeMap := make(map[int]int)
	for age, count := range ageMap {
		if age == 0 {
			newAgeMap[6] += count
			newAgeMap[8] = count
		} else {
			newAgeMap[age-1] += count
		}
	}

	return newAgeMap
}

func countPop(ageMap map[int]int) int {
	sum := 0
	for _, count := range ageMap {
		sum += count
	}

	return sum
}
