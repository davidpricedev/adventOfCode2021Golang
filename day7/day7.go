package day7

import (
	"aoc21go/utils"
	"fmt"
	"math"
	"strings"
)

func Day7() {
	const inputFile = "day7/input.txt"
	fmt.Println("part1: ", part1(inputFile))
	fmt.Println("part2: ", part2(inputFile))
}

func loadData(filename string) (map[int]int, int) {
	initialPositions := utils.AtoiMap(strings.Split(strings.TrimSpace(utils.ReadFile(filename)), ","))
	positionMap := make(map[int]int)
	maxPos := 0
	for _, pos := range initialPositions {
		if maxPos < pos {
			maxPos = pos
		}

		positionMap[pos] += 1
	}

	return positionMap, maxPos
}

func part1(filename string) int {
	positionMap, maxPos := loadData(filename)
	costMap := make(map[int]int)
	for i := 0; i <= maxPos; i++ {
		costMap[i] = calculatePart1Cost(positionMap, i)
	}

	return getLowestCost(costMap)
}

func part2(filename string) int {
	positionMap, maxPos := loadData(filename)
	costMap := make(map[int]int)
	for i := 0; i <= maxPos; i++ {
		costMap[i] = calculatePart2Cost(positionMap, i)
	}

	return getLowestCost(costMap)
}

func calculatePart1Cost(positionMap map[int]int, destPos int) int {
	cost := 0
	for pos, count := range positionMap {
		distance := int(math.Abs(float64(pos - destPos)))
		cost += distance * count
	}

	return cost
}

func calculatePart2Cost(positionMap map[int]int, destPos int) int {
	cost := 0
	for pos, count := range positionMap {
		distance := int(math.Abs(float64(pos - destPos)))
		singleCost := distance * (distance + 1) / 2
		cost += singleCost * count
	}

	return cost
}

func getLowestCost(costMap map[int]int) int {
	minCost := costMap[0]
	for _, cost := range costMap {
		if minCost > cost {
			minCost = cost
		}
	}

	return minCost
}
