package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/config"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/disable"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/enable"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/fail"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/join"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/leave"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/list_mem"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/list_self"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/cmd/serve"
)

var rootCmd = &cobra.Command{
	Use:   "msl",
	Short: "membershiplist",
	Long:  `Machine Programming 2 - Distributed Group Membership`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(join.New(), leave.New(), fail.New(), serve.New(), config.New(), list_mem.New(), list_self.New(), enable.New(), disable.New())
}
