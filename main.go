package main

import "github.com/bonds0097/blackjacksim/blackjack"
import "fmt"

func main() {
	g := &blackjack.Game{
		Shoe: blackjack.NewShoe(6),
		Players: []*blackjack.Player{
			&blackjack.Player{"Alfredo", nil},
			&blackjack.Player{"Chance", nil},
		},
	}

	g.Deal()

	fmt.Print(g)
}
