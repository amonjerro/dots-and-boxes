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
	identifier         uuid.UUID
	width              int
	height             int
	playerOrder        []PlayerType
	totalPlayers       int
	currentPlayerTurn  int
	scores             []int
	edges              map[string]bool
	nodes              []string
	squares            map[string]Square
	squaresByOpenEdges []utils.Set[string]
}

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
		identifier:         uuid.New(),
		width:              width,
		height:             height,
		nodes:              nodes,
		edges:              make(map[string]bool),
		squares:            make(map[string]Square),
		squaresByOpenEdges: squaresByOpenEdges,
		totalPlayers:       playerCount,
	}, nil
}

func (g *Game) UpdateCurrentTurn() {
	g.currentPlayerTurn = g.currentPlayerTurn + 1
	if g.currentPlayerTurn > g.totalPlayers {
		g.currentPlayerTurn = 0
	}
}

func (g *Game) Initialize() {
	// Set up
	g.scores = make([]int, g.totalPlayers)

	// Nodes
	for x := 1; x < g.width; x++ {
		for y := 1; y < g.height; y++ {
			coordinateString := fmt.Sprintf("%d,%d", x, y)
			g.nodes = append(g.nodes, coordinateString)
		}
	}

	// Exterior Edges
	for x := 1; x < g.width-1; x++ {
		top := fmt.Sprintf("%d,1-%d,1", x, x+1)
		bot := fmt.Sprintf("%d,%d-%d,%d", x, g.height-1, x+1, g.height-1)
		g.edges[top] = true
		g.edges[bot] = true
	}

	for y := 1; y < g.height-1; y++ {
		left := fmt.Sprintf("1,%d-1,%d", y, y+1)
		right := fmt.Sprintf("%d,%d-%d,%d", g.width-1, y, g.width-1, y+1)
		g.edges[left] = true
		g.edges[right] = true
	}

	// Interior Edges
	for x := 1; x < g.width-1; x++ {
		for y := 1; y < g.height-1; y++ {
			coordinateString := fmt.Sprintf("%d,%d", x, y)
			topEdge := fmt.Sprintf("%d,%d-%d,%d", x, y, x+1, y)
			botEdge := fmt.Sprintf("%d,%d-%d,%d", x, y+1, x+1, y+1)
			leftEdge := fmt.Sprintf("%d,%d-%d,%d", x, y, x, y+1)
			rightEdge := fmt.Sprintf("%d,%d-%d,%d", x+1, y, x+1, y+1)

			count := utils.Reduce([]string{topEdge, botEdge, leftEdge, rightEdge}, func(acc int, val string) int {
				_, exists := g.edges[val]
				if exists {
					return acc
				}
				return acc + 1
			}, 0)
			square := Square{
				edges:               []string{topEdge, botEdge, leftEdge, rightEdge},
				openConnectionCount: count,
			}
			g.squares[coordinateString] = square
			g.squaresByOpenEdges[count].Add(coordinateString)
		}
	}
}

func (g *Game) GetNextPlayerType() PlayerType {
	return g.playerOrder[g.currentPlayerTurn]
}
