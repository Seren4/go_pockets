package api

import "learngo/httpgordle/internal/session"

// ToGameResponse converts a session.Game into an GameResponse.
func ToGameResponse(game session.Game) GameResponse {
	apigame := GameResponse{
		ID:           string(game.ID),
		AttemptsLeft: game.AttemptsLeft,
		Guesses:      make([]Guess, len(game.Guesses)),
		Status:       string(game.Status),
		// TODO WordLength
	}

	for index := 0; index < len(game.Guesses); index++ {
		apigame.Guesses[index].Word = game.Guesses[index].Word
		apigame.Guesses[index].Feedback = game.Guesses[index].Feedback

	}

	if game.AttemptsLeft == 0 {
		apigame.Solution = "" // TODO solution
	}
	return apigame
}
