package membership

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// MemberList is a struct that contains a map of Members
type Membership struct {
	ID      string             // ID of the node
	Members map[string]*Member // map of members
	mu      sync.Mutex         // mutex
}

// New creates a new membership
func New() (*Membership, error) {
	member, err := NewMemberSelf()
	if err != nil {
		return nil, fmt.Errorf("failed to create a new membership: %w", err)
	}
	return &Membership{ID: member.ID, Members: map[string]*Member{
		member.ID: member,
	}}, nil
}

// NewEmpty creates a new empty membership
func NewEmpty() *Membership {
	return &Membership{Members: map[string]*Member{}}
}

// IncreaseSelfHeartbeat increases the heartbeat of itself
func (m *Membership) IncreaseSelfHeartbeat() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Members[m.ID].IncreaseHeartbeat()
}

func (m *Membership) UpdateSelfState(state State) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Members[m.ID].UpdateState(m.Members[m.ID].Heartbeat, state)
}

// Update updates the membership list
func (m *Membership) Update(ms *Membership) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Iterate through the members in the membership list
	for _, member := range ms.Members {
		// Update the membership list with the member
		m.updateMember(member)
	}
}

// UpdateMember updates the membership list with a new member
func (m *Membership) updateMember(member *Member) {
	// case 1: member is not in the membership list
	if _, ok := m.Members[member.ID]; !ok {
		m.Members[member.ID] = member
		// TODO: prettier log
		logrus.Infof("JOINED: %v", member)
		return
	}
	// case 2: member is in the membership list
	// case 2.1: member is marked as failed, don't update the state
	if m.Members[member.ID].State == FAILED {
		return
	}
	// case 2.2: member is marked as left, don't update the state
	if m.Members[member.ID].State == LEFT {
		return
	}
	// case 2.3: member is marked as alive
	if m.Members[member.ID].State == ALIVE {
		// case 2.3.1: new member is marked as failed or left, update the state
		if member.State == FAILED || member.State == LEFT {
			m.Members[member.ID].UpdateState(member.Heartbeat, member.State)
		}
		// case 2.3.2: new member is marked as alive with higher heartbeat number and with equal or higher incarnation number, update the heartbeat
		if member.State == ALIVE && m.Members[member.ID].Heartbeat < member.Heartbeat && m.Members[member.ID].Incarnation <= member.Incarnation {
			m.Members[member.ID].UpdateHeartbeatAndIncarnation(member.Heartbeat, member.Incarnation)
		}
		// case 2.3.4: new member is marked as suspected with equal or higher incarnation number, update state, and incarnation number
		if member.State == SUSPECTED && m.Members[member.ID].Incarnation <= member.Incarnation {
			// case self alive but received suspected
			if m.ID == member.ID {
				m.Members[member.ID].UpdateStateAndIncarnation(ALIVE, member.Incarnation+1)
			} else {
				m.Members[member.ID].UpdateStateAndIncarnation(member.State, member.Incarnation)
			}
		}
		return
	}
	// case 2.4: member is marked as suspected
	if m.Members[member.ID].State == SUSPECTED {
		// case 2.4.1: new member is marked as failed or left, update the state
		if member.State == FAILED || member.State == LEFT {
			m.Members[member.ID].UpdateState(member.Heartbeat, member.State)
		}
		// case 2.4.2: new member is marked as alive with higher incarnation number, update the state, heartbeat, and incarnation number
		if member.State == ALIVE && m.Members[member.ID].Incarnation < member.Incarnation {
			m.Members[member.ID].UpdateStateHeartbeatAndIncarnation(member.State, member.Heartbeat, member.Incarnation)
		}
		return
	}
}

// DetectFailure detects failure
func (m *Membership) DetectFailure(failureDetectionTimeout time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, member := range m.Members {
		if member.State == ALIVE && time.Now().UnixMilli() > member.LastUpdateTime+failureDetectionTimeout.Milliseconds() {
			// TODO: suspected
			member.UpdateState(member.Heartbeat, FAILED)
		}
	}
}

// CleanUp cleans up the membership list
func (m *Membership) CleanUp(cleanupTimeout time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, member := range m.Members {
		if (member.State == FAILED || member.State == LEFT) && time.Now().UnixMilli() > member.LastUpdateTime+cleanupTimeout.Milliseconds() {
			delete(m.Members, id)
		}
	}
}

// Serialize serializes the membership list
func Serialize(m *Membership) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Deserialize deserializes the membership list
func Deserialize(b []byte) (*Membership, error) {
	m := NewEmpty()
	buf := bytes.Buffer{}
	buf.Write(b)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// String
func (m *Membership) String() string {
	return fmt.Sprintf("SelfID: %s\nMembership: %s\n", m.ID, m.Members)
}

// Get name of member
func (m *Membership) GetName() string {
	return strings.Split(m.ID, "_")[0]
}

// Get heartbeat target members' hostnames
func (m *Membership) GetHeartbeatTargetMembers() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	// TODO: introducers and algorithms for selecting heartbeat target members
	hostnames := []string{"fa23-cs425-8701.cs.illinois.edu", "fa23-cs425-8702.cs.illinois.edu"}
	for i, hostname := range hostnames {
		if hostname == m.GetName() {
			hostnames = append(hostnames[:i], hostnames[i+1:]...)
			break
		}
	}
	return hostnames
}

// Get snapshot of membership list
func (m *Membership) GetSnapshot() *Membership {
	m.mu.Lock()
	defer m.mu.Unlock()
	snapshot := NewEmpty()
	snapshot.ID = m.ID
	for _, member := range m.Members {
		snapshot.Members[member.ID] = member.GetSnapshot()
	}
	return snapshot
}
