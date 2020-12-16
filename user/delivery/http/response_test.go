package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	r := &Response{
		Status:  200,
		Data:    "Hello",
		Message: "Hello message",
	}

	rec := httptest.NewRecorder()
	err := r.JSON(rec)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
