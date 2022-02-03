package service

import (
	"context"

	"github.com/google/uuid"
)

var (
	_ TogglService = (*togglProvider)(nil)
)

func NewTogglService(store TogglRepository) TogglService {
	return &togglProvider{store: store}
}

type TogglService interface {
	CreateDeck(ctx context.Context, shuffle bool, codes ...string) (*Deckdto, error)
	OpenDeck(ctx context.Context, deckId uuid.UUID) (*Deck, error)
	DrawCard(ctx context.Context, deckId uuid.UUID, numberToDraw int) ([]*Card, error)
}

type togglProvider struct {
	store TogglRepository
}

func (t *togglProvider) CreateDeck(ctx context.Context, shuffle bool, codes ...string) (*Deckdto, error) {
	deck := &Deck{
		DeckID: uuid.New(),
	}
	if codes != nil {
		// the user has supplied initial card codes so we must use that
		deck.Cards = findCards(codes...)
	} else {
		deck.Cards = allCards()
	}
	if shuffle {
		deck.Shuffle()
	}
	if err := t.store.CreateDeck(ctx, deck); err != nil {
		return nil, err
	}
	return deck.WithoutCards(), nil
}

func (t *togglProvider) OpenDeck(ctx context.Context, deckId uuid.UUID) (*Deck, error) {
	return t.store.GetDeck(ctx, deckId)
}

func (t *togglProvider) DrawCard(ctx context.Context, deckId uuid.UUID, count int) ([]*Card, error) {
	deck, err := t.store.GetDeck(ctx, deckId)
	if err != nil {
		return nil, err
	}
	if deck.Remaining() < count {
		return nil, ErrInvalidNumberOfCardsToDraw
	}
	drawn, remaining := deck.Cards[0:count], deck.Cards[count:len(deck.Cards)]
	deck.Cards = remaining
	if err := t.store.UpdateDeck(ctx, deck); err != nil {
		return nil, err
	}
	return drawn, nil
}
