package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/amonjerro/dots-and-boxes/internal/cache"
	"github.com/amonjerro/dots-and-boxes/internal/game"
)

type MessageHandler interface {
	Handle(m *Message) error
}

type GameStartHandler struct{}

func (handler GameStartHandler) Handle(msg *Message) error {
	m, ok := msg.Body.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to parse message")
	}

	game, err := game.NewGame(m["width"].(int), m["height"].(int), m["player_count"].(int))
	if err != nil {
		return err
	}

	cacheClient := &cache.RedisClient{}
	ctx := context.Background()
	cacheClient.Set(ctx, game.Identifier.String(), game, time.Duration(24*time.Hour))

	return nil
}
