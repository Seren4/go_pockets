package newgame

import "net/http"



// The benefit of this method over http.Handle is that we don’t have to 
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

func Handle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game"))
}