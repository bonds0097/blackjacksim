package blackjack

import (
	"fmt"
)

// DealerHand represents the two cards dealt to the dealer at the start of a
// Blackjack game.
type DealerHand struct {
	FirstCard Card
	HoleCard  Card
}

func (h DealerHand) String() string {
	return fmt.Sprintf("%s (%s)", h.FirstCard, h.HoleCard)
}

// PlayerHand represents a hand of cards that belongs to a non-dealer player.
type PlayerHand struct {
	FirstCard       Card
	SecondCard      Card
	AdditionalCards []Card
	CanSplit        bool
	CanSurrender    bool
}

func (h PlayerHand) String() string {
	s := fmt.Sprintf("%v, %v", h.FirstCard, h.SecondCard)

	for _, c := range h.AdditionalCards {
		s += fmt.Sprintf(", %v", c)
	}

	return s
}
