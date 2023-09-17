package handler

import (
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/heartbeat"
)

type FailHandler struct{}

func (h *FailHandler) Handle(args []string) (string, error) {
	heartbeat := heartbeat.GetInstance()
	heartbeat.Stop()
	return "Failing", nil
}
