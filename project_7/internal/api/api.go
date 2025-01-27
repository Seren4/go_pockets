package api


const (
	  // GameID is the name of the field that stores the game's identifier
		GameID = "id"

    // NewGameRoute is the path to create a new game.
		NewGameRoute  = "/games"
		
		// GetStatusRoute is the path to get the status of a game identified by its id.
		GetStatusRoute = "/games/{" + GameID + "}"
)

// GameResponse contains the information about a game.
type GameResponse struct {
	ID string `json:"id"`
	AttemptsLeft byte `json:"attempts_left"`
	Guesses []Guess `json:"guesses"`                             
	WordLength byte `json:"word_length"`
	// The solution is only given when the game is over, omit otherwise
	Solution string `json:"solution,omitempty"`                  
	Status string `json:"status"`
}

// Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word string `json:"word"`
	Feedback string `json:"feedback"`      
}