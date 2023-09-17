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
	timestamp := time.Now().Unix()
	return &Member{
		ID:             fmt.Sprintf("%s-%d", name, timestamp),
		Heartbeat:      0,
		LastUpdateTime: time.Now().Unix(),
		State:          ALIVE,
		Incarnation:    0,
	}, nil
}

func (m *Member) UpdateState(heartbeat int, state State) {
	m.Heartbeat = heartbeat
	m.State = state
	m.LastUpdateTime = time.Now().Unix()
	// TODO: prettier log
	logrus.Infof("Update State: %v", m)
}

func (m *Member) UpdateHeartbeatAndIncarnation(heartbeat, incarnation int) {
	m.Heartbeat = heartbeat
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().Unix()
	// TODO: prettier log
	logrus.Infof("Update Heartbeat and Incarnation: %v", m)
}

func (m *Member) UpdateStateAndIncarnation(state State, incarnation int) {
	m.State = state
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().Unix()
	// TODO: prettier log
	logrus.Infof("Update State and Incarnation: %v", m)
}

func (m *Member) UpdateStateHeartbeatAndIncarnation(state State, heartbeat, incarnation int) {
	m.State = state
	m.Heartbeat = heartbeat
	m.Incarnation = incarnation
	m.LastUpdateTime = time.Now().Unix()
	// TODO: prettier log
	logrus.Infof("Update State, Heartbeat and Incarnation: %v", m)
}

func (m *Member) IncreaseHeartbeat() {
	m.Heartbeat++
	m.LastUpdateTime = time.Now().Unix()
	// TODO: prettier log
	logrus.Infof("Increase Heartbeat: %v", m)
}

// String
func (m *Member) String() string {
	return fmt.Sprintf("Member{ID=%s, Heartbeat=%d, LastUpdateTime=%d, State=%s, Incarnation=%d}", m.ID, m.Heartbeat, m.LastUpdateTime, m.State, m.Incarnation)
}
