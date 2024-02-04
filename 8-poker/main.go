package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	cardCount = 2
)

var values = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var suits = []string{"S", "C", "D", "H"}

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		process(in)
	}
}

func process(in *bufio.Reader) {
	var total int
	fmt.Fscan(in, &total)

	deck := newDeck()
	hands := make([][]string, total)

	for player := 0; player < total; player++ {
		hands[player] = make([]string, 3)
		for i := 0; i < cardCount; i++ {
			var card string
			fmt.Fscan(in, &card)
			delete(deck, card)
			hands[player][i] = card
		}
	}

	results := []string{}
	for card := range deck {
		winning := true
		hands[0][cardCount] = card
		playerValue := handValue(hands[0])
		for player := 1; player < total; player++ {
			hands[player][cardCount] = card
			if handValue(hands[player]) > playerValue {
				winning = false
				break
			}
		}
		if winning {
			results = append(results, card)
		}
	}

	fmt.Println(len(results))
	for _, card := range results {
		fmt.Println(card)
	}
}

func newDeck() map[string]bool {
	d := make(map[string]bool, 52)
	for k := range values {
		for _, suit := range suits {
			d[k+suit] = true
		}
	}
	return d
}

func handValue(hand []string) int {
	total := 0
	counts := map[string]int{}
	for _, card := range hand {
		counts[card[:1]]++
	}
	for name, count := range counts {
		mult := 1
		for i := 1; i < count; i++ {
			mult *= 10
		}
		val := values[name] * mult
		if val > total {
			total = val
		}
	}
	return total
}
