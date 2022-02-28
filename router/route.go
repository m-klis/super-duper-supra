package router

import (
	"noteapp/controller"

	"github.com/go-chi/chi"
)

func NoteRouter(noteController controller.NoteController) *chi.Mux {
	router := chi.NewMux()

	router.Route("/note", func(r chi.Router) {
		r.Get("/", noteController.FindNotes)
		r.Post("/", noteController.CreateNote)
		r.Route("/{noteid}", func(r chi.Router) {
			r.Get("/", noteController.FindNote)
		})
	})

	return router
}
