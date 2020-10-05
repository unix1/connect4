package game

import (
	"errors"
	"fmt"
)

// Game is the main game containing the game config and board
type Game struct {
	config Config
	board  board
}

// Config is the game configuration
type Config struct {
	cols    int
	rows    int
	players int
	connect int
}

// Player is a 0-indexed integer to identify a player in the game
type Player int

func (p Player) String() string {
	return fmt.Sprintf("%d", int(p)+1)
}

var (
	defaultConfig = Config{
		cols:    7,
		rows:    6,
		players: 2,
		connect: 4,
	}
)

// NewGame creates a new game with default configuration
func NewGame() (*Game, error) {
	return NewCustomGame(defaultConfig)
}

// NewCustomGame creates a new game with custom configuration
func NewCustomGame(conf Config) (*Game, error) {
	board, err := newBoard(conf.rows, conf.cols)
	if err != nil {
		return nil, err
	}
	game := &Game{
		config: conf,
		board:  board,
	}
	return game, nil
}

// Move is a turn that each player makes
type Move struct {
	Player Player
	Col    int
}

// MoveResult is the result type returned after a player makes a move
type MoveResult struct {
	Status Status
	Player *Player
}

// Status is an enum for the expectation of the next event
type Status int

const (
	StatusNextMove Status = iota
	StatusDraw
	StatusWin
)

// Move makes a turn for a specific player
func (g *Game) Move(m Move) (MoveResult, error) {
	if err := g.validateMove(m); err != nil {
		return MoveResult{Status: StatusNextMove}, err
	}
	maxLength, err := g.board.move(m)
	if err != nil {
		return MoveResult{Status: StatusNextMove}, err
	}
	if maxLength >= g.config.connect {
		return MoveResult{Status: StatusWin, Player: &m.Player}, nil
	}
	if g.board.isFull() {
		return MoveResult{Status: StatusDraw}, nil
	}
	return MoveResult{Status: StatusNextMove}, nil
}

func (g *Game) validateMove(m Move) error {
	if err := g.validatePlayer(m.Player); err != nil {
		return err
	}
	if err := g.validateCol(m.Col); err != nil {
		return err
	}
	return nil
}

func (g *Game) validatePlayer(p Player) error {
	var err error
	if int(p) >= g.config.players || p < 0 {
		err = errors.New("invalid player")
	}
	return err
}

func (g *Game) validateCol(col int) error {
	var err error
	if col >= g.config.cols || col < 0 {
		err = errors.New("invalid column")
	}
	return err
}
