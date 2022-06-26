package day4

import (
	"fmt"
	"strconv"
	"strings"
)

// Maintain 2 representations
// A 2d table of the values in table
// and 2 1d arrays counting the number of hits in each column and row
// to find if we had a bingo, we just look for any of the column counts or row counts that are 5
// the board also contains an identifier boardId
type Board struct {
	boardId    int
	table      [5][5]int
	rowHits    [5]int
	columnHits [5]int
	bingoed    bool
	// Not really used for anything apart for debugging
	calls [][]int // each entry is [row, col, call#]
}

func NewBoard(lines []string, boardId int) *Board {
	var table [5][5]int
	//fmt.Println("lines: ", lines, " boardid: ", boardId)
	for i, line := range lines {
		//fmt.Println("i: ", i, ", line: ", line)
		strValues := strings.Fields(strings.TrimSpace(line))
		//fmt.Println("strvalues: ", strValues)
		for j := 0; j < 5; j++ {
			num, _ := strconv.Atoi(strValues[j])
			//fmt.Println("i: ", i, ", j: ", j, ", value: ", num, ", rawvalue: ", strValues[j])
			table[i][j] = num
		}
	}

	return &Board{
		table:   table,
		boardId: boardId,
	}
}

func (board Board) Inspect() {
	fmt.Println("boardId: ", board.boardId, ", has bingo?: ", board.HasBingo())
}

func (board Board) InspectFull() {
	fmt.Println("boardId: ", board.boardId)
	fmt.Println("rowHits: ", board.rowHits)
	fmt.Println("columnHits: ", board.columnHits)
	fmt.Println("table: ", board.table)
}

func (board *Board) HasBingo() bool {
	if board.bingoed {
		return true
	}

	for _, x := range board.rowHits {
		if x == 5 {
			board.bingoed = true
			return true
		}
	}

	for _, y := range board.columnHits {
		if y == 5 {
			board.bingoed = true
			return true
		}
	}

	return false
}

func (board *Board) ApplyCall(calledNumber int) {
	for i, row := range board.table {
		for j, x := range row {
			if x == calledNumber {
				//fmt.Println("==== Hit [", board.boardId, "] call: ", calledNumber, ", x: ", x)
				//fmt.Println("Before apply ---------------")
				//board.InspectFull()
				board.rowHits[i] += 1
				board.columnHits[j] += 1
				callEntry := []int{i, j, x}
				board.calls = append(board.calls, callEntry)
				board.table[i][j] = -1
				//fmt.Println("After apply ---------------")
				//board.InspectFull()
			}
		}
	}
}

func (board Board) CalculateScore(calledNumber int) int {
	var sum int = 0
	for _, row := range board.table {
		for _, x := range row {
			if x >= 0 {
				sum += x
			}
		}
	}

	//fmt.Println("sum: ", sum, ", calledNumber: ", calledNumber)
	return sum * calledNumber
}
