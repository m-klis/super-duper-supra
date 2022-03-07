package controller

import "net/http"

type NoteController interface {
	FindNotes(http.ResponseWriter, *http.Request)
	FindNote(http.ResponseWriter, *http.Request)
	CreateNote(http.ResponseWriter, *http.Request)
	UpdateNote(http.ResponseWriter, *http.Request)
	DeleteNote(http.ResponseWriter, *http.Request)
}
