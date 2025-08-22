package game

import (
	uuid "github.com/google/uuid"

	"github.com/amonjerro/dots-and-boxes/internal/utils"
)

type PlayerType int

const (
	HumanPlayer PlayerType = iota
	AIPlayer
)

type Game struct {
	identifier         uuid.UUID
	width              int
	height             int
	playerOrder        []PlayerType
	totalPlayers       int
	currentPlayerTurn  int
	scores             []int
	squaresByOpenEdges []utils.Set[string]
}

func NewGame(width int, height int) (*Game, error) {

	squaresByOpenEdges := []utils.Set[string]{
		{Elements: make(map[string]bool)}, // 0
		{Elements: make(map[string]bool)}, // 1
		{Elements: make(map[string]bool)}, // 2
		{Elements: make(map[string]bool)}, // 3
		{Elements: make(map[string]bool)}, // 4
	}

	return &Game{
		identifier:         uuid.New(),
		width:              width,
		height:             height,
		squaresByOpenEdges: squaresByOpenEdges,
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

func (g *Game) GetNextPlayerType() PlayerType {
	return g.playerOrder[g.currentPlayerTurn]
}
