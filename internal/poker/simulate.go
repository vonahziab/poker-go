package poker

import "runtime"

func Simulate(n int) Stats {
	stats := make(Stats)
	deck := NewDeck()

	for range n {
		ShuffleDeck(deck)

		combination := Hand(deck[:5]).Evaluate()

		stats[combination]++
	}

	return stats
}

func SimulateConcurrent(n int) Stats {
	workers := runtime.NumCPU()
	results := make(chan Stats, workers)

	for i := range workers {
		count := n / workers

		if i < n%workers {
			count++
		}

		go func(count int) {
			results <- Simulate(count)
		}(count)
	}

	stats := make(Stats)

	for range workers {
		workerStats := <-results

		for combination, count := range workerStats {
			stats[combination] += count
		}
	}

	return stats
}
