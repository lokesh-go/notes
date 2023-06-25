package game

import (
	"fmt"

	"github.com/lokesh-go/notes/problems/tic-tac-toe/board"
	"github.com/lokesh-go/notes/problems/tic-tac-toe/players"
)

type Game struct {
	players []players.Player
	board   *board.Board
}

// New ...
func New(players []players.Player, boardSize int) *Game {
	return &Game{
		players: players,
		board:   board.Init(boardSize),
	}
}

// Start ...
func (g *Game) Start() {
	gameFinish := false
	for !gameFinish {
		for _, player := range g.players {
			finish := g.playTurn(player)
			if finish {
				gameFinish = true
				break
			}
		}
	}
}

func (g *Game) playTurn(player players.Player) (finish bool) {
	// Check free cells
	if !g.board.IsFreeCells() {
		fmt.Println("!!!!! TIE !!!!!")
		return true
	}

	// Take input from user
	var cellNo int
	fmt.Printf("\n%s -> input cell no: ", player.Name)
	fmt.Scan(&cellNo)
	piecePlayed := player.PlayingPiece.Play()

	// Add piece on board
	added := g.board.AddPiece(cellNo, piecePlayed)
	if !added {
		fmt.Println("Invalid position try in next turn!")
	}

	// Check winner
	win := g.board.CheckWinner(cellNo, piecePlayed)
	if win {
		fmt.Printf("***** %s is Winner *****", player.Name)
		return true
	}

	// Return
	return false
}
