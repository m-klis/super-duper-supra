package main

import (
	"fmt"
	"net/http"
	"noteapp/config"
	"noteapp/controller"
	"noteapp/exception"
	"noteapp/repository"
	"noteapp/router"
	"noteapp/service"

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

	// app := chi.NewRouter()
	// app.Use(middleware.Logger)
	// app.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write(([]byte("Hello super duper supra")))
	// })

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("server started at localhost:3000")
	err := server.ListenAndServe()
	exception.PanicIfNeeded(err)
}
