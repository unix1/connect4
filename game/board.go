package game

import (
	"errors"
)

type board [][]slot

// newBoard creates a new board for the game
func newBoard(cols, rows int) (board, error) {
	if rows <= 0 || cols <= 0 {
		return nil, errors.New("rows or columns cannot be less than one")
	}
	board := make([][]slot, cols)
	for i := 0; i < cols; i++ {
		col := make([]slot, rows)
		board[i] = col
	}
	return board, nil
}

// move makes a turn for the player on the board; error is returned if the move
// is invalid - e.g. adding to an already full column
func (b board) move(m Move) (int, error) {
	var maxLength int
	row, err := b.findFirstOpenSlotInColumn(m.Col)
	if err != nil {
		return maxLength, err
	}
	b[m.Col][row].occupied = &m.Player
	maxLength = b.updateConnections(m.Col, row)
	return maxLength, nil
}

func (b board) findFirstOpenSlotInColumn(col int) (int, error) {
	var i int
	for i, slot := range b[col] {
		if slot.occupied == nil {
			return i, nil
		}
	}
	return i, errors.New("invalid move, column full")
}

// updateConnections updates connections for a slot in given coordinates and
// returns the max length of the affected connections.
func (b board) updateConnections(col, row int) int {
	var max int
	pairs := b.surroundingDirectionalPairs(col, row)
	for _, pair := range pairs {
		if pair.slot1 != nil {
			pair.slot1.connectSlot(pair.dir, &b[col][row])
		}
		if pair.slot2 != nil {
			pair.slot2.connectSlot(pair.dir, &b[col][row])
		}
		newLength := b[col][row].getConnectionLength(pair.dir)
		if newLength > max {
			max = newLength
		}
	}
	return max
}

type dirSlot struct {
	dir   direction
	slot1 *slot
	slot2 *slot
}

func (b board) surroundingDirectionalPairs(col, row int) []dirSlot {
	var pairs []dirSlot
	player := b[col][row].occupied
	// horizontal
	{
		var slot1, slot2 *slot
		if !b.isAtLeftEdge(col, row) && b[col-1][row].isOccupiedBy(player) {
			slot1 = &b[col-1][row]
		}
		if !b.isAtRightEdge(col, row) && b[col+1][row].isOccupiedBy(player) {
			slot2 = &b[col+1][row]
		}
		if slot1 != nil || slot2 != nil {
			pairs = append(pairs, dirSlot{dir: horizontal, slot1: slot1, slot2: slot2})
		}
	}
	// vertical
	{
		var slot1, slot2 *slot
		if !b.isAtBottomEdge(col, row) && b[col][row-1].isOccupiedBy(player) {
			slot1 = &b[col][row-1]
		}
		if !b.isAtTopEdge(col, row) && b[col][row+1].isOccupiedBy(player) {
			slot2 = &b[col][row+1]
		}
		if slot1 != nil || slot2 != nil {
			pairs = append(pairs, dirSlot{dir: vertical, slot1: slot1, slot2: slot2})
		}
	}
	// diagonal left up to right down \
	{
		var slot1, slot2 *slot
		if !b.isAtLeftEdge(col, row) && !b.isAtTopEdge(col, row) && b[col-1][row+1].isOccupiedBy(player) {
			slot1 = &b[col-1][row+1]
		}
		if !b.isAtRightEdge(col, row) && !b.isAtBottomEdge(col, row) && b[col+1][row-1].isOccupiedBy(player) {
			slot2 = &b[col+1][row-1]
		}
		if slot1 != nil || slot2 != nil {
			pairs = append(pairs, dirSlot{dir: diaglurd, slot1: slot1, slot2: slot2})
		}
	}
	// diagonal left down to right tup /
	{
		var slot1, slot2 *slot
		if !b.isAtLeftEdge(col, row) && !b.isAtBottomEdge(col, row) && b[col-1][row-1].isOccupiedBy(player) {
			slot1 = &b[col-1][row-1]
		}
		if !b.isAtRightEdge(col, row) && !b.isAtTopEdge(col, row) && b[col+1][row+1].isOccupiedBy(player) {
			slot2 = &b[col+1][row+1]
		}
		if slot1 != nil || slot2 != nil {
			pairs = append(pairs, dirSlot{dir: diagldru, slot1: slot1, slot2: slot2})
		}
	}
	return pairs
}

func (b board) isAtLeftEdge(col, row int) bool {
	var is bool
	if col == 0 {
		is = true
	}
	return is
}

func (b board) isAtRightEdge(col, row int) bool {
	var is bool
	if col == len(b)-1 {
		is = true
	}
	return is
}

func (b board) isAtTopEdge(col, row int) bool {
	var is bool
	if row == len(b[col])-1 {
		is = true
	}
	return is
}

func (b board) isAtBottomEdge(col, row int) bool {
	var is bool
	if row == 0 {
		is = true
	}
	return is
}

func (b board) isFull() bool {
	for _, col := range b {
		if col[len(col)-1].occupied == nil {
			return false
		}
	}
	return true
}
