package main

func main() {
	/*var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = newCard()
	fmt.Println(card) */

	/*
		cards := []string{newCard(), "Ace of Diamonds"}
		cards = append(cards, "Six of Spades")
		for i, card := range cards {
			fmt.Println(i, card)
		}*/
	/*
		cards := newDeck()
		hand, remainingCard := deal(cards, 5)

		hand.print()
		remainingCard.print()
	*/

	/*
		cards := newDeck()
		fmt.Println(cards.toString())
	*/

	/*
		cards := newDeck()
		cards.saveToFile("mycards")
	*/

	/*
		cards := newDeckFromFile("mycards")
		cards.print()
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*
func newCard() string {
	return "Five of Diamonds"
}*/
