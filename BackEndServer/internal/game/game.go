package game

import (
	"fmt"

	uuid "github.com/google/uuid"

	"github.com/amonjerro/dots-and-boxes/internal/utils"
)

type PlayerType int

const (
	HumanPlayer PlayerType = iota
	AIPlayer
)

type Square struct {
	edges               []string
	openConnectionCount int
}

type Game struct {
	Identifier         uuid.UUID           `json:"game_id"`
	Width              int                 `json:"width"`
	Height             int                 `json:"height"`
	PlayerOrder        []PlayerType        `json:"player_order,omitempty"`
	TotalPlayers       int                 `json:"total_players"`
	CurrentPlayerTurn  int                 `json:"current_player_turn"`
	Scores             []int               `json:"scores,omitempty"`
	Edges              map[string]bool     `json:"edges,omitempty"`
	Nodes              []string            `json:"nodes,omitempty"`
	Squares            map[string]Square   `json:"squares"`
	SquaresByOpenEdges []utils.Set[string] `json:"squares_by_edge"`
}

// Creates the pointer to a new game
func NewGame(width int, height int, playerCount int) (*Game, error) {

	squaresByOpenEdges := []utils.Set[string]{
		{Elements: make(map[string]bool), Count: 0}, // 0
		{Elements: make(map[string]bool), Count: 0}, // 1
		{Elements: make(map[string]bool), Count: 0}, // 2
		{Elements: make(map[string]bool), Count: 0}, // 3
		{Elements: make(map[string]bool), Count: 0}, // 4

	}

	nodes := []string{}

	return &Game{
		Identifier:         uuid.New(),
		Width:              width,
		Height:             height,
		Nodes:              nodes,
		Edges:              make(map[string]bool),
		Squares:            make(map[string]Square),
		SquaresByOpenEdges: squaresByOpenEdges,
		TotalPlayers:       playerCount,
	}, nil
}

// Updates a game's current turn marker
func (g *Game) UpdateCurrentTurn() {
	g.CurrentPlayerTurn = g.CurrentPlayerTurn + 1
	if g.CurrentPlayerTurn > g.TotalPlayers {
		g.CurrentPlayerTurn = 0
	}
}

// Initalizes a games internal variables and states when creating it from scratch
func (g *Game) Initialize() {
	// Set up
	g.Scores = make([]int, g.TotalPlayers)

	// Nodes
	for x := 1; x < g.Width; x++ {
		for y := 1; y < g.Height; y++ {
			coordinateString := fmt.Sprintf("%d,%d", x, y)
			g.Nodes = append(g.Nodes, coordinateString)
		}
	}

	// Exterior Edges
	for x := 1; x < g.Width-1; x++ {
		top := fmt.Sprintf("%d,1-%d,1", x, x+1)
		bot := fmt.Sprintf("%d,%d-%d,%d", x, g.Height-1, x+1, g.Height-1)
		g.Edges[top] = true
		g.Edges[bot] = true
	}

	for y := 1; y < g.Height-1; y++ {
		left := fmt.Sprintf("1,%d-1,%d", y, y+1)
		right := fmt.Sprintf("%d,%d-%d,%d", g.Width-1, y, g.Width-1, y+1)
		g.Edges[left] = true
		g.Edges[right] = true
	}

	// Interior Edges
	for x := 1; x < g.Width-1; x++ {
		for y := 1; y < g.Height-1; y++ {
			coordinateString := fmt.Sprintf("%d,%d", x, y)
			topEdge := fmt.Sprintf("%d,%d-%d,%d", x, y, x+1, y)
			botEdge := fmt.Sprintf("%d,%d-%d,%d", x, y+1, x+1, y+1)
			leftEdge := fmt.Sprintf("%d,%d-%d,%d", x, y, x, y+1)
			rightEdge := fmt.Sprintf("%d,%d-%d,%d", x+1, y, x+1, y+1)

			count := utils.Reduce([]string{topEdge, botEdge, leftEdge, rightEdge}, func(acc int, val string) int {
				_, exists := g.Edges[val]
				if exists {
					return acc
				}
				return acc + 1
			}, 0)
			square := Square{
				edges:               []string{topEdge, botEdge, leftEdge, rightEdge},
				openConnectionCount: count,
			}
			g.Squares[coordinateString] = square
			g.SquaresByOpenEdges[count].Add(coordinateString)
		}
	}
}

// Gets the type of the next player. Important for AI player turns
func (g *Game) GetNextPlayerType() PlayerType {
	return g.PlayerOrder[g.CurrentPlayerTurn]
}
