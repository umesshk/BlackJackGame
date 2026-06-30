package blackjack

import (
	"fmt"
	"strings"

	"github.com/umesshk/CardDeck/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	str := make([]string, len(h))

	for i := range h {
		str[i] = h[i].String()
	}

	return strings.Join(str, ", ")

}

func (h Hand) DealerString() string {
	return h[0].String() + ",**HIDDEN**"
}

func DrawCards(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]

}

func (h Hand) Score() int {

	minScore := h.MinScore()

	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore

}

func (h Hand) MinScore() int {

	score := 0

	for _, c := range h {
		score += min(int(c.Rank), 10)
	}

	return score

}

func Shuffle(gs GameSate) GameSate {
	ret := Clone(gs)
	ret.Deck = deck.CreateDeck(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs GameSate) GameSate {
	ret := Clone(gs)

	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)

	var card deck.Card

	for range 2 {
		card, ret.Deck = DrawCards(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = DrawCards(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}

	ret.State = int(StatePlayerTurn)

	return ret

}

func Stand(gs GameSate) GameSate {
	ret := Clone(gs)
	ret.State++

	return ret
}

func Hit(gs GameSate) GameSate {
	ret := Clone(gs)
	hand := ret.CurrentPlayer()

	var card deck.Card

	card, ret.Deck = DrawCards(ret.Deck)

	*hand = append(*hand, card)

	if hand.Score() > 21 {
		return Stand(ret)
	}

	return ret
}

func EndGame(gs GameSate) GameSate {

	ret := Clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()

	fmt.Println("Player : ", ret.Player)
	fmt.Println("Player Score : ", pScore)
	fmt.Println("Dealer : ", ret.Dealer)
	fmt.Println("Dealer : ", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You Busted")

	case dScore > 21:
		fmt.Println("Dealer Busted")

	case pScore > dScore:
		fmt.Println("You Win!!")

	case dScore > pScore:
		fmt.Println("You Lose!!")

	default:
		fmt.Println("Draw")
	}

	fmt.Println()

	return ret
}

func InitGame() {

	var gs GameSate

	gs = Shuffle(gs)

	for range 10 {
		gs = Deal(gs)

		var input string

		for gs.State == int(StatePlayerTurn) {

			fmt.Println("Dealer : ", gs.Dealer.DealerString())
			fmt.Println("Player : ", gs.Player)

			fmt.Println("What would you like to do (h)it or (s)tand ? ")

			fmt.Scanf("%s\n", &input)

			switch input {
			case "h":
				gs = Hit(gs)

			case "s":
				gs = Stand(gs)

			default:
				fmt.Println("Invalid Case")
			}

		}

		gs = EndGame(gs)
	}
}
