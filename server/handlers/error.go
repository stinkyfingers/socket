package handlers

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Status  int    `json:"statusCode"`
	W       http.ResponseWriter
}

func (h HttpError) HandleErr() {

	errWriter := struct {
		HttpError
		ErrStr string `json:"error"`
	}{
		h,
		"",
	}
	if h.Error != nil {
		errWriter.ErrStr = h.Error.Error()
	}
	j, err := json.Marshal(errWriter)
	if err != nil {
		h.W.Write([]byte("error erroring error"))
		return
	}
	h.W.WriteHeader(h.Status)
	h.W.Header().Set("Content-Type", "application/json")
	h.W.Write(j)
	return
}
