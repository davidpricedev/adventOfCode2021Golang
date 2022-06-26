package day5

import (
	"aoc21go/utils"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func NewPointFromString(coordStr string) *Point {
	coords := utils.AtoiMap(strings.Split(strings.TrimSpace(coordStr), ","))
	return NewPoint(coords[0], coords[1])
}

func (point Point) Inspect() {
	fmt.Println("Point (", point.x, ",", point.y, ")")
}

func (point Point) toString() string {
	return fmt.Sprintf("Point(%d,%d)", point.x, point.y)
}

type Line struct {
	start Point
	end   Point
}

func NewLineFromString(lineText string) *Line {
	pointsArr := strings.Split(strings.TrimSpace(lineText), " -> ")
	return &Line{*NewPointFromString(pointsArr[0]), *NewPointFromString(pointsArr[1])}
}

func (line Line) Inspect() {
	fmt.Println("Line (", line.start.toString(), " -> ", line.end.toString(), ")")
}

func (line Line) StartX() int {
	return line.start.x
}

func (line Line) StartY() int {
	return line.start.y
}

func (line Line) EndX() int {
	return line.end.x
}

func (line Line) EndY() int {
	return line.end.y
}

func (line Line) MinX() int {
	if line.start.x <= line.end.x {
		return line.start.x
	} else {
		return line.end.x
	}
}

func (line Line) MinY() int {
	if line.start.y <= line.end.y {
		return line.start.y
	} else {
		return line.end.y
	}
}

func (line Line) MaxX() int {
	if line.start.x > line.end.x {
		return line.start.x
	} else {
		return line.end.x
	}
}

func (line Line) MaxY() int {
	if line.start.y > line.end.y {
		return line.start.y
	} else {
		return line.end.y
	}
}

func (line Line) IsCardinal() bool {
	return line.StartX() == line.EndX() || line.StartY() == line.EndY()
}

func (line Line) ToPoints() []*Point {
	xd := line.MaxX() - line.MinX()
	yd := line.MaxY() - line.MinY()
	var length int
	if xd > yd {
		length = xd
	} else {
		length = yd
	}

	var xinc int
	if line.start.x == line.end.x {
		xinc = 0
	} else if line.start.x < line.end.x {
		xinc = 1
	} else {
		xinc = -1
	}

	var yinc int
	if line.start.y == line.end.y {
		yinc = 0
	} else if line.start.y < line.end.y {
		yinc = 1
	} else {
		yinc = -1
	}

	var linePoints []*Point
	for i, j := line.StartX(), line.StartY(); len(linePoints) <= length; i, j = i+xinc, j+yinc {
		linePoints = append(linePoints, NewPoint(i, j))
	}

	return linePoints
}
