package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Init initializes the logger to log to the specified file and terminal.
func Init(logPath string) {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(io.MultiWriter(f, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
}
