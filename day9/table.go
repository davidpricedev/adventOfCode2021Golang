package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Table2d struct {
	tableData [][]int
	height    int
	width     int
}

func NewTable2dFromFileData(filedata []string) *Table2d {
	var tableData [][]int
	height := len(filedata)
	var width int
	for _, line := range filedata {
		cleanLine := strings.TrimSpace(line)
		var lineData []int
		for _, xstr := range strings.Split(cleanLine, "") {
			x, _ := strconv.Atoi(xstr)
			lineData = append(lineData, x)
		}
		width = len(lineData)
		tableData = append(tableData, lineData)
	}

	return &Table2d{tableData, height, width}
}

func (table Table2d) Get(x int, y int) (int, error) {
	if y < 0 || y > table.height {
		return -99, fmt.Errorf("y (%d) is out of bounds (0, %d)", y, table.height)
	}
	if x < 0 || x > table.width {
		return -99, fmt.Errorf("x (%d) is out of bounds (0, %d)", x, table.width)
	}

	return table.tableData[y][x], nil
}

// Return adjacent points and values
func (table Table2d) GetAdjacent(x int, y int) ([]*point2d, []int) {
	var points []*point2d
	var values []int
	// west
	if x > 0 {
		points = append(points, NewPoint(x-1, y))
		value, _ := table.Get(x-1, y)
		values = append(values, value)
	}
	// east
	if x < (table.width - 1) {
		points = append(points, NewPoint(x+1, y))
		value, _ := table.Get(x+1, y)
		values = append(values, value)
	}
	// north
	if y > 0 {
		points = append(points, NewPoint(x, y-1))
		value, _ := table.Get(x, y-1)
		values = append(values, value)
	}
	// south
	if y < (table.height - 1) {
		points = append(points, NewPoint(x, y+1))
		value, _ := table.Get(x, y+1)
		values = append(values, value)
	}

	return points, values
}

// If we stipulate not only within the basin but also uphill, then we'll have less worry about duplication
func (table Table2d) GetAdjBasinUphill(point point2d) []point2d {
	currentValue, _ := table.Get(point.x, point.y)
	adjPointsRaw, _ := table.GetAdjacent(point.x, point.y)
	var adjPoints []point2d
	for _, p := range adjPointsRaw {
		value, _ := table.Get(p.x, p.y)
		if value < 9 && value >= currentValue {
			adjPoints = append(adjPoints, *p)
		}
	}

	return adjPoints
}

func (table Table2d) IsLocalMin(x int, y int) bool {
	value, _ := table.Get(x, y)
	if value == 9 {
		return false
	}

	_, adjValues := table.GetAdjacent(x, y)
	islower := true
	for _, z := range adjValues {
		islower = islower && (value < z)
	}

	return islower
}
