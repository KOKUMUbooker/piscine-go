package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	playerCount := 4
	cardsPerPlayer := 3

	for i := 0; i < playerCount; i++ {
		PrintStr("Player ")
		PrintInt(i + 1)
		PrintStr(": ")

		for j := 0; j < cardsPerPlayer; j++ {
			cardIndex := i*cardsPerPlayer + j
			PrintInt(deck[cardIndex])
			if j != cardsPerPlayer-1 {
				PrintStr(", ")
			}
		}

		z01.PrintRune('\n')
	}
}

func PrintStr2(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		d := n % 10
		digits = append([]rune{rune(d + '0')}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
