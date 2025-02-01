package newgame_test

import (
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/session"

	"net/http"
	"net/http/httptest"
	"testing"
	"regexp"
	"strings"

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

	// idFinderRegexp is a regular expression that will ensure the body contains an id field with a value that contains
	// only letters (uppercase and/or lowercase) and/or digits.
	idFinderRegexp := regexp.MustCompile(`.+"id":"([a-zA-Z0-9]+)".+`) 
	id := idFinderRegexp.FindStringSubmatch(recorder.Body.String())  
	body := strings.Replace(recorder.Body.String(), id[1], "123456", 1)                  
	if len(id) != 2 {                                               
		t.Fatal("cannot find one id in the json output")
	}
	assert.JSONEq(t, `{"id":"123456","attempts_left":3,"guesses":[],"word_length":0,"status":"Playing"}`, body) 
}

// Stubbing the repo
type gameAdderStub struct {
	err error
}

func (g gameAdderStub) Add(_ session.Game) error {
	return g.err
}
