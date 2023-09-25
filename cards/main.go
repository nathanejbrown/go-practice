package main

func main() {
	cards := newDeck()
	// cards.saveToFile("full_deck")
	cards.shuffle()
	evenOrOdd()
}