package day3

import (
	"aoc21go/utils"
	"fmt"
	"strconv"
	"strings"
)

func extractGammaEpsilon(data []string) ([]string, []string) {
	var oneCounts []int
	for i, line := range data {
		values := strings.Split(line, "")
		if i == 0 {
			for _, s := range values {
				x, _ := strconv.Atoi(s)
				oneCounts = append(oneCounts, x)
			}
		} else {
			for j, s := range values {
				x, _ := strconv.Atoi(s)
				oneCounts[j] += x
			}
		}
		//fmt.Println(oneCounts)
	}
	half := float64(len(data)) / float64(2)
	var gammaArray []string
	var epsilonArray []string
	//fmt.Println("oneCounts: ", oneCounts, ", length: ", len(data), ", half: ", half)
	for _, x := range oneCounts {
		if float64(x) >= half {
			gammaArray = append(gammaArray, "1")
			epsilonArray = append(epsilonArray, "0")
		} else {
			gammaArray = append(gammaArray, "0")
			epsilonArray = append(epsilonArray, "1")
		}
	}
	return gammaArray, epsilonArray
}

func part1(data []string) int64 {
	gammaArray, epsilonArray := extractGammaEpsilon(data)
	gamma, _ := strconv.ParseInt(strings.Join(gammaArray, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonArray, ""), 2, 64)
	return gamma * epsilon
}

func part2(data []string) int64 {
	var lastFilteredO2 []string = data
	var lastFilteredCO2 []string = data
	for i := 0; i < len(data[0]); i++ {
		newFilteredO2 := []string{}
		gamma, _ := extractGammaEpsilon(lastFilteredO2)
		//fmt.Println("gamma: ", gamma)
		if len(lastFilteredO2) > 1 {
			for _, line := range lastFilteredO2 {
				values := strings.Split(line, "")
				if values[i] == gamma[i] {
					newFilteredO2 = append(newFilteredO2, line)
				}
			}
			//fmt.Println("next o2: ", newFilteredO2)
			lastFilteredO2 = newFilteredO2
		}
		newFilteredCO2 := []string{}
		_, epsilon := extractGammaEpsilon(lastFilteredCO2)
		//fmt.Println("epsilon: ", epsilon)
		if len(lastFilteredCO2) > 1 {
			for _, line := range lastFilteredCO2 {
				values := strings.Split(line, "")
				if values[i] == epsilon[i] {
					newFilteredCO2 = append(newFilteredCO2, line)
				}
			}
			//fmt.Println("next co2: ", newFilteredCO2)
			lastFilteredCO2 = newFilteredCO2
		}
	}
	//fmt.Println("o2: ", lastFilteredO2, "co2: ", lastFilteredCO2)
	o2, _ := strconv.ParseInt(lastFilteredO2[0], 2, 64)
	co2, _ := strconv.ParseInt(lastFilteredCO2[0], 2, 64)
	return o2 * co2
}

func Day3() {
	const inputFile = "day3/input.txt"
	data := utils.ReadFileToLines(inputFile)
	fmt.Println("part1: " + strconv.FormatInt(part1(data), 10))
	fmt.Println("part2: " + strconv.FormatInt(part2(data), 10))
}
