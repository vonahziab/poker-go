package poker

import "testing"

func TestSimulateBankrollAllLose(t *testing.T) {
	res := SimulateBankroll(1000, 10, 50, Paytable{})

	if res.Hands != 50 {
		t.Errorf("Hands = %d, want 50", res.Hands)
	}

	if res.Final != 500 {
		t.Errorf("Final = %d, want 500", res.Final)
	}

	if res.Min != 500 || res.Max != 1000 {
		t.Errorf("Min/Max = %d/%d, want 500/1000", res.Min, res.Max)
	}
}

func TestSimulateBankrollBusted(t *testing.T) {
	res := SimulateBankroll(100, 10, 1000, Paytable{})

	if res.Hands != 10 {
		t.Errorf("Hands = %d, want 10", res.Hands)
	}

	if res.Final != 0 || !res.Busted {
		t.Errorf("Final = %d, Busted = %v, want 0, true", res.Final, res.Busted)
	}
}

func TestSimulateBankrollAllWin(t *testing.T) {
	table := Paytable{}
	for rank := HighCard; rank <= RoyalFlush; rank++ {
		table[rank] = 2
	}

	res := SimulateBankroll(100, 10, 20, table)

	if res.Final != 300 {
		t.Errorf("Final = %d, want 300", res.Final)
	}

	if res.RTP() != 200 {
		t.Errorf("RTP() = %v, want 200", res.RTP())
	}

	if res.Net() != 200 {
		t.Errorf("Net() = %d, want 200", res.Net())
	}
}

func TestSimulateBankrollInvalidBet(t *testing.T) {
	res := SimulateBankroll(100, 0, 10, DefaultPaytable)

	if res.Hands != 0 || res.Final != 100 {
		t.Errorf("Hands = %d, Final = %d, want 0, 100", res.Hands, res.Final)
	}
}

func TestSimulateBankrollStatsMatchHands(t *testing.T) {
	res := SimulateBankroll(1_000_000, 1, 1000, DefaultPaytable)

	if got := res.Stats.Total(); got != res.Hands {
		t.Errorf("Stats.Total() = %d, Hands = %d", got, res.Hands)
	}
}
