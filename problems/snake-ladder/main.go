package main

import (
	"github.com/lokesh-go/notes/problems/snake-ladder/game"
	"github.com/lokesh-go/notes/problems/snake-ladder/players"
)

func main() {
	// Initialise new game
	players := []players.Player{players.New("P1"), players.New("P2")}
	newGame := game.New(players, 10, 7, 5, 1)

	// Run
	newGame.Run()
}
