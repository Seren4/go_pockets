package getstatus_test

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGet(t *testing.T) {

	// Create a request.
	req, err := http.NewRequest(http.MethodGet, "/games/", nil)
	require.NoError(t, err)

	// add path parameter
	req.SetPathValue(api.GameID, "33")

	// Create a response recorder
	recorder := httptest.NewRecorder()
	handleFunc := getstatus.Handle(gameFinderStub{})

	// Call the function
	handleFunc(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

// Stubbing the repo
type gameFinderStub struct {
	err  error
	game session.Game
}

func (g gameFinderStub) Find(_ session.GameID) (session.Game, error) {
	return g.game, g.err
}
