package httperr_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"{{.ServiceName}}/pkg/httperr"
)

func TestError(t *testing.T) {
	err := httperr.Error{HTTPStatus: http.StatusNotFound, Code: "101", Message: "NOT_FOUND"}

	assert.Equal(t, "NOT_FOUND", err.Error(), "Incorrect error")
}
