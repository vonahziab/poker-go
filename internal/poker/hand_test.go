package poker

import "testing"

func TestRankCountsIsStraight(t *testing.T) {
	tests := []struct {
		name  string
		ranks []Rank
		want  bool
	}{
		{"обычный стрит", []Rank{Five, Six, Seven, Eight, Nine}, true},
		{"стрит с тузом", []Rank{Ace, Two, Three, Four, Five}, true},
		{"нет стрита", []Rank{Two, Three, Four, Five, Seven}, false},
		{"пара", []Rank{Two, Two, Three, Four, Five}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var counts RankCounts
			for _, rank := range tt.ranks {
				counts[rank]++
			}

			if got := counts.IsStraight(); got != tt.want {
				t.Errorf("IsStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluateHand(t *testing.T) {
	tests := []struct {
		name string
		hand []Card
		want HandRank
	}{
		{"старшая карта", []Card{
			{Rank: Two, Suit: Hearts}, {Rank: Five, Suit: Spades}, {Rank: Nine, Suit: Clubs},
			{Rank: Jack, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, HighCard},
		{"пара", []Card{
			{Rank: Two, Suit: Hearts}, {Rank: Two, Suit: Spades}, {Rank: Five, Suit: Clubs},
			{Rank: Nine, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, Pair},
		{"две пары", []Card{
			{Rank: Two, Suit: Hearts}, {Rank: Two, Suit: Spades}, {Rank: Five, Suit: Clubs},
			{Rank: Five, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, TwoPair},
		{"сет", []Card{
			{Rank: Seven, Suit: Hearts}, {Rank: Seven, Suit: Spades}, {Rank: Seven, Suit: Clubs},
			{Rank: Nine, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, ThreeOfAKind},
		{"стрит", []Card{
			{Rank: Five, Suit: Hearts}, {Rank: Six, Suit: Spades}, {Rank: Seven, Suit: Clubs},
			{Rank: Eight, Suit: Diamonds}, {Rank: Nine, Suit: Hearts},
		}, Straight},
		{"стрит A-2-3-4-5", []Card{
			{Rank: Ace, Suit: Hearts}, {Rank: Two, Suit: Spades}, {Rank: Three, Suit: Clubs},
			{Rank: Four, Suit: Diamonds}, {Rank: Five, Suit: Hearts},
		}, Straight},
		{"флеш", []Card{
			{Rank: Two, Suit: Clubs}, {Rank: Five, Suit: Clubs}, {Rank: Nine, Suit: Clubs},
			{Rank: Jack, Suit: Clubs}, {Rank: King, Suit: Clubs},
		}, Flush},
		{"фулл-хаус", []Card{
			{Rank: Three, Suit: Hearts}, {Rank: Three, Suit: Spades}, {Rank: Three, Suit: Clubs},
			{Rank: King, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, FullHouse},
		{"каре", []Card{
			{Rank: Nine, Suit: Hearts}, {Rank: Nine, Suit: Spades}, {Rank: Nine, Suit: Clubs},
			{Rank: Nine, Suit: Diamonds}, {Rank: King, Suit: Hearts},
		}, FourOfAKind},
		{"стрит-флеш", []Card{
			{Rank: Five, Suit: Spades}, {Rank: Six, Suit: Spades}, {Rank: Seven, Suit: Spades},
			{Rank: Eight, Suit: Spades}, {Rank: Nine, Suit: Spades},
		}, StraightFlush},
		{"стрит-флеш с тузом внизу", []Card{
			{Rank: Ace, Suit: Diamonds}, {Rank: Two, Suit: Diamonds}, {Rank: Three, Suit: Diamonds},
			{Rank: Four, Suit: Diamonds}, {Rank: Five, Suit: Diamonds},
		}, StraightFlush},
		{"роял-флеш", []Card{
			{Rank: Ten, Suit: Hearts}, {Rank: Jack, Suit: Hearts}, {Rank: Queen, Suit: Hearts},
			{Rank: King, Suit: Hearts}, {Rank: Ace, Suit: Hearts},
		}, RoyalFlush},
		{"флеш-стрит до короля не роял", []Card{
			{Rank: Nine, Suit: Hearts}, {Rank: Ten, Suit: Hearts}, {Rank: Jack, Suit: Hearts},
			{Rank: Queen, Suit: Hearts}, {Rank: King, Suit: Hearts},
		}, StraightFlush},
		{"стрит от десятки не флеш, не роял", []Card{
			{Rank: Ten, Suit: Hearts}, {Rank: Jack, Suit: Spades}, {Rank: Queen, Suit: Hearts},
			{Rank: King, Suit: Hearts}, {Rank: Ace, Suit: Hearts},
		}, Straight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hand(tt.hand).Evaluate(); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandRankString(t *testing.T) {
	tests := []struct {
		rank HandRank
		want string
	}{
		{HighCard, "Старшая карта"},
		{Pair, "Пара"},
		{TwoPair, "Две пары"},
		{ThreeOfAKind, "Сет"},
		{Straight, "Стрит"},
		{Flush, "Флеш"},
		{FullHouse, "Фулл-хаус"},
		{FourOfAKind, "Каре"},
		{StraightFlush, "Стрит-флеш"},
		{RoyalFlush, "Роял-флеш"},
	}

	for _, tt := range tests {
		if got := tt.rank.String(); got != tt.want {
			t.Errorf("HandRank(%d).String() = %q, want %q", tt.rank, got, tt.want)
		}
	}
}
