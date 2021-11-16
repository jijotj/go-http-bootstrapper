package lib_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"{{.ServiceName}}/pkg/lib"
)

func TestWriteResponseJSON(t *testing.T) {
	var w bytes.Buffer
	response := map[string]string{"A": "one", "B": "two"}

	lib.WriteResponseJSON(&w, response)

	assert.JSONEq(t, `{"A": "one", "B": "two"}`, w.String(), "Incorrect response written")
}

func TestWriteResponseJSONWhenError(t *testing.T) {
	var w bytes.Buffer
	response := func() string { return "response" }

	lib.WriteResponseJSON(&w, response)

	assert.Empty(t, w.String(), "Unexpected response written")
}
