package snake

import (
	"math/rand"
	"time"
)

// New
func New(boardSize, s int) (snakes map[int]int) {
	// Initialize
	rand.Seed(time.Now().UnixNano())
	snakes = make(map[int]int)
	var head int
	var tail int
	min := 1
	max := boardSize * boardSize

	for i := 0; i < s; i++ {
		for {
			head = rand.Intn(max-min) + min
			tail = rand.Intn(max-min) + min
			_, exists := snakes[head]
			if head > tail && !exists {
				break
			}
		}
		snakes[head] = tail
	}

	// Returns
	return snakes
}
