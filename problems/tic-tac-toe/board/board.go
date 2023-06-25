package board

import "fmt"

type Board struct {
	Cell [][]string
}

// Init ...
func Init(size int) *Board {
	// Init board
	board := initBoard(size)

	// Print
	board.print()

	// Returns
	return board
}

func initBoard(size int) *Board {
	// Make board row
	cell := make([][]string, size)
	for i := range cell {
		// Add cells
		cell[i] = make([]string, size)
	}

	// Returns
	return &Board{cell}
}

func (b *Board) AddPiece(cellNo int, pieceType string) bool {
	// Gets board position
	row := (cellNo - 1) / len(b.Cell)
	col := (cellNo - 1) % len(b.Cell)

	if b.Cell[row][col] != "" {
		return false
	}
	b.Cell[row][col] = pieceType
	b.print()
	return true
}

func (b *Board) CheckWinner(cellNo int, pieceType string) bool {
	// Gets board position
	row := (cellNo - 1) / len(b.Cell)
	col := (cellNo - 1) % len(b.Cell)

	rowMatch := true
	colMatch := true
	diagonalMatch := true
	antiDiagonalMatch := true

	// Check row
	for i := 0; i < len(b.Cell); i++ {
		if b.Cell[row][i] == "" || b.Cell[row][i] != pieceType {
			rowMatch = false
		}
	}

	// Check col
	for i := 0; i < len(b.Cell); i++ {
		if b.Cell[i][col] == "" || b.Cell[i][col] != pieceType {
			colMatch = false
		}
	}

	// Check diagonal
	i := 0
	j := 0
	for i < len(b.Cell) {
		if b.Cell[i][j] == "" || b.Cell[i][j] != pieceType {
			diagonalMatch = false
		}

		i = i + 1
		j = j + 1
	}

	// Check anti diagonal
	i = 0
	j = len(b.Cell) - 1
	for i < len(b.Cell) {
		if b.Cell[i][j] == "" || b.Cell[i][j] != pieceType {
			antiDiagonalMatch = false
		}

		i++
		j--
	}

	return rowMatch || colMatch || diagonalMatch || antiDiagonalMatch
}

func (b *Board) IsFreeCells() bool {
	for i := 0; i < len(b.Cell); i++ {
		for j := 0; j < len(b.Cell); j++ {
			if b.Cell[i][j] == "" {
				return true
			}
		}
	}
	return false
}

func (b *Board) print() {
	fmt.Println()
	for rowIndex, rowData := range b.Cell {
		fmt.Printf("%d-%d -> |", (len(rowData)*(len(rowData)-rowIndex-1))+1, len(rowData)*(len(rowData)-rowIndex))
		for colIndex := range rowData {
			val := b.Cell[len(rowData)-rowIndex-1][colIndex]
			if val == "" {
				fmt.Printf("  %s |", val)
			} else {
				fmt.Printf(" %s |", val)
			}
		}
		fmt.Println()
	}
}
