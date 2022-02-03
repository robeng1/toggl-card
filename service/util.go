package service

var (
	suits  = [4]string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values = [13]string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
)

var (
	// internal data structures to hold cards to avoid
	// regenerating them on every call
	cardList []*Card
	cardDict map[string]Card
)

// It's is assumed all codes are valid card codes
func findCards(codes ...string) []*Card {
	allCards()
	initCards := make([]*Card, len(codes))
	for i, code := range codes {
		c := cardDict[code]
		initCards[i] = &c
	}
	return initCards
}

func allCards() []*Card {
	if len(cardList) > 0 {
		return copyCardList(cardList)
	}
	cardList = make([]*Card, 0)
	for _, suit := range suits {
		for _, value := range values {
			cardList = append(cardList, &Card{Value: value, Suit: suit, Code: string(value[0]) + string(suit[0])})
		}
	}
	cardDict = make(map[string]Card)
	for _, card := range cardList {
		cardDict[card.Code] = *card
	}
	return copyCardList(cardList)
}

func copyCardList(cards []*Card) []*Card {
	b := make([]*Card, len(cards))
	// return a copy to avoid modifying the original cards
	copy(b, cardList)
	return b
}
