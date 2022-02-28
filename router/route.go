package router

import (
	"noteapp/controller"

	"github.com/go-chi/chi"
)

func NoteRouter(noteController controller.NoteController) *chi.Mux {
	router := chi.NewMux()

	router.Get("/note", noteController.FindAll)

	return router
}
