package game

type Game struct {
	maxX int
	maxY int
}

func NewGame() *Game {
	return &Game{
		maxX: 100,
		maxY: 100,
	}
}
