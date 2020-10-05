package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	game, err := NewGame()
	require.NoError(t, err)
	assert.Equal(t, defaultConfig, game.config)
}

func TestGameMoveInvalid(t *testing.T) {
	game, _ := NewGame()
	_, err := game.Move(Move{Player: 0, Col: -1})
	assert.Error(t, err)
	_, err = game.Move(Move{Player: -1, Col: 0})
	assert.Error(t, err)
}

// Short game, player 2 (the "o") should win in connect3
// |  o|
// | ox|
// |oxx|
// +---+
func TestGameWin(t *testing.T) {
	conf := Config{
		cols:    3,
		rows:    3,
		connect: 3,
		players: 2,
	}
	g, err := NewCustomGame(conf)
	require.NoError(t, err)
	res, err := g.Move(Move{Player: 0, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, StatusNextMove, res.Status)
	res, err = g.Move(Move{Player: 1, Col: 0})
	require.NoError(t, err)
	assert.Equal(t, StatusNextMove, res.Status)
	res, err = g.Move(Move{Player: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, StatusNextMove, res.Status)
	res, err = g.Move(Move{Player: 1, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, StatusNextMove, res.Status)
	res, err = g.Move(Move{Player: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, StatusNextMove, res.Status)
	res, err = g.Move(Move{Player: 1, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, StatusWin, res.Status)
	assert.Equal(t, 1, int(*res.Player))
}
