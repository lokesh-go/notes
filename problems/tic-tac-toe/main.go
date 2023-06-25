package main

import (
	"github.com/lokesh-go/notes/problems/tic-tac-toe/game"
	"github.com/lokesh-go/notes/problems/tic-tac-toe/players"
)

func main() {
	players := []players.Player{players.New("P1", "X"), players.New("P2", "O")}
	game := game.New(players, 3)
	game.Start()
}
