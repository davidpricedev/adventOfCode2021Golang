package day4

import (
	"aoc21go/utils"
	"fmt"
	"strings"
)

func loadData(inputFile string) ([]int, []*Board) {
	lines := utils.ReadFileToLines(inputFile)
	calls := utils.AtoiMap(strings.Split(strings.TrimSpace(lines[0]), ","))
	var boards []*Board
	for i := 2; i < len(lines); i += 6 {
		boardNum := (i - 2) / 6
		//fmt.Println("i: ", i, ", boardNum: ", boardNum)
		//fmt.Println("lines: ", lines[i:i+5])
		boards = append(boards, NewBoard(lines[i:i+5], boardNum))
	}
	return calls, boards
}

func part1(inputFile string) int {
	calls, boards := loadData(inputFile)
	for _, call := range calls {
		//fmt.Println("call: ", call)
		for _, board := range boards {
			board.ApplyCall(call)
			if board.HasBingo() {
				return board.CalculateScore(call)
			}
		}
	}

	PrintBoards(boards)
	return -1
}

func part2(inputFile string) int {
	calls, boards := loadData(inputFile)
	totalBoards := len(boards)
	bingos := 0
	for _, call := range calls {
		for _, board := range boards {
			if !board.HasBingo() {
				board.ApplyCall(call)
				if board.HasBingo() {
					bingos += 1
				}
			}

			// we just logged the final bingo
			if bingos == totalBoards {
				return board.CalculateScore(call)
			}
		}
	}

	PrintBoards(boards)
	return -1
}

func Day4() {
	const inputFile = "day4/input.txt"
	fmt.Println("part1: ", part1(inputFile))
	fmt.Println("part2: ", part2(inputFile))
}

func PrintBoards(boards []*Board) {
	for _, x := range boards {
		x.InspectFull()
	}
}
