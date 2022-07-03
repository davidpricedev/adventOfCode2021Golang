package day9

import (
	"aoc21go/utils"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

func Day9() {
	const inputFile = "day9/input.txt"
	//const inputFile = "day9/inputSample.txt"
	fmt.Println("part1: ", part1(inputFile))
	fmt.Println("part2: ", part2(inputFile))
}

func loadData(file string) *Table2d {
	lines := utils.ReadFileToLines(file)
	return NewTable2dFromFileData(lines)
}

func part1(file string) int {
	table := loadData(file)
	lowpoints := getLowPoints(table)
	var lowvalues []int
	for _, p := range lowpoints {
		value, _ := table.Get(p.x, p.y)
		lowvalues = append(lowvalues, value)
	}

	return utils.SumInts(lowvalues) + len(lowpoints)
}

func getLowPoints(table *Table2d) []*point2d {
	var lowpoints []*point2d
	for j := 0; j < table.height; j++ {
		for i := 0; i < table.width; i++ {
			if table.IsLocalMin(i, j) {
				lowpoints = append(lowpoints, NewPoint(i, j))
			}
		}
	}

	return lowpoints
}

func part2(file string) int {
	table := loadData(file)
	lowpoints := getLowPoints(table)
	var sizes []int
	for _, p := range lowpoints {
		sizes = append(sizes, getBasinSize(table, *p))
	}

	sort.Ints(sizes)
	biggest3 := sizes[len(sizes)-3:]
	return utils.MultInts(biggest3)
}

func getBasinSize(table *Table2d, point point2d) int {
	initBasinPoints := mapset.NewSet(point)
	initCandidates := table.GetAdjBasinUphill(point)
	//fmt.Println("initBasin: ", initBasinPoints, ", initCandidates: ", initCandidates)
	return growBasin(table, initBasinPoints, mapset.NewSet(initCandidates...))
}

// recursive algo
func growBasin(table *Table2d, inputBasinPoints mapset.Set[point2d], inputCandidates mapset.Set[point2d]) int {
	//fmt.Println("basin points: ", inputBasinPoints, ", candidates: ", inputCandidates)

	// end condition
	if inputCandidates.Cardinality() == 0 {
		return inputBasinPoints.Cardinality()
	}

	// progress the basin points
	//basinPoints := mapset.NewSet[*point2d]()
	basinPoints := inputBasinPoints.Union(inputCandidates)

	// generate new candidates
	newCandidates := mapset.NewSet[point2d]()
	for point := range inputCandidates.Iter() {
		// the difference call is necessary to break up infinite loops for cases when the terain is flat
		adjs := mapset.NewSet(table.GetAdjBasinUphill(point)...).Difference(basinPoints)
		newCandidates = newCandidates.Union(adjs)
	}

	// tail rec call
	return growBasin(table, basinPoints, newCandidates)
}
