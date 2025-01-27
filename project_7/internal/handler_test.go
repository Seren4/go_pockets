package internal

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"learngo/httpgordle/internal/handlers/newgame"


	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

)


func TestHandle(t *testing.T) {
	// Create a request.
	req, err := http.NewRequest(http.MethodPost, "/games", nil)  
	require.NoError(t, err)
	// Create a response recorder
	recorder := httptest.NewRecorder() 

	// Call the function
	newgame.Handle(recorder, req)  
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":null,"word_length":0,"status":""}`, recorder.Body.String())
}