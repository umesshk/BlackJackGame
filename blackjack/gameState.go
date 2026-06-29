package blackjack

import (
	"github.com/umesshk/CardDeck/deck"
)

type GameSate struct {
	Deck   []deck.Card
	State  int
	Player Hand
	Dealer Hand
}

func Clone(gs GameSate) GameSate {
	ret := GameSate{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make([]deck.Card, len(gs.Player)),
		Dealer: make([]deck.Card, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)

	return ret

}
