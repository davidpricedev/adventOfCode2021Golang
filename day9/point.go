package day9

type point2d struct {
	x int
	y int
}

func NewPoint(x int, y int) *point2d {
	return &point2d{x, y}
}
