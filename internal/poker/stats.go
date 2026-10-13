package poker

// Stats хранит, сколько раз выпала каждая комбинация.
type Stats map[HandRank]int

// Total возвращает общее число раздач.
func (s Stats) Total() int {
	sum := 0

	for _, count := range s {
		sum += count
	}

	return sum
}

// Percent возвращает долю комбинации от всех раздач в процентах.
func (s Stats) Percent(rank HandRank) float64 {
	total := s.Total()
	if total == 0 {
		return 0
	}

	return float64(s[rank]) / float64(total) * 100
}
