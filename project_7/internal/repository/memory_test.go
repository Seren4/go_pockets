package repository

import (
	"learngo/httpgordle/internal/session"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	gr := New()

	id := "1"
	newGame := session.Game{
		ID:           session.GameID(id),
		AttemptsLeft: 3,
		Guesses:      make([]session.Guess, 0),
		Status:       session.StatusLost,
	}
	gr.Add(newGame)

	game, err := gr.Find(newGame.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, session.StatusLost, game.Status)

	_, err2 := gr.Find(session.GameID("44"))
	if err2 == nil {
		t.Errorf("expected %q, got: %q", session.ErrNotFound, err2)
	}

	newGame.Status = session.StatusPlaying
	errUpd := gr.Update(newGame)
	assert.Equal(t, nil, errUpd)

	gameUpd, err := gr.Find(newGame.ID)
	assert.Equal(t, session.StatusPlaying, gameUpd.Status)
	assert.Equal(t, nil, err)

}
