package main

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []rune) {
	if len(deck) != 12 {
		z01.Println("The deck should contain exactly 12 cards.")
		return
	}

	players := 4
	cardsPerPlayer := len(deck) / players

	for i := 0; i < players; i++ {
		start := i * cardsPerPlayer
		end := (i + 1) * cardsPerPlayer

		playerCards := deck[start:end]

		z01.Printf("Player %d: ", i+1)
		for _, card := range playerCards {
			z01.PrintRune(card)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	// Sample deck of 12 cards represented as runes
	deck := []rune{'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q'}

	DealAPackOfCards(deck)
}
