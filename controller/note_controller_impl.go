package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"noteapp/exception"
	"noteapp/model"
	"noteapp/service"

	"github.com/go-chi/chi"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (c *NoteControllerImpl) FindNotes(w http.ResponseWriter, r *http.Request) {
	noteResponses, err := c.NoteService.FindNotes(r.Context())

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   noteResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	exception.PanicIfNeeded(err)
}

func (c *NoteControllerImpl) FindNote(w http.ResponseWriter, r *http.Request) {
	noteId := chi.URLParam(r, "noteid")
	fmt.Println(noteId)
	fmt.Println("Cek Atas")
	if noteId == "" {
		exception.PanicIfNeeded(errors.New("noteId nil"))
	}
	res, err := c.NoteService.FindNote(r.Context(), noteId)

	var webResponse model.WebResponse
	if err != nil {
		webResponse = model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Failed",
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	exception.PanicIfNeeded(err)
}
