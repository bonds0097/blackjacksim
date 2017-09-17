package blackjack

import (
	"fmt"
)

// Game represents the internal state of a blackjack game.
type Game struct {
	Shoe       *Shoe
	Players    []*Player
	DealerHand DealerHand
}

func (g Game) String() string {
	s := fmt.Sprintf("Dealer - %v\n", g.DealerHand)
	for _, p := range g.Players {
		s += fmt.Sprintf("%s - %v\n", p.Name, p.Hand)
	}

	return s
}

// Deal returns initial hands for the dealer and assigns each player a hand.
func (g *Game) Deal() {
	for i := 0; i < 2; i++ {
		for _, p := range g.Players {
			if i == 0 {
				p.Hand = &PlayerHand{}
				p.Hand.FirstCard = g.Shoe.Draw()
			}

			p.Hand.SecondCard = g.Shoe.Draw()
		}
	}

	g.DealerHand = DealerHand{g.Shoe.Draw(), g.Shoe.Draw()}
}
