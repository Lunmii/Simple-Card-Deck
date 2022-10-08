package main

import "fmt"

func main() {
	cards := []string{"Ace Of Spade", randomCard()}
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func randomCard() string {
	return "Black Jack"
}
