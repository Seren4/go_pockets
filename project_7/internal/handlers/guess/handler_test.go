package guess_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/api"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

)


func TestHandleGet(t *testing.T) {

	// Create a request with a non nil body; (body is an io.Reader)
	body := strings.NewReader(`{"guess":"pocket"}`)
	req, err := http.NewRequest(http.MethodPut,"/games/", body)
	require.NoError(t, err)
	
	// add path parameter
	req.SetPathValue(api.GameID, "34") 

	// Create a response recorder
	recorder := httptest.NewRecorder() 

	// Call the function
	guess.Handle(recorder, req)  
	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t,"application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"34","attempts_left":0,"guesses":null,"word_length":0,"status":""}`, recorder.Body.String())
}