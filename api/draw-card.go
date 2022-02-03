package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (api *Handler) DrawCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, ok := vars["id"]; !ok {
		WriteError(w, http.StatusBadRequest, errors.New("must provide deck id"))
	} else {
		deckId, err := uuid.Parse(id)
		if err != nil {
			WriteError(w, http.StatusBadRequest, errors.New("invalid ID"))
		}
		var count int
		if n, ok := vars["count"]; !ok {
			WriteError(w, http.StatusBadRequest, errors.New("must provide number to draw"))
		} else {
			count, err = strconv.Atoi(n)
			if err != nil {
				WriteError(w, http.StatusBadRequest, errors.New("not a valid number"))
			}
		}
		cards, err := api.toggl.DrawCard(r.Context(), deckId, count)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, err)
			return
		}
		resp := make(map[string]interface{})
		resp["cards"] = cards
		WriteData(w, http.StatusOK, cards)
	}
}
