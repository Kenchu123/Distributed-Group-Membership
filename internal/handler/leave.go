package handler

import (
	"time"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/heartbeat"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/membership"
)

type LeaveHandler struct{}

func (h *LeaveHandler) Handle(args []string) (string, error) {
	instance := heartbeat.GetInstance()
	if instance.IsRunning == false {
		return "Not in the group", nil
	}
	// change the state of the node to leave
	instance.Membership.UpdateSelfState(membership.LEFT)
	// TODO: fine tuning the time sleep here
	time.Sleep(heartbeat.HEARTBEAT_INTERVAL * 3)
	instance.Stop()
	return "Leaving the group", nil
}
