package restErrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	// TODO: Complete Test Cases
}
func TestNewBadRequestError(t *testing.T) {
	// TODO: Complete Test Cases
}
func TestNewNotFoundError(t *testing.T) {
	// TODO: Complete Test Cases
}
func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is the message!", errors.New("Database error!"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "This is the message!", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "Database error!", err.Causes[0])
}
