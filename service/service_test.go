package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/robeng1/toggl/service"
	"github.com/robeng1/toggl/storage/memory"
)

func TestToggleService(t *testing.T) {
	for _, tester := range serviceTests {
		store := memory.NewStore()
		svc := service.NewTogglService(store)
		tester(t, svc)
	}
}

var serviceTests = []func(*testing.T, service.TogglService){
	testCreateDeck,
	testCreateDeckWithCodes,
	testOpenDeck,
	testDrawCard,
}

func testCreateDeck(t *testing.T, store service.TogglService) {
	deck, err := store.CreateDeck(context.Background(), false)
	require.NoError(t, err)
	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, 52, deck.Remaining)
}

func testCreateDeckWithCodes(t *testing.T, store service.TogglService) {
	codes := []string{"AS", "2C", "5H"}
	deck, err := store.CreateDeck(context.Background(), false, codes...)
	require.NoError(t, err)
	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, 3, deck.Remaining)
	actualDeck, err := store.OpenDeck(context.Background(), deck.DeckID)
	require.NoError(t, err)
	assert.Equal(t, "AS", actualDeck.Cards[0].Code)
	assert.Equal(t, "2C", actualDeck.Cards[1].Code)
	assert.Equal(t, "5H", actualDeck.Cards[2].Code)
}

func testOpenDeck(t *testing.T, store service.TogglService) {
	deck, err := store.CreateDeck(context.Background(), false)
	require.NoError(t, err)
	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, 52, deck.Remaining)
	actualDeck, err := store.OpenDeck(context.Background(), deck.DeckID)
	require.NoError(t, err)
	assert.Equal(t, false, actualDeck.Shuffled)
	assert.Equal(t, 52, actualDeck.Remaining())
}

func testDrawCard(t *testing.T, store service.TogglService) {
	deck, err := store.CreateDeck(context.Background(), false)
	require.NoError(t, err)
	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, 52, deck.Remaining)
	actualDeck, err := store.OpenDeck(context.Background(), deck.DeckID)
	require.NoError(t, err)
	assert.Equal(t, false, actualDeck.Shuffled)
	assert.Equal(t, 52, actualDeck.Remaining())
	drawn, err := store.DrawCard(context.Background(), actualDeck.DeckID, 5)
	require.NoError(t, err)
	assert.Equal(t, 5, len(drawn))
	assert.Equal(t, 47, actualDeck.Remaining())
}
