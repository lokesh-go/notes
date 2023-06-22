package board

import (
	"fmt"

	"github.com/lokesh-go/notes/problems/snake-ladder/ladder"
	"github.com/lokesh-go/notes/problems/snake-ladder/snake"
)

type Board struct {
	Cell [][]int
}

// Initialise board
func Init(size, snakes, ladders int) *Board {
	// Init board
	board := initBoard(size)

	// Add snakes
	board.addSnakes(snakes)

	// Add ladders
	board.addLadders(ladders)

	// Print board
	board.print()

	// Returns
	return board
}

func initBoard(size int) *Board {
	// Make board row
	cell := make([][]int, size)
	for i := range cell {
		// Add cells
		cell[i] = make([]int, size)
	}

	// Returns
	return &Board{cell}
}

func (b *Board) addSnakes(s int) {
	// Gets new snakes
	snakes := snake.New(len(b.Cell), s)

	for head, tail := range snakes {
		row := head / len(b.Cell)
		col := head % len(b.Cell)

		// Add snake
		b.Cell[row][col] = tail
	}
}

func (b *Board) addLadders(l int) {
	// Gets new ladders
	ladders := ladder.New(len(b.Cell), l)

	for start, end := range ladders {
		row := start / len(b.Cell)
		col := start % len(b.Cell)

		// Add ladders
		b.Cell[row][col] = end
	}
}

func (b *Board) print() {
	fmt.Println()
	fmt.Println("###############    BOARD     ###############")
	fmt.Println()
	for rowIndex, rowValue := range b.Cell {
		fmt.Printf("%d-%d. |", ((len(rowValue)-rowIndex-1)*len(rowValue))+1, (len(rowValue)-rowIndex)*len(rowValue))
		for colIndex := range rowValue {
			fmt.Printf(" %d |", b.Cell[len(rowValue)-rowIndex-1][colIndex])
		}
		fmt.Println()
	}
}
