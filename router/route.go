package router

import (
	"context"
	"log"
	"net/http"
	"noteapp/controller"
	"os"
	"time"

	_ "noteapp/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NoteRouter(noteController controller.NoteController) http.Handler {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)

	r.Route("/note", func(r chi.Router) {
		r.Get("/", noteController.FindNotes)
		r.Post("/", noteController.CreateNote)
		r.Route("/{noteid}", func(r chi.Router) {
			r.Get("/", noteController.FindNote)
			r.Put("/", noteController.UpdateNote)
			r.Delete("/", noteController.DeleteNote)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:48081/swagger/doc.json"), //The url pointing to API definition
	))

	return r
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server corectly: %v\n", err)
		os.Exit(1)
	}
}
