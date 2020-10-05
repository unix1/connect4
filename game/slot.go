package game

type slot struct {
	occupied  *Player
	connected connections
}

func (s *slot) connectSlot(d direction, sl *slot) {
	if s.getConnection(d) != nil && sl.getConnection(d) != nil {
		// both slots have existing connections in this direction, merge;
		// note that this will affect this slot's connection and all other
		// slots that hold this connection
		for _, slot := range sl.getConnection(d).slots {
			s.getConnection(d).addSlot(slot)
		}
		sl.setConnection(d, s.getConnection(d))
	} else {
		// one or both slots don't have a current connection
		if s.getConnection(d) == nil && sl.getConnection(d) == nil {
			conn := &connection{}
			conn.addSlot(s)
			conn.addSlot(sl)
			s.setNewConnection(d, &conn)
			sl.setNewConnection(d, &conn)
		} else if s.getConnection(d) == nil {
			conn := sl.getConnection(d)
			conn.addSlot(s)
			s.setConnection(d, conn)
		} else if sl.getConnection(d) == nil {
			conn := s.getConnection(d)
			conn.addSlot(sl)
			sl.setConnection(d, conn)
		}
	}
}

func (s *slot) setNewConnection(d direction, c **connection) {
	if s.connected == nil {
		s.connected = make(map[direction]**connection)
	}
	s.connected[d] = c
}

func (s *slot) setConnection(d direction, c *connection) {
	if s.connected == nil {
		s.setNewConnection(d, &c)
	}
	if _, ok := s.connected[d]; !ok {
		s.setNewConnection(d, &c)
	}
	*s.connected[d] = c
}

func (s *slot) getConnection(d direction) *connection {
	if s.connected == nil {
		return nil
	}
	if _, ok := s.connected[d]; !ok {
		return nil
	}
	if s.connected[d] == nil {
		return nil
	}
	return *s.connected[d]
}

func (s *slot) getConnectionLength(d direction) int {
	var length int
	if conn := s.getConnection(d); conn != nil {
		length = conn.length()
	}
	return length
}

func (s *slot) isOccupiedBy(p *Player) bool {
	if s.occupied == nil {
		return false
	}
	if p == nil {
		return false
	}
	if *p == *s.occupied {
		return true
	}
	return false
}
