package serve

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command/server"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/logger"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server`,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	logger.Init("msl.log")
	server, err := server.New(port)
	if err != nil {
		logrus.Fatal(err)
	}
	server.Run()
}

func New() *cobra.Command {
	return serveCmd
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "7132", "port to listen on")
}
