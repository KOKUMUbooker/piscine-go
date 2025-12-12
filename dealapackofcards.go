package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	playerCount := 4
	cardsPerPlayer := len(deck) / playerCount

	for i := 0; i < playerCount; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			cardIndex := i*cardsPerPlayer + j
			fmt.Print(deck[cardIndex])
			if j != cardsPerPlayer-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
