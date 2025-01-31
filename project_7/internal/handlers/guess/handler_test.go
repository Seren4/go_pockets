package guess_test

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGet(t *testing.T) {

	// Create a request with a non nil body; (body is an io.Reader)
	body := strings.NewReader(`{"guess":"pocket"}`)
	req, err := http.NewRequest(http.MethodPut, "/games/", body)
	require.NoError(t, err)

	// add path parameter
	req.SetPathValue(api.GameID, "34")

	// Create a response recorder
	recorder := httptest.NewRecorder()

	handleFunc := guess.Handle(gameUpdaterStub{})
	// Call the function

	handleFunc(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"34","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

// Stubbing the repo
type gameUpdaterStub struct {
	err error
}

func (g gameUpdaterStub) Update(_ session.Game) error {
	return g.err
}
