package controller

import "net/http"

type NoteController interface {
	FindNotes(http.ResponseWriter, *http.Request)
	FindNote(http.ResponseWriter, *http.Request)
}
