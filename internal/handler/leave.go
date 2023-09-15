package handler

import "github.com/sirupsen/logrus"

type LeaveHandler struct{}

func (h *LeaveHandler) Handle(args []string) (string, error) {
	logrus.Info("Self Leaving the group")
	return "Leaving the group", nil
}
