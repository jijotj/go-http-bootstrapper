package lib_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"{{.ServiceName}}/pkg/httperr"
	"{{.ServiceName}}/pkg/lib"
	"{{.ServiceName}}/pkg/testlib/mocks"
)

func TestWrite(t *testing.T) {
	mockWriter := new(mocks.ResponseWriter)
	recordingWriter := lib.NewRecordingWriter(mockWriter)
	errMsg := httperr.Error{
		HTTPStatus: http.StatusBadGateway,
		Code:       "502",
		Message:    "BAD_GATEWAY",
	}
	bytes, err := json.Marshal(errMsg)
	require.NoError(t, err, "Unexpected JSON marshal error")

	mockWriter.On("Write", bytes).Return(len(bytes), nil).Once()
	mockWriter.On("WriteHeader", http.StatusBadGateway).Return().Once()

	recordingWriter.WriteHeader(http.StatusBadGateway)
	n, err := recordingWriter.Write(bytes)

	assert.NoError(t, err, "Unexpected response write error")
	assert.Equal(t, len(bytes), n, "Incorrect number of bytes written")
	assert.Equal(t, errMsg, recordingWriter.Err, "Incorrect error recorded in writer")

	mock.AssertExpectationsForObjects(t, mockWriter)
}

func TestWriteForJSONUnmarshalError(t *testing.T) {
	mockWriter := new(mocks.ResponseWriter)
	recordingWriter := lib.NewRecordingWriter(mockWriter)

	bytes, err := json.Marshal("{}")
	require.NoError(t, err, "Unexpected JSON marshal error")

	mockWriter.On("Write", bytes).Return(len(bytes), nil).Once()
	mockWriter.On("WriteHeader", http.StatusBadGateway).Return().Once()

	recordingWriter.WriteHeader(http.StatusBadGateway)
	n, err := recordingWriter.Write(bytes)

	assert.NoError(t, err, "Unexpected response write error")
	assert.Equal(t, len(bytes), n, "Incorrect number of bytes written")
	assert.Equal(t, httperr.UnknownError, recordingWriter.Err, "Incorrect error recorded in writer")

	mock.AssertExpectationsForObjects(t, mockWriter)
}

func TestWriteHeader(t *testing.T) {
	mockWriter := new(mocks.ResponseWriter)
	recordingWriter := lib.NewRecordingWriter(mockWriter)

	mockWriter.On("WriteHeader", http.StatusBadGateway).Return().Once()

	recordingWriter.WriteHeader(http.StatusBadGateway)

	assert.Equal(t, http.StatusBadGateway, recordingWriter.Status, "Incorrect header status code")

	mock.AssertExpectationsForObjects(t, mockWriter)
}

func TestHeader(t *testing.T) {
	mockWriter := new(mocks.ResponseWriter)
	recordingWriter := lib.NewRecordingWriter(mockWriter)

	mockWriter.On("Header").Return(nil).Once()

	header := recordingWriter.Header()

	assert.Nil(t, header, "Incorrect header map")

	mock.AssertExpectationsForObjects(t, mockWriter)
}
