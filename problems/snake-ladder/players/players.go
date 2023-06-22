package players

// Player ...
type Player struct {
	Name     string
	Position int
}

// New ...
func New(name string) Player {
	return Player{
		Name:     name,
		Position: 0,
	}
}
