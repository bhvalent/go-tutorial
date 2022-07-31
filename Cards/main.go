package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards = cards.shuffle()
	cards.print()
}
