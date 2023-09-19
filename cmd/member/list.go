package member

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/command/client"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List membershiplist",
	Long:  `List membershiplist in json format`,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	client, err := client.New(configPath, machineRegex)
	if err != nil {
		logrus.Fatalf("failed to create command client: %v", err)
	}
	results := client.Run([]string{string(command.LIST)})
	for _, r := range results {
		if r.Err != nil {
			logrus.Errorf("failed to send command to %s: %v\n", r.Hostname, r.Err)
			continue
		}
		logrus.Printf("%s: %s\n", r.Hostname, r.Message)
	}
}
