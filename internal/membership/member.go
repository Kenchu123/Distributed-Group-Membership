package membership

// Member is a struct that contains the ID, heartbeat, last update time, and state of a member
type Member struct {
	ID             string
	Heartbeat      int
	LastUpdateTime int
	State          State
}

// NewMember creates a new member
func NewMember() Member {
	return Member{}
}

// Update updates the heartbeat and last update time of a member
func (m *Member) Update(heartbeat int, lastUpdateTime int, state State) {
	// If the heartbeat is less than the current heartbeat, return
	if m.Heartbeat > heartbeat {
		return
	}
	m.Heartbeat = heartbeat
	m.LastUpdateTime = lastUpdateTime
	m.State = state
	// TODO: handle state changes
}
