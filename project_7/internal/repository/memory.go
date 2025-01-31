package repository

import (
	"fmt"
	"learngo/httpgordle/internal/session"
	"sync"
)

// GameRepository holds all the current games.
type GameRepository struct {
	mutex   sync.Mutex
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	// Lock the reading and the writing of the game.
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	_, found := gr.storage[game.ID]
	if found {
		// Cannot add the same game twice
		return fmt.Errorf("gameID %s already exists", game.ID)
	}
	gr.storage[game.ID] = game

	return nil
}

// Find finds the game in memory (if any).
func (gr *GameRepository) Find(gameID session.GameID) (session.Game, error) {
	// Lock the reading of the game.
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	game, found := gr.storage[gameID]
	if !found {
		// Cannot find the game.
		return session.Game{}, fmt.Errorf("can't find game %s: %q", game.ID, session.ErrNotFound)
	}

	return game, nil
}

// Update updates the game in memory (if any).
func (gr *GameRepository) Update(game session.Game) error {
	_, err := gr.Find(game.ID)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	// Lock the writing of the game.

	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	gr.storage[game.ID] = game

	return nil
}
