package blackjack

import (
	"fmt"

	"github.com/umesshk/CardDeck/deck"
)

type Hand []deck.Card

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
