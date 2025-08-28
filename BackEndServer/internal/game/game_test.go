package game

import (
	"fmt"
	"testing"
)

func TestGameInstantiation(t *testing.T) {
	parameters := struct {
		width       int
		height      int
		playerCount int
	}{
		10, 10, 2,
	}
	game, err := NewGame(parameters.width, parameters.height, parameters.playerCount)
	if err != nil {
		t.Fatalf("Game Initialization Failed")
	}

	if game.Height != parameters.height {
		t.Fatalf("Game Height incorrectly set")
	}

	if game.Width != parameters.width {
		t.Fatalf("Game Width incorrectly set")
	}

	if game.TotalPlayers != parameters.playerCount {
		t.Fatalf("Player count incorrectly set")
	}

	if len(game.SquaresByOpenEdges) != 5 {
		t.Fatalf("Squares by open edges incorrect possibility setup")
	}
}

func TestGameInitialize(t *testing.T) {
	parameters := struct {
		width       int
		height      int
		playerCount int
	}{
		10, 10, 2,
	}
	testEdges := []string{
		fmt.Sprintf("%d,%d-%d,%d", 1, 1, 2, 1),
		fmt.Sprintf("%d,%d-%d,%d", 1, parameters.height-1, 2, parameters.height-1),
		fmt.Sprintf("%d,%d-%d,%d", parameters.width-1, 1, parameters.width-1, 2),
		fmt.Sprintf("%d,%d-%d,%d", 1, 1, 1, 2),
	}

	game, err := NewGame(parameters.width, parameters.height, parameters.playerCount)
	if err != nil {
		t.Fatalf("Game Initialization Failed")
	}

	game.Initialize()

	if len(game.Nodes) != (game.Height-1)*(game.Width-1) {
		t.Fatalf("Node count is not correctly set")
	}

	if game.SquaresByOpenEdges[0].Count != 0 {
		t.Fatalf("Enclosed squares exist in new game")
	}

	if game.SquaresByOpenEdges[1].Count != 0 {
		t.Fatalf("Near-enclosed squares exist in new game")
	}

	for _, tt := range testEdges {
		val, exists := game.Edges[tt]
		if !exists {
			t.Fatalf("Border edge not set %v", tt)
		}
		if !val {
			t.Fatalf("Border edge value improperly configured for edge %v", tt)
		}
	}

}
