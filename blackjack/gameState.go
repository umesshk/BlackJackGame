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

type State int

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

func (gs *GameSate) CurrentPlayer() *Hand {
	switch gs.State {
	case int(StatePlayerTurn):
		return &gs.Player
	case int(StateDealerTurn):
		return &gs.Dealer
	default:
		panic("Not a Valid State")
	}
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
