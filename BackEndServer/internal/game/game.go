package game

import (
	uuid "github.com/google/uuid"
)

type PlayerType int

const (
	HumanPlayer PlayerType = iota
	AIPlayer
)

type Game struct {
	identifier        uuid.UUID
	width             int
	height            int
	playerOrder       []PlayerType
	totalPlayers      int
	currentPlayerTurn int
	scores            []int
	openSquares       map[int][]string
}

func NewGame(width int, height int) (*Game, error) {

	return &Game{
		identifier: uuid.New(),
		width:      width,
		height:     height,
	}, nil
}

func (g *Game) UpdateCurrentTurn() {
	g.currentPlayerTurn = g.currentPlayerTurn + 1
	if g.currentPlayerTurn > g.totalPlayers {
		g.currentPlayerTurn = 0
	}
}

func (g *Game) Initialize(width int, height int) {
	g.scores = make([]int, g.totalPlayers)

}
