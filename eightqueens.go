package piscine

import "github.com/01-edu/z01"

const n = 8

func EightQueens() {
	board := []int{}
	res := [][]int{}
	var visited [n + 1]bool

	nQueenUtil(1, board, &res, visited)

	for _, resBoard := range res {
		printBoard(resBoard)
	}
}

func nQueenUtil(col int, board []int, res *[][]int, visited [n + 1]bool) {
	// If all queens got safely placed, add board to res
	if col > n {
		*res = append(*res, board[:])
		return
	}

	for row := 1; row <= n; row++ {
		// If row aleady visited, move to next iteration
		if visited[row] {
			continue
		}

		// Ensure queen is safe before proceeding
		if !isQueenSafe(col, row, board) {
			continue
		}

		// Mark row as visited
		visited[row] = true

		// Place queen
		board = append(board, row)

		// Recur for next col
		nQueenUtil(col+1, board, res, visited)

		// Backtrack
		board = board[:len(board)-1] // Remove last item
		visited[row] = false
	}
}

// Helper func to check if placement is safe
func isQueenSafe(curCol int, curRow int, board []int) bool {
	for i, placedRow := range board {
		placedCol := i + 1

		// Check for diagonals
		if absInt(placedCol-curCol) == absInt(placedRow-curRow) {
			return false
		}
	}

	return true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(board []int) {
	for _, num := range board {
		z01.PrintRune(rune('0' + num))
	}
	z01.PrintRune('\n')
}
