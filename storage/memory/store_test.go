package memory_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/robeng1/toggl/service"
	"github.com/robeng1/toggl/storage/memory"
)

func TestMemoryStore(t *testing.T) {
	for _, tester := range memoryStoreTesters {
		store := memory.NewStore()
		tester(t, store)
	}
}

var memoryStoreTesters = []func(*testing.T, service.TogglRepository){
	testCreateDeck,
	testUpdateDeck,
	testGetDeck,
}

func testCreateDeck(t *testing.T, store service.TogglRepository) {
	id := uuid.New()
	deck := &service.Deck{DeckID: id, Shuffled: false}
	err := store.CreateDeck(context.Background(), deck)
	require.NoError(t, err)
	assert.Equal(t, id, deck.DeckID)
	assert.Equal(t, false, deck.Shuffled)
}

func testUpdateDeck(t *testing.T, store service.TogglRepository) {
	id := uuid.New()
	deck := &service.Deck{DeckID: id, Shuffled: false}
	err := store.CreateDeck(context.Background(), deck)
	require.NoError(t, err)
	assert.Equal(t, id, deck.DeckID)
	assert.Equal(t, false, deck.Shuffled)
	deck.Shuffled = true
	err = store.UpdateDeck(context.Background(), deck)
	require.NoError(t, err)
	deck, err = store.GetDeck(context.Background(), id)
	require.NoError(t, err)
	assert.Equal(t, id, deck.DeckID)
	assert.Equal(t, true, deck.Shuffled)
}

func testGetDeck(t *testing.T, store service.TogglRepository) {
	id := uuid.New()
	deck := &service.Deck{DeckID: id, Shuffled: false}
	err := store.CreateDeck(context.Background(), deck)
	require.NoError(t, err)
	assert.Equal(t, id, deck.DeckID)
	assert.Equal(t, false, deck.Shuffled)
	deck.Shuffled = true
	err = store.UpdateDeck(context.Background(), deck)
	require.NoError(t, err)
	deck, err = store.GetDeck(context.Background(), id)
	require.NoError(t, err)
	assert.Equal(t, id, deck.DeckID)
	assert.Equal(t, true, deck.Shuffled)
}
