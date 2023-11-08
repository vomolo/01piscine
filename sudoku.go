package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	if checkErrors(arr) {
		fmt.Println("Error: Invalid Sudoku input")
	} else {
		sudoku, err := parseSudoku(arr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			if solveSudoku(sudoku) {
				printSudoku(sudoku)
			} else {
				fmt.Println("Error: No solution found")
			}
		}
	}
}

func checkErrors(arr []string) bool {
	if len(arr) != 9 {
		return true
	}
	for _, str := range arr {
		if len(str) != 9 {
			return true
		}
		for _, ch := range str {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
	}
	return false
}

func parseSudoku(arr []string) ([][]int, error) {
	sudoku := make([][]int, 9)
	for i := 0; i < 9; i++ {
		sudoku[i] = make([]int, 9)
		for j, ch := range arr[i] {
			if ch == '.' {
				sudoku[i][j] = 0
			} else {
				num := int(ch - '0')
				if num < 1 || num > 9 {
					return nil, fmt.Errorf("Invalid character '%c' in Sudoku input", ch)
				}
				sudoku[i][j] = num
			}
		}
	}
	return sudoku, nil
}

func solveSudoku(sudoku [][]int) bool {
	empty, row, col := findEmptyLocation(sudoku)
	if !empty {
		return true // All cells are filled
	}

	for num := 1; num <= 9; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num

			if solveSudoku(sudoku) {
				return true
			}

			sudoku[row][col] = 0 // Reset if the current configuration doesn't lead to a solution
		}
	}

	return false
}

func findEmptyLocation(sudoku [][]int) (bool, int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if sudoku[row][col] == 0 {
				return true, row, col
			}
		}
	}
	return false, -1, -1
}

func isSafe(sudoku [][]int, row, col, num int) bool {
	return !usedInRow(sudoku, row, num) && !usedInCol(sudoku, col, num) && !usedInBox(sudoku, row-row%3, col-col%3, num)
}

func usedInRow(sudoku [][]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if sudoku[row][col] == num {
			return true
		}
	}
	return false
}

func usedInCol(sudoku [][]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if sudoku[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(sudoku [][]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if sudoku[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

func printSudoku(sudoku [][]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print(sudoku[row][col], " ")
		}
		fmt.Println()
	}
}
