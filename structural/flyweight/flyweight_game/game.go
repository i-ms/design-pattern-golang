package flyweight_game

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *Game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
