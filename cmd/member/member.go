package member

import (
	"github.com/spf13/cobra"
)

var configPath string
var machineRegex string
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Manage group membership",
	Long:  `Manage group membership`,
}

func New() *cobra.Command {
	return memberCmd
}

func init() {
	memberCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".msl/config.yml", "path to config file")
	memberCmd.PersistentFlags().StringVarP(&machineRegex, "machine-regex", "m", ".*", "regex for machines to join (e.g. \"0[1-9]\")")
	memberCmd.AddCommand(joinCmd, leaveCmd, failCmd, listCmd)
}
