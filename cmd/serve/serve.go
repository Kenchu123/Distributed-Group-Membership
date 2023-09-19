package serve

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command/server"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/logger"
)

var configPath string
var logPath string
var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server`,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	logger.Init(logPath)
	server, err := server.New(configPath, port)
	if err != nil {
		logrus.Fatal(err)
	}
	server.Run()
}

func New() *cobra.Command {
	return serveCmd
}

func init() {
	serveCmd.Flags().StringVarP(&configPath, "config", "c", ".msl/config.yml", "path to config file")
	serveCmd.Flags().StringVarP(&logPath, "log", "l", "logs/msl.log", "path to log file")
	serveCmd.Flags().StringVarP(&port, "port", "p", "7132", "port to listen on")
}
