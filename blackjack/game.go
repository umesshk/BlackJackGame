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

func InitGame() {

	cards := deck.CreateDeck(deck.Deck(2), deck.Shuffle)

	var card deck.Card

	var p1, p2 Hand

	for i := 0; i < 2; i++ {
		card, cards = cards[0], cards[1:]
		p1 = append(p1, card)

		card, cards = cards[0], cards[1:]
		p2 = append(p2, card)

	}

	fmt.Println(p1)
	fmt.Println(p2)

}
