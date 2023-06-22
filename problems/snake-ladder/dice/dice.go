package dice

import "math/rand"

// Dice
type Dice struct {
	NoOfDices int
}

// New
func New(count int) *Dice {
	return &Dice{count}
}

// Roll
func (d *Dice) Roll() (result int) {
	result = 0
	min := 1
	max := 6
	for i := 0; i < d.NoOfDices; i++ {
		result = result + rand.Intn(max-min) + min
	}
	return result
}
