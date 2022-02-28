package main

import (
	"net/http"
	"noteapp/config"
	"noteapp/controller"
	"noteapp/exception"
	"noteapp/repository"
	"noteapp/router"
	"noteapp/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

func main() {
	// setup config
	configuration := config.New()
	db := config.NewPostgresDatabase(configuration)
	validate := validator.New()

	noteRepository := repository.NewNoteRepository()
	noteService := service.NewNoteService(noteRepository, db, validate)
	noteController := controller.NewNoteController(noteService)

	router := router.NoteRouter(noteController)

	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("Hello super duper supra")))
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	exception.PanicIfNeeded(err)
}
