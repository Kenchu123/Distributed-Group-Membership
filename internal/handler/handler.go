package handler

import "fmt"

// Handler is an interface that defines a Handle method
type Handler interface {
	Handle(args []string) (string, error)
}

// RootHandler is a struct that contains a map of Handlers
type RootHandler struct {
	Handlers map[Command]Handler
}

// NewRootHandler returns a new RootHandler
func NewRootHandler() *RootHandler {
	return &RootHandler{
		Handlers: map[Command]Handler{
			JOIN:      &JoinHandler{},
			LEAVE:     &LeaveHandler{},
			FAIL:      &FailHandler{},
			SUSPICION: &SuspicionHandler{},
			DROPRATE:  &DropRateHandler{},
		},
	}
}

// Handle takes a command and returns the result of the command
func (h *RootHandler) Handle(args []string) (string, error) {
	cmd := Command(args[0])
	handler, ok := h.Handlers[cmd]
	if !ok {
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
	return handler.Handle(args)
}
