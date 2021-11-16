package lib

import (
	"encoding/json"
	"net/http"

	"{{.ServiceName}}/pkg/httperr"
)

type RecordingWriter struct {
	http.ResponseWriter
	Status int
	Length int
	Err    httperr.Error
}

func NewRecordingWriter(w http.ResponseWriter) *RecordingWriter {
	return &RecordingWriter{ResponseWriter: w, Status: http.StatusOK, Length: 0, Err: httperr.Error{}}
}

func (rw *RecordingWriter) Write(b []byte) (int, error) {
	if rw.Status != http.StatusOK {
		err := httperr.Error{}
		if unmarshalErr := json.Unmarshal(b, &err); unmarshalErr != nil {
			rw.Err = httperr.UnknownError
		} else {
			rw.Err = err
		}
	}

	n, err := rw.ResponseWriter.Write(b)
	rw.Length += n
	return n, err
}

func (rw *RecordingWriter) WriteHeader(status int) {
	rw.Status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *RecordingWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}
