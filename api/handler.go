package api

import "github.com/robeng1/toggl/service"

func NewHandler(store service.TogglRepository) *Handler {
	toggl := service.NewTogglService(store)
	return &Handler{toggl: toggl}
}

type Handler struct {
	toggl service.TogglService
}
