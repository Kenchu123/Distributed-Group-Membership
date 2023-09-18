package config

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/handler"
)

var enable bool
var suspicionCmd = &cobra.Command{
	Use:   "suspicion",
	Short: "Set suspicion",
	Long:  `Set suspicion`,
	Run:   suspicion,
}

func suspicion(cmd *cobra.Command, args []string) {
	client, err := client.New(configPath, machineRegex)
	if err != nil {
		logrus.Fatalf("failed to create command client: %v", err)
	}
	results := client.Run([]string{string(handler.SUSPICION), strconv.FormatBool(enable)})
	for _, r := range results {
		if r.Err != nil {
			logrus.Errorf("failed to send command to %s: %v\n", r.Hostname, r.Err)
			continue
		}
		logrus.Printf("%s: %s\n", r.Hostname, r.Message)
	}
}

func init() {
	suspicionCmd.Flags().BoolVarP(&enable, "enable", "e", false, "enable or disable suspicion")
}
