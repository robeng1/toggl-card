package service

import (
	"context"

	"github.com/google/uuid"
)

type TogglRepository interface {
	CreateDeck(ctx context.Context, deck *Deck) error
	UpdateDeck(ctx context.Context, deck *Deck) error
	GetDeck(ctx context.Context, deckId uuid.UUID) (*Deck, error)
}
