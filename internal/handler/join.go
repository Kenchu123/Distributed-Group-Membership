package handler

import "github.com/sirupsen/logrus"

type JoinHandler struct{}

func (h *JoinHandler) Handle(args []string) (string, error) {
	logrus.Info("Self joining the group")
	return "Joining the group", nil
}
