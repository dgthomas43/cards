package main

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()
	hand := deck{}
	cards.shuffle()

	hand, cards = deal(cards, 5)
	hand.print()
	cards.print()
	//hand.saveToFile("my_hand")
}
