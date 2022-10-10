package main

func main() {
	cards := deck{"Ace Of Spade", randomCard()}
	cards = append(cards, "Six of spades")

	cards.print()
}

func randomCard() string {
	return "Black Jack"
}
