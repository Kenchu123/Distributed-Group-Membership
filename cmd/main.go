package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/logger"
)

func main() {
	logger.Init("membership.log")
	logrus.Info("Hello, World!")
}
