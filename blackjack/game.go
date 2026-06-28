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

	var input string

	for input != "s" {
		fmt.Println("Dealer : ", dealer.DealerString())
		fmt.Println("Player : ", player)

		fmt.Println("What would you like to do (h)it or (s)tand ? ")

		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, cards = DrawCards(cards)
			player = append(player, card)

		}

	}

	pScore, dScore := player.Score(), dealer.Score()

	fmt.Println("Player : ", player)
	fmt.Println("Player Score : ", pScore)
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

}
