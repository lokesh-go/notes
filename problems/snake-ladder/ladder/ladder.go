package ladder

import (
	"math/rand"
)

// New
func New(boardSize, l int) (ladders map[int]int) {
	// Initialize
	ladders = make(map[int]int)
	var start int
	var end int
	min := 1
	max := boardSize * boardSize

	for i := 0; i < l; i++ {
		for {
			start = rand.Intn(max-min) + min
			end = rand.Intn(max-min) + min
			_, exists := ladders[start]
			if start < end && !exists {
				break
			}
		}
		ladders[start] = end
	}

	// Returns
	return ladders
}
