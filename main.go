package main

import (
	"fmt"
	"log"
	"net/http"
	"noteapp/config"
	"noteapp/controller"
	"noteapp/repository"
	route "noteapp/router"
	"noteapp/service"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/go-playground/validator/v10"
)

const isoFormat = "2006-01-02T15:04:05.000-0700"
const port = ":48081"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:48081
// @BasePath /note

// // @securityDefinitions.apikey BearerAuth
// // @in header
// // @name Authorization
// // @Security BearerAuth
func main() {
	runtime.GOMAXPROCS(2)
	// setup config
	configuration := config.New()
	db := config.NewPostgresDatabase(configuration)
	defer db.Close()
	validate := validator.New()

	noteRepository := repository.NewNoteRepository()
	noteService := service.NewNoteService(noteRepository, db, validate)
	noteController := controller.NewNoteController(noteService)

	router := route.NoteRouter(noteController)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		log.Print(err)
	}()

	defer route.Stop(server)
	log.Print("we're up and running on " + port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch), "in server")
	log.Println("Stopping API Server")
}
