package newgame_test

import (
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/session"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	// Create anonymous function
	handleFunc := newgame.Handle(gameAdderStub{})
	// Create a request.
	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)
	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the function
	handleFunc(recorder, req)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

// Stubbing the repo
type gameAdderStub struct {
	err error
}

func (g gameAdderStub) Add(_ session.Game) error {
	return g.err
}
