package handler

import "github.com/sirupsen/logrus"

type FailHandler struct{}

func (h *FailHandler) Handle(args []string) (string, error) {
	logrus.Info("Self Failing")
	return "Failing", nil
}
