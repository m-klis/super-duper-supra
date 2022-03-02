package router

import (
	"noteapp/controller"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func NoteRouter(noteController controller.NoteController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/note", func(r chi.Router) {
		r.Get("/", noteController.FindNotes)
		r.Post("/", noteController.CreateNote)
		r.Route("/{noteid}", func(r chi.Router) {
			r.Get("/", noteController.FindNote)
		})
	})

	return r
}
