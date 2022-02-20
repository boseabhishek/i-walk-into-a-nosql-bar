package domain

// Game ...
type Game struct {
	Title    string
	TopScore int // TODO: think about allocation later?
}

// Gamer ...
type Gamer struct {
	ID     string
	Played []Game
}

// GamerService ...
type GamerService interface {
	FindGameAndScore(ID string) ([]Game, error)
}
