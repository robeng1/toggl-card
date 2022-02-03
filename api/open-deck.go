package api

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (api *Handler) OpenDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, ok := vars["id"]; !ok {
		WriteError(w, http.StatusBadRequest, errors.New("must provide deck id"))
	} else {
		deckId, err := uuid.Parse(id)
		if err != nil {
			WriteError(w, http.StatusBadRequest, errors.New("invalid ID"))
		}
		deck, err := api.toggl.OpenDeck(r.Context(), deckId)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, err)
			return
		}
		WriteData(w, http.StatusOK, deck)
	}
}
