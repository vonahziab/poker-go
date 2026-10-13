package main

import (
	"flag"
	"fmt"
	"os"

	"poker/internal/poker"
)

func main() {
	balance := flag.Int("balance", 1000, "начальный баланс")
	bet := flag.Int("bet", 10, "ставка на раздачу")
	hands := flag.Int("hands", 10_000, "максимум раздач")
	flag.Parse()

	if *bet <= 0 || *balance < *bet || *hands <= 0 {
		fmt.Fprintln(os.Stderr, "нужно: bet > 0, balance >= bet, hands > 0")
		os.Exit(2)
	}

	res := poker.SimulateBankroll(*balance, *bet, *hands, poker.DefaultPaytable)

	printResult(res, *bet)
}

func printResult(res poker.BankrollResult, bet int) {
	fmt.Printf("Ставка:          %d\n", bet)
	fmt.Printf("Сыграно раздач:  %d\n", res.Hands)
	fmt.Printf("Баланс:          %d -> %d (%+d)\n", res.Start, res.Final, res.Net())
	fmt.Printf("Минимум/максимум: %d / %d\n", res.Min, res.Max)
	fmt.Printf("Возврат (RTP):   %.2f%%\n", res.RTP())

	if res.Busted {
		fmt.Println("Деньги закончились.")
	}

	fmt.Println()

	for rank := poker.RoyalFlush; rank >= poker.HighCard; rank-- {
		fmt.Printf("%-15s %9d  %8.4f%%\n", rank, res.Stats[rank], res.Stats.Percent(rank))
	}
}
