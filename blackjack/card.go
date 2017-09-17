package blackjack

// Card represents a card in a blackjack game.
type Card struct {
	Name   string
	Value1 int
	Value2 int
}

func (c Card) String() string {
	return c.Name
}
