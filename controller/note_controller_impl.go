package controller

import (
	"encoding/json"
	"net/http"
	"noteapp/exception"
	"noteapp/model"
	"noteapp/service"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (c *NoteControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	noteResponses := c.NoteService.FindAll(r.Context())
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   noteResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	exception.PanicIfNeeded(err)
}
