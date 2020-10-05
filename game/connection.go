package game

type direction int

const (
	horizontal direction = iota
	vertical
	diaglurd // diagonal left up -> right down: \
	diagldru // diagonal left down -> right up: /
)

type connections map[direction]**connection

type connection struct {
	slots []*slot
}

func (c *connection) addSlot(s *slot) {
	c.slots = append(c.slots, s)
}

func (c *connection) length() int {
	return len(c.slots)
}
