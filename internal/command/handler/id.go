package handler

import (
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/heartbeat"
)

type IDHandler struct{}

func (h *IDHandler) Handle(args []string) (string, error) {
	instance, err := heartbeat.GetInstance()
	if err != nil {
		return "", err
	}
	if !instance.IsRunning {
		return "Not in the membership list", nil
	}
	id := instance.Membership.ID
	return id, nil
}
