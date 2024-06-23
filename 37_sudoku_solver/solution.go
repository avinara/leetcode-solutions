package main

import "strconv"

func solveSudoku(board [][]byte) {
	length := len(board[0])
	boxes := make([]map[string]bool, 9)
	rows := make([]map[string]bool, 9)
	cols := make([]map[string]bool, 9)

	for i := 0; i < length; i++ {
		boxes[i] = make(map[string]bool)
		rows[i] = make(map[string]bool)
		cols[i] = make(map[string]bool)
	}

	for r := 0; r < length; r++ {
		for c := 0; c < length; c++ {
			if string(board[r][c]) != "." {
				val := string(board[r][c])
				boxId := getBoxId(r, c)
				boxes[boxId][val] = true
				rows[r][val] = true
				cols[c][val] = true
			}
		}
	}
	solveBacktrack(board, boxes, rows, cols, 0, 0)
}

func solveBacktrack(board [][]byte, boxes, rows, cols []map[string]bool, r int, c int) bool {
	period := "."
	if r == len(board[0]) || c == len(board[0]) {
		return true
	}
	if string(board[r][c]) == period {
		for num := 1; num <= 9; num++ {
			numString := strconv.Itoa(num)
			board[r][c] = numString[0]
			boxId := getBoxId(r, c)
			box := boxes[boxId]
			row := rows[r]
			col := cols[c]

			if isValid(box, row, col, numString) {
				box[numString] = true
				row[numString] = true
				col[numString] = true

				if c == len(board[0])-1 {
					if solveBacktrack(board, boxes, rows, cols, r+1, 0) {
						return true
					}
				} else {
					if solveBacktrack(board, boxes, rows, cols, r, c+1) {
						return true
					}
				}
				delete(box, numString)
				delete(row, numString)
				delete(col, numString)
			}
			board[r][c] = period[0]
		}
	} else {
		if c == len(board[0])-1 {
			if solveBacktrack(board, boxes, rows, cols, r+1, 0) {
				return true
			}
		} else {
			if solveBacktrack(board, boxes, rows, cols, r, c+1) {
				return true
			}
		}
	}
	return false
}

func isValid(box, row, col map[string]bool, num string) bool {
	if box[num] || row[num] || col[num] {
		return false
	}
	return true
}

func getBoxId(r int, c int) int {
	colVal := c / 3
	rowVal := (r / 3) * 3
	return rowVal + colVal
}
