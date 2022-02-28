package controller

import "net/http"

type NoteController interface {
	CreateNote(http.ResponseWriter, *http.Request)
	FindNotes(http.ResponseWriter, *http.Request)
	FindNote(http.ResponseWriter, *http.Request)
}
