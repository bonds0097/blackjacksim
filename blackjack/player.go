package blackjack

// Player represents a non-dealer participant in a game of blackjack.
type Player struct {
	Name string
	Hand *PlayerHand
}

func (p Player) String() string {
	return p.Name
}
