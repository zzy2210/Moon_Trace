package cmd

import (
	"Moon_Trace/web"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start moon trace web api",
	Run: func(cmd *cobra.Command, args []string) {
		if err := web.Start(); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	// webCmd.Flags().StringVar()
}
