package guess

import ( 
	"net/http"
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"log"
)



// The benefit of this method over http.Handle is that we don’t have to 
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

func Handle(w http.ResponseWriter, req *http.Request) {

	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the ID of the game", http.StatusBadRequest)
		return
	}
	// TODO verify id
	// log package is not thread- safe.
	// So keep in mind that it can lead to unordered logs and complicate later testing.
	log.Printf("retrieve status of the game with id: %v", id)

	// Read the request, containing the guess, from the body of the input.
	// Instantiate the request struct
	r := api.GuessRequest{} 
	// Decode the body
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		// If the JSON doesn’t parse, it’s a bad request
		http.Error(w, err.Error(), http.StatusBadRequest)  
		return
	}

	apiGame := api.GameResponse{
		ID: id,
	}
	// Encode the game into JSON
	error := json.NewEncoder(w).Encode(apiGame)   
	if error != nil {             
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)                                
		}     

}