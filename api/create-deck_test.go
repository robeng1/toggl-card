package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/robeng1/toggl/service"
	"github.com/robeng1/toggl/storage/memory"
)

func testServer() *httptest.Server {
	router := mux.NewRouter()
	handler := NewHandler(memory.NewStore())
	router.HandleFunc("/decks", handler.CreateDeck)
	router.HandleFunc("/decks/{id}", handler.OpenDeck)
	router.HandleFunc("/decks/{id}/draw/{count}", handler.DrawCard)
	return httptest.NewServer(router)
}

func TestHandler_CreateDeck(t *testing.T) {
	server := testServer()
	defer server.Close()
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	baseURL := server.URL+"/decks"

	t.Run("create deck", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL, "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		defer resp.Body.Close()
	})

	t.Run("create deck with codes", func(t *testing.T) {
		requestBody := bytes.NewBuffer([]byte(`{"shuffle": false}`))
		resp, err := client.Post(baseURL+"?cards=AS,KD,AC,2C,KH", "application/json", requestBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		defer resp.Body.Close()
		var deck service.Deckdto
		err = json.NewDecoder(resp.Body).Decode(&deck)
		require.NoError(t, err)
		assert.Equal(t, 5, deck.Remaining)
	})

}
