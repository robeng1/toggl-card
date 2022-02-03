package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/robeng1/toggl/service"
)

func TestHandler_OpenDeck(t *testing.T) {
	server := testServer()
	defer server.Close()
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	baseURL := server.URL+"/decks"

	t.Run("open deck", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL, "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		var deck service.Deckdto
		err = json.NewDecoder(resp.Body).Decode(&deck)
		require.NoError(t, err)
		openedResp, err := client.Get(baseURL + fmt.Sprintf("/%s", deck.DeckID.String()))
		require.NoError(t, err)

		var openedDeck service.Deck
		err = json.NewDecoder(openedResp.Body).Decode(&openedDeck)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, openedResp.StatusCode)
		assert.Equal(t, 52, openedDeck.Remaining())
		assert.Equal(t, deck.Remaining, openedDeck.Remaining())
	})

	
}
