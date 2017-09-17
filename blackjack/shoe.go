package blackjack

import (
	"math/rand"
	"time"
)

// Shoe represents a shoe holding the cards that will be used during play.
type Shoe struct {
	cards []int
	i     int
}

// NewShoe returns a shuffled shoe containing decks number of 52-card decks.
func NewShoe(decks int) *Shoe {
	shoe := &Shoe{}

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	shoe.cards = rand.Perm(deckSize * decks)

	return shoe
}

// Draw returns the next card in the shoe.
func (s *Shoe) Draw() Card {
	c := deck[s.cards[s.i]%len(deck)]
	s.i++

	return c
}
