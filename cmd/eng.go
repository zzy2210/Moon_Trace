package cmd

import (
	"Moon_Trace/eng"
	"github.com/spf13/cobra"
)

var engCmd = &cobra.Command{
	Short: "start eng",
	Run: func(cmd *cobra.Command, args []string) {
		eng.Execute()
	},
}
