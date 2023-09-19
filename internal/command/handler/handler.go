package handler

import (
	"fmt"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command"
)

// Handler is an interface that defines a Handle method
type Handler interface {
	Handle(args []string) (string, error)
}

// RootHandler is a struct that contains a map of Handlers
type RootHandler struct {
	Handlers map[command.Command]Handler
}

// NewRootHandler returns a new RootHandler
func NewRootHandler() *RootHandler {
	return &RootHandler{
		Handlers: map[command.Command]Handler{
			command.JOIN:      &JoinHandler{},
			command.LEAVE:     &LeaveHandler{},
			command.FAIL:      &FailHandler{},
			command.SUSPICION: &SuspicionHandler{},
			command.DROPRATE:  &DropRateHandler{},
		},
	}
}

// Handle takes a command and returns the result of the command
func (h *RootHandler) Handle(args []string) (string, error) {
	cmd := command.Command(args[0])
	handler, ok := h.Handlers[cmd]
	if !ok {
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
	return handler.Handle(args)
}