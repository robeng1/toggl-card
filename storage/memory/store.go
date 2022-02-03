package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/robeng1/toggl/service"
)

var (
	_ service.TogglRepository = (*Store)(nil)
)

func NewStore() service.TogglRepository{
	d := make(map[uuid.UUID]*service.Deck)
	return &Store{data: d}
}

type Store struct {
	data map[uuid.UUID]*service.Deck
	mu   sync.Mutex
}

func (s *Store) CreateDeck(ctx context.Context, deck *service.Deck) error {
	if _, ok := s.data[deck.DeckID]; ok {
		return errors.New("deck with same ID already exists")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[deck.DeckID] = deck
	return nil
}
func (s *Store) UpdateDeck(ctx context.Context, deck *service.Deck) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[deck.DeckID] = deck
	return nil
}
func (s *Store) GetDeck(ctx context.Context, deckId uuid.UUID) (*service.Deck, error) {
	if deck, ok := s.data[deckId]; !ok {
		return nil, errors.New("deck was not found")
	} else {
		return deck, nil
	}
}
