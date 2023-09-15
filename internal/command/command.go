package command

// State is a string that represents the state of a member
type Command string

// Constants for the different states of a member
const (
	JOIN  Command = "JOIN"
	LEAVE Command = "LEAVE"
	FAIL  Command = "FAIL"
)
