package handler

import (
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/heartbeat"
)

type JoinHandler struct{}

func (h *JoinHandler) Handle(args []string) (string, error) {
	heartbeat := heartbeat.GetInstance()
	if heartbeat.IsRunning == true {
		return "Already in the group", nil
	}
	heartbeat.Start()
	return "Start Heartbeating", nil
}
