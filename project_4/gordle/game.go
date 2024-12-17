package gordle

import "fmt"

// Game holds all the information we need to play a game of Gordle.
type Game struct{}

// New returns a Game, which can be used to play.
func New() *Game {
	g := &Game{}
	return g
}

// Play runs the game.
// ( creating a methon on an object is achieved by writing a pointer receiver on the Game stucture)
func (g *Game) Play()  {
	fmt.Println("Welcome to Gordie")
	fmt.Println("Enter a guess:\n")
}