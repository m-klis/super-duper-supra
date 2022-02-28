package controller

import "net/http"

type NoteController interface {
	FindAll(http.ResponseWriter, *http.Request)
}
