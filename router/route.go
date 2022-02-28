package router

import (
	"noteapp/controller"

	"github.com/go-chi/chi/v5"
)

func NoteRouter(noteController controller.NoteController) *chi.Mux {
	router := chi.NewMux()

	router.Get("/note", noteController.FindAll)

	return router
}
