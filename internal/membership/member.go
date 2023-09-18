package membership

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Member is a struct
type Member struct {
	ID             string
	Heartbeat      int
	LastUpdateTime int64
	State          State
	Incarnation    int
}

// NewMemberSelf creates a new member of itself information
func NewMemberSelf() (*Member, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}
	return &Member{
		ID:             fmt.Sprintf("%s_%d", name, time.Now().Unix()),
		Heartbeat:      0,
		LastUpdateTime: time.Now().UnixMilli(),
		State:          ALIVE,
		Incarnation:    0,
	}, nil
}

func (m *Member) UpdateState(heartbeat int, state State) {
	m.Heartbeat = heartbeat
	m.State = state
	m.LastUpdateTime = time.Now().UnixMilli()
	logrus.Infof("[STATE CHANGE] %s change state to %s", m.ID, m.State)
}

func (m *Member) UpdateHeartbeatAndIncarnation(heartbeat, incarnation int) {
	m.Heartbeat = heartbeat
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().UnixMilli()
	logrus.Infof("[HEARTBEAT, INCARNATION] %s change heartbeat to %d, incarnation to %d", m.ID, m.Heartbeat, m.Incarnation)
}

func (m *Member) UpdateStateAndIncarnation(state State, incarnation int) {
	m.State = state
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().UnixMilli()
	logrus.Infof("[STATE, INCARNATION] %s change state to %s, incarnation to %d", m.ID, m.State, m.Incarnation)
}

func (m *Member) UpdateStateHeartbeatAndIncarnation(state State, heartbeat, incarnation int) {
	m.State = state
	m.Heartbeat = heartbeat
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().UnixMilli()
	logrus.Infof("[STATE, HEARTBEAT, INCARNATION] %s change state to %s, heartbeat to %d, incarnation to %d", m.ID, m.State, m.Heartbeat, m.Incarnation)
}

func (m *Member) IncreaseHeartbeat() {
	m.Heartbeat++
	m.LastUpdateTime = time.Now().UnixMilli()
	logrus.Infof("[HEARTBEAT] %s increase heartbeat to %d", m.ID, m.Heartbeat)
}

func (m *Member) GetSnapshot() *Member {
	return &Member{
		ID:             m.ID,
		Heartbeat:      m.Heartbeat,
		LastUpdateTime: m.LastUpdateTime,
		State:          m.State,
		Incarnation:    m.Incarnation,
	}
}

func (m *Member) String() string {
	return fmt.Sprintf("\n{\n\tID: %s\n\tHeartbeat: %d\n\tLastUpdateTime: %d\n\tState: %s\n\tIncarnation: %d\n}", m.ID, m.Heartbeat, m.LastUpdateTime, m.State, m.Incarnation)
}
