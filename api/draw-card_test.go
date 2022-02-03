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

func TestHandler_DrawCard(t *testing.T) {
	server := testServer()
	defer server.Close()
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	baseURL := server.URL+"/decks"
	t.Run("draw card", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL, "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		var deck service.Deck
		err = json.NewDecoder(resp.Body).Decode(&deck)
		require.NoError(t, err)
		drawCardResp, err := client.Get(baseURL + fmt.Sprintf("/%s/draw/%s", deck.DeckID.String(), "3"))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, drawCardResp.StatusCode)
	})

	t.Run("bad input", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL, "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		var deck service.Deck
		err = json.NewDecoder(resp.Body).Decode(&deck)
		require.NoError(t, err)
		drawCardResp, err := client.Get(baseURL + fmt.Sprintf("/%s/draw/%s", deck.DeckID.String(), "a"))
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, drawCardResp.StatusCode)
	})

	t.Run("invalid ID", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL, "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		var deck service.Deck
		err = json.NewDecoder(resp.Body).Decode(&deck)
		require.NoError(t, err)
		drawCardResp, err := client.Get(baseURL + fmt.Sprintf("/%s/draw/%s", "20", "a"))
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, drawCardResp.StatusCode)
	})

}
