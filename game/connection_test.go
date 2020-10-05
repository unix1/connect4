package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	c := &connection{}
	s1 := &slot{}
	s2 := &slot{}
	c.addSlot(s1)
	c.addSlot(s2)
	assert.Equal(t, 2, c.length())
}
