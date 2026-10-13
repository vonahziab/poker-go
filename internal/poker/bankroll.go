package poker

// Paytable задаёт, во сколько ставок возвращается выплата за комбинацию.
// 0 — ставка проиграна, 1 — ставка возвращена, 3 — выигрыш 2 к 1 и т. д.
type Paytable map[HandRank]int

// DefaultPaytable — таблица выплат по умолчанию.
var DefaultPaytable = Paytable{
	HighCard:      0,
	Pair:          1,
	TwoPair:       3,
	ThreeOfAKind:  4,
	Straight:      6,
	Flush:         8,
	FullHouse:     12,
	FourOfAKind:   40,
	StraightFlush: 100,
	RoyalFlush:    500,
}

// BankrollResult — итог игры на дистанции.
type BankrollResult struct {
	Start    int   // начальный баланс
	Final    int   // конечный баланс
	Min      int   // самый низкий баланс за игру
	Max      int   // самый высокий баланс за игру
	Hands    int   // сыграно раздач
	Wagered  int   // сумма всех ставок
	Returned int   // сумма всех выплат
	Busted   bool  // баланса не хватает на следующую ставку
	Stats    Stats // сколько раз выпала каждая комбинация
}

// Net возвращает итоговый выигрыш (отрицательный — проигрыш).
func (r BankrollResult) Net() int {
	return r.Final - r.Start
}

// RTP возвращает долю выплат от ставок в процентах.
func (r BankrollResult) RTP() float64 {
	if r.Wagered == 0 {
		return 0
	}

	return float64(r.Returned) / float64(r.Wagered) * 100
}

// SimulateBankroll играет раздачи по ставке bet, пока не сыграно hands
// раздач или не закончились деньги.
func SimulateBankroll(balance, bet, hands int, table Paytable) BankrollResult {
	res := BankrollResult{
		Start: balance,
		Min:   balance,
		Max:   balance,
		Stats: make(Stats),
	}

	if bet > 0 {
		deck := NewDeck()

		for res.Hands < hands && balance >= bet {
			ShuffleDeck(deck)

			rank := Hand(deck[:5]).Evaluate()
			payout := bet * table[rank]

			balance += payout - bet

			res.Hands++
			res.Wagered += bet
			res.Returned += payout
			res.Stats[rank]++
			res.Min = min(res.Min, balance)
			res.Max = max(res.Max, balance)
		}
	}

	res.Final = balance
	res.Busted = balance < bet

	return res
}
