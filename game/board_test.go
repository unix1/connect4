package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	board, err := newBoard(3, 5)
	require.NoError(t, err)
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, slot{}, board[i][j])
		}
	}
}

func TestNewBoardError(t *testing.T) {
	_, err := newBoard(3, 0)
	assert.Error(t, err)
	_, err = newBoard(-4, 5)
	assert.Error(t, err)
}

func TestBoardMoveVertical(t *testing.T) {
	board, _ := newBoard(3, 4)
	length, err := board.move(Move{Player: 0, Col: 0})
	require.NoError(t, err)
	// length will be 0 because no connection exists yet
	assert.Equal(t, 0, length)
	for i := 0; i < 3; i++ {
		length, err = board.move(Move{Player: 0, Col: 0})
		require.NoError(t, err)
		assert.Equal(t, i+2, length)
	}
	_, err = board.move(Move{Player: 0, Col: 0})
	assert.Error(t, err)
}

func TestBoardMoveHorizontal(t *testing.T) {
	board, _ := newBoard(4, 3)
	length, err := board.move(Move{Player: 0, Col: 0})
	require.NoError(t, err)
	// length will be 0 because no connection exists yet
	assert.Equal(t, 0, length)
	for i := 0; i < 3; i++ {
		length, err = board.move(Move{Player: 0, Col: i + 1})
		require.NoError(t, err)
		assert.Equal(t, i+2, length)
	}
}

// Tests the following diagonal scenario
// |  o |
// | ox |
// |oxx |
// +----+
func TestBoardMoveDiagonal(t *testing.T) {
	board, _ := newBoard(4, 3)
	length, err := board.move(Move{Player: 0, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, 0, length)
	length, err = board.move(Move{Player: 1, Col: 0})
	require.NoError(t, err)
	assert.Equal(t, 0, length)
	length, err = board.move(Move{Player: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, 2, length)
	length, err = board.move(Move{Player: 1, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, 2, length)
	length, err = board.move(Move{Player: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, 2, length)
	length, err = board.move(Move{Player: 1, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, 3, length)
}

// Tests the following scenario merging existing connections:
// the big O is the last move that should result in length of 5.
// |     |
// |     |
// |ooOoo|
// +-----+
func TestBoardMoveMerge(t *testing.T) {
	board, _ := newBoard(5, 3)
	length, err := board.move(Move{Player: 0, Col: 0})
	require.NoError(t, err)
	assert.Equal(t, 0, length)
	length, err = board.move(Move{Player: 0, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, 2, length)
	length, err = board.move(Move{Player: 0, Col: 3})
	require.NoError(t, err)
	assert.Equal(t, 0, length)
	length, err = board.move(Move{Player: 0, Col: 4})
	require.NoError(t, err)
	assert.Equal(t, 2, length)
	length, err = board.move(Move{Player: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, 5, length)
}

func TestBoardIsFull(t *testing.T) {
	board, _ := newBoard(2, 3)
	for i := 0; i < 3; i++ {
		assert.False(t, board.isFull())
		_, err := board.move(Move{Player: 0, Col: 0})
		require.NoError(t, err)
		_, err = board.move(Move{Player: 1, Col: 1})
		require.NoError(t, err)
	}
	assert.True(t, board.isFull())
}
