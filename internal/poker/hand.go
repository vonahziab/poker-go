package poker

type HandRank int
type Hand []Card

const (
	HighCard HandRank = iota + 1
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var handRankNames = []string{
	"",
	"Старшая карта",
	"Пара",
	"Две пары",
	"Сет",
	"Стрит",
	"Флеш",
	"Фулл-хаус",
	"Каре",
	"Стрит-флеш",
	"Роял-флеш",
}

func (h HandRank) String() string {
	return handRankNames[h]
}

const rankSlots = int(Ace) + 1 // индексы 0..14, используем 2..14

// RankCounts хранит, сколько карт каждого ранга в руке (индекс = ранг).
type RankCounts [rankSlots]int

func (h Hand) CountRanks() RankCounts {
	var counts RankCounts

	for _, card := range h {
		counts[card.Rank]++
	}

	return counts
}

func (h Hand) IsFlush() bool {
	suit := h[0].Suit

	for _, card := range h[1:] {
		if card.Suit != suit {
			return false
		}
	}

	return true
}

func (counts RankCounts) IsStraight() bool {
	// Обычные последовательности: 2-3-4-5-6 ... 10-J-Q-K-A
	for start := 2; start <= 10; start++ {
		straight := true

		for rank := start; rank < start+5; rank++ {
			if counts[rank] != 1 {
				straight = false
				break
			}
		}

		if straight {
			return true
		}
	}

	// A-2-3-4-5
	return counts[Ace] == 1 &&
		counts[Two] == 1 &&
		counts[Three] == 1 &&
		counts[Four] == 1 &&
		counts[Five] == 1
}

// IsBroadway сообщает, что в руке есть 10-J-Q-K-A.
func (counts RankCounts) IsBroadway() bool {
	for rank := Ten; rank <= Ace; rank++ {
		if counts[rank] != 1 {
			return false
		}
	}

	return true
}

func (h Hand) Evaluate() HandRank {
	counts := h.CountRanks()

	pairs := 0
	three := 0
	four := 0

	for _, count := range counts {
		switch count {
		case 2:
			pairs++
		case 3:
			three++
		case 4:
			four++
		}
	}

	isFlush := h.IsFlush()
	isStraight := counts.IsStraight()

	switch {
	case isFlush && isStraight && counts.IsBroadway():
		return RoyalFlush
	case isFlush && isStraight:
		return StraightFlush
	case isFlush:
		return Flush
	case isStraight:
		return Straight
	case four == 1:
		return FourOfAKind
	case three == 1 && pairs == 1:
		return FullHouse
	case three == 1:
		return ThreeOfAKind
	case pairs == 2:
		return TwoPair
	case pairs == 1:
		return Pair
	default:
		return HighCard
	}
}
