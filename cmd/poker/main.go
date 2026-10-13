package main

import (
	"flag"
	"fmt"
	"time"

	"poker/internal/poker"
)

func main() {
	n := flag.Int("n", 1_000_000, "количество симуляций")
	flag.Parse()

	start := time.Now()
	sequential := poker.Simulate(*n)
	sequentialTime := time.Since(start)

	start = time.Now()
	concurrent := poker.SimulateConcurrent(*n)
	concurrentTime := time.Since(start)

	fmt.Println("Sequential:", sequentialTime)
	fmt.Println("Concurrent:", concurrentTime)

	fmt.Println("Sequential total:", sequential.Total())
	fmt.Println("Concurrent total:", concurrent.Total())

	printStats(concurrent)
}

// printStats выводит комбинации от сильной к слабой с долей от общего числа раздач.
func printStats(stats poker.Stats) {
	for rank := poker.RoyalFlush; rank >= poker.HighCard; rank-- {
		fmt.Printf("%-15s %9d  %8.4f%%\n", rank, stats[rank], stats.Percent(rank))
	}
}
