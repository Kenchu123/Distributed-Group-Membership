package membership

// State is a string that represents the state of a member
type State string

// Constants for the different states of a member
const (
	Alive  State = "ALIVE"
	Failed State = "FAILED"
	Left   State = "LEFT"
)
