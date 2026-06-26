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

func InitGame() {

	cards := deck.CreateDeck(deck.Deck(2), deck.Shuffle)

	var card deck.Card

	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = DrawCards(cards)
			*hand = append(*hand, card)
		}

	}

	fmt.Println("Player : ", player)
	fmt.Println("Dealer : ", dealer.DealerString())
}
