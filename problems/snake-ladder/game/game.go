package game

import (
	"fmt"
	"time"

	"github.com/lokesh-go/notes/problems/snake-ladder/board"
	"github.com/lokesh-go/notes/problems/snake-ladder/dice"
	"github.com/lokesh-go/notes/problems/snake-ladder/players"
)

type Game struct {
	board   *board.Board
	players []players.Player
	dice    *dice.Dice
}

// New ...
func New(players []players.Player, boardSize, noOfSnakes, noOfLadder, noOfDices int) *Game {
	// Initialise new game
	return &Game{
		board:   board.Init(boardSize, noOfSnakes, noOfLadder),
		players: players,
		dice:    dice.New(noOfDices),
	}
}

func (g *Game) Run() {
	gameClose := false
	for !gameClose {
		fmt.Println("\n$$$ NEW TURN $$$")
		for _, player := range g.players {
			winner := g.playTurn(player)
			if winner {
				gameClose = true
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (g *Game) playTurn(player players.Player) (winner bool) {
	fmt.Printf("\n@@@@@ Playing: %s\n", player.Name)

	// Roll dice
	result := g.dice.Roll()
	fmt.Printf("Your rolling result: %d\n", result)
	newPosition := player.Position + result

	// Condition for winner
	allCellsCount := len(g.board.Cell) * len(g.board.Cell)
	if newPosition == allCellsCount {
		fmt.Printf("******** Player: %s is WINNER ********", player.Name)
		return true
	}

	// Edge case
	if newPosition > allCellsCount {
		fmt.Println("<<<< You can't move, Wait for next turn! >>>>")
		return false
	}

	// Gets board position
	row := (newPosition - 1) / len(g.board.Cell)
	col := (newPosition - 1) % len(g.board.Cell)

	// Checks for ladder or snake
	if g.board.Cell[row][col] > 0 {
		// Checks
		if g.board.Cell[row][col] > newPosition {
			fmt.Println("You got Ladder >>>>>>>>>>>>>")
		} else {
			fmt.Println("You got Snake  <<<<<<<<<<<<<<<")
		}

		// Refresh position
		newPosition = g.board.Cell[row][col]
		row = (newPosition - 1) / len(g.board.Cell)
		col = (newPosition - 1) % len(g.board.Cell)
	}

	// Update position
	for ind, p := range g.players {
		if p.Name == player.Name {
			g.players[ind].Position = newPosition
		}
	}

	fmt.Printf("Your Position: %d\n", newPosition)
	fmt.Printf("Board Position - Row: %d-%d, Cell: %d\n", (row*len(g.board.Cell))+1, (row+1)*len(g.board.Cell), col+1)

	return false
}
