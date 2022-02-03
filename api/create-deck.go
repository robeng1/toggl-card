package api

import (
	"net/http"
	"strings"
)

func (api *Handler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var creatDeckBody struct {
		Shuffle bool
	}
	if err := Payload(r, &creatDeckBody); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}
	query := r.URL.Query()
	code := query.Get("cards")
	var codes []string
	if code != "" {
		codes = strings.Split(strings.TrimSpace(code), ",")
	}
	deck, err := api.toggl.CreateDeck(r.Context(), creatDeckBody.Shuffle, codes...)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteData(w, http.StatusCreated, deck)
}
