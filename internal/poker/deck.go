package poker

import "math/rand/v2"

func NewDeck() []Card {
	var deck []Card

	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Spades; suit++ {
			card := Card{
				Rank: rank,
				Suit: suit,
			}

			deck = append(deck, card)
		}
	}

	return deck
}

func ShuffleDeck(deck []Card) {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}
