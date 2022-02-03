package service

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	Shuffled bool      `json:"shuffled,omitempty"`
	Cards    []*Card   `json:"cards,omitempty"`
	DeckID   uuid.UUID `json:"deck_id,omitempty"`
}

func (deck Deck) Remaining() int {
	return len(deck.Cards)
}

// Shuffle : shuffle the cards in the deck 4 times
func (deck *Deck) Shuffle() {
	if deck != nil {
		rand.Seed(time.Now().UnixNano())
		for s := 1; s <= 5; s++ {
			rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
		}
		deck.Shuffled = true
	}
}

func (deck *Deck) WithoutCards() *Deckdto {
	if deck != nil {
		return &Deckdto{
			DeckID:    deck.DeckID,
			Shuffled:  deck.Shuffled,
			Remaining: deck.Remaining(),
		}
	}
	return nil
}

// MarshalJSON is implementation of json.Marshaller
func (deck Deck) MarshalJSON() ([]byte, error) {
	return marshalJSON(deck)
}

// custom implementation of the json.Marshaller to add
// remaining to the fields
func marshalJSON(deck Deck) ([]byte, error) {
	data := make(map[string]interface{})
	data["shuffled"] = deck.Shuffled
	data["remaining"] = deck.Remaining()
	data["cards"] = deck.Cards
	data["deck_id"] = deck.DeckID
	return json.Marshal(&data)
}

// A data transfer struct that omits the cards of the deck
type Deckdto struct {
	Shuffled  bool      `json:"shuffled,omitempty"`
	Remaining int       `json:"remaining,omitempty"`
	DeckID    uuid.UUID `json:"deck_id,omitempty"`
}
