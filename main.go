package main

import (
	"github.com/guths/caju-transaction-approver/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.CmdServe)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
