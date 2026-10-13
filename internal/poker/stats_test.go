package poker

import "testing"

func TestStatsTotal(t *testing.T) {
	stats := Stats{Pair: 3, Flush: 2, HighCard: 5}

	if got := stats.Total(); got != 10 {
		t.Errorf("Total() = %d, want 10", got)
	}
}

func TestStatsPercent(t *testing.T) {
	stats := Stats{Pair: 1, HighCard: 3}

	tests := []struct {
		name string
		rank HandRank
		want float64
	}{
		{"есть в статистике", Pair, 25},
		{"большинство", HighCard, 75},
		{"нет в статистике", Flush, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stats.Percent(tt.rank); got != tt.want {
				t.Errorf("Percent(%v) = %v, want %v", tt.rank, got, tt.want)
			}
		})
	}
}

func TestStatsPercentEmpty(t *testing.T) {
	if got := (Stats{}).Percent(Pair); got != 0 {
		t.Errorf("Percent() на пустой статистике = %v, want 0", got)
	}
}
