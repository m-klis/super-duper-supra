package main

import (
	"log"
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

	log.Print("we're up and running")
	err := http.ListenAndServe(":8080", router)
	exception.PanicIfNeeded(err)
}
