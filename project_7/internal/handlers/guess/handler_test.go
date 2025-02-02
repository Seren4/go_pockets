package guess_test

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/session"
	"learngo/httpgordle/internal/gordle"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGet(t *testing.T) {

	game, _ := gordle.New([]string { "pocket" })

	handleFunc := guess.Handle(gameGuesserStub{session.Game{
		ID:           "34",
		Gordle:       *game,
		AttemptsLeft: 3,
		Guesses:      nil,
		Status:       session.StatusPlaying,
	}})


	// Create a request with a non nil body; (body is an io.Reader)
	body := strings.NewReader(`{"guess":"pocket"}`)
	req, err := http.NewRequest(http.MethodPut, "/games/", body)
	require.NoError(t, err)

	// add path parameter
	req.SetPathValue(api.GameID, "34")

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the function

	handleFunc(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"34","attempts_left":2,"guesses":[{"word":"pocket", "feedback":"✅✅✅✅✅✅"}],"word_length":0,"status":"Won"}`, recorder.Body.String())
}

// Stubbing the repo
type gameGuesserStub struct {
	game session.Game
}

func (g gameGuesserStub) Update(_ session.Game) error {
	return nil
}

func (g gameGuesserStub) Find(_ session.GameID) (session.Game, error) {
	return g.game, nil
}
