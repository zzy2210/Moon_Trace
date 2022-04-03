package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "moon_trance",
	Short: "moon_trace is a tool for safe",
}

func Excute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
