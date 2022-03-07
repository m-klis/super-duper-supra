package controller

import (
	"encoding/json"
	"errors"
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
	var webResponse = model.NewWebResponse()
	noteResponses, err := c.NoteService.FindNotes(r.Context())
	exception.CheckError(err)

	if err != nil {
		webResponse.Code = http.StatusBadRequest
		webResponse.Status = "Failed"
		webResponse.Message = err
	} else {
		webResponse.Code = http.StatusOK
		webResponse.Status = "OK"
		webResponse.Data = noteResponses
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	exception.CheckError(err)
}

func (c *NoteControllerImpl) FindNote(w http.ResponseWriter, r *http.Request) {
	var webResponse = model.NewWebResponse()
	noteId := chi.URLParam(r, "noteid")
	if noteId == "" {
		exception.CheckError(errors.New("noteId nil"))
		webResponse.Code = http.StatusBadRequest
		webResponse.Status = "Failed"
		webResponse.Message = "check id again"
	} else {
		res, err := c.NoteService.FindNote(r.Context(), noteId)
		if err != nil {
			webResponse.Code = http.StatusBadRequest
			webResponse.Status = "Failed"
			webResponse.Message = err
		} else {
			webResponse.Code = http.StatusOK
			webResponse.Status = "OK"
			webResponse.Data = res
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	exception.CheckError(err)
}

func (c *NoteControllerImpl) CreateNote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	noteCreateRequest := model.NoteCreateRequest{}
	err := decoder.Decode(&noteCreateRequest)
	// exception.CheckError(err)
	var webResponse = model.NewWebResponse()
	if err != nil {
		webResponse.Code = http.StatusNotAcceptable
		webResponse.Status = "Failed"
		webResponse.Message = err
	} else {
		noteResponse, err := c.NoteService.CreateNote(r.Context(), noteCreateRequest)
		exception.CheckError(err)
		if err != nil {
			webResponse.Code = http.StatusBadRequest
			webResponse.Status = "Failed"
			webResponse.Message = err
		} else {
			webResponse.Code = http.StatusOK
			webResponse.Status = "OK"
			webResponse.Data = noteResponse
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	exception.CheckError(err)
}

func (c *NoteControllerImpl) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var webResponse = model.NewWebResponse()
	noteId := chi.URLParam(r, "noteid")
	if noteId == "" {
		webResponse.Code = http.StatusBadRequest
		webResponse.Status = "Failed"
		webResponse.Message = "check id again"
	} else {
		decoder := json.NewDecoder(r.Body)
		noteUpdateRequest := model.NoteUpdateRequest{}
		err := decoder.Decode(&noteUpdateRequest)
		exception.CheckError(err)
		if err != nil {
			webResponse.Code = http.StatusNotAcceptable
			webResponse.Status = "Failed"
			webResponse.Message = err
		} else {
			noteUpdateRequest.ID = noteId
			noteResponse, err := c.NoteService.UpdateNote(r.Context(), noteUpdateRequest)
			exception.CheckError(err)
			webResponse.Code = http.StatusOK
			webResponse.Status = "OK"
			webResponse.Data = noteResponse
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	exception.CheckError(err)
}

func (c *NoteControllerImpl) DeleteNote(w http.ResponseWriter, r *http.Request) {
	var webResponse = model.NewWebResponse()
	noteId := chi.URLParam(r, "noteid")
	if noteId == "" {
		webResponse.Code = http.StatusBadRequest
		webResponse.Status = "Failed"
		webResponse.Message = "check id again"
	} else {
		err := c.NoteService.DeleteNote(r.Context(), noteId)
		if err != nil {
			webResponse.Code = http.StatusBadRequest
			webResponse.Status = "Failed"
			webResponse.Message = err
		} else {
			webResponse.Code = http.StatusOK
			webResponse.Status = "OK"
			webResponse.Message = "note deleted"
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	exception.CheckError(err)
}
