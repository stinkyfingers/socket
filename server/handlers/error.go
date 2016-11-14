package handlers

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Status  int    `json:"statusCode"`
	w       http.ResponseWriter
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
		h.w.Write([]byte("error erroring error"))
		return
	}
	h.w.WriteHeader(h.Status)
	h.w.Header().Set("Content-Type", "application/json")
	h.w.Write(j)
	return
}
