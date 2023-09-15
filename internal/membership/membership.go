package membership

import (
	"bytes"
	"encoding/gob"
)

// MemberList is a struct that contains a map of Members
type Membership struct {
	Members map[string]Member
}

// NewMembership creates a new membership
func NewMembership() *Membership {
	return &Membership{Members: make(map[string]Member)}
}

// Update updates the membership list
func (m *Membership) Update(ms *Membership) {
	// TODO: Iterate through the members in the membership list
}

// Serialize serializes the membership list
func Serialize(m *Membership) ([]byte, error) {
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
	m := NewMembership()
	buf := bytes.Buffer{}
	buf.Write(b)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
