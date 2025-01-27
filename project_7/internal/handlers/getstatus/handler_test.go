package getstatus_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/api"



	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

)


func TestHandleGet(t *testing.T) {

	// Create a request.
	req, err := http.NewRequest(http.MethodGet, "/games/", nil)  
	require.NoError(t, err)

	// add path parameter
	req.SetPathValue(api.GameID, "33") 

	// Create a response recorder
	recorder := httptest.NewRecorder() 

	// Call the function
	getstatus.Handle(recorder, req)  
	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"33","attempts_left":0,"guesses":null,"word_length":0,"status":""}`, recorder.Body.String())
}