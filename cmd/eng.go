package cmd

import (
	"Moon_Trace/eng"
	"github.com/spf13/cobra"
)

var engArgs eng.Args
var engCmd = &cobra.Command{
	Short: "start eng",
	Run: func(cmd *cobra.Command, args []string) {
		eng.Execute(&engArgs)
	},
}

func init() {
	rootCmd.AddCommand(engCmd)
	engCmd.Flags().StringVar(&engArgs.ConfigPath, "config_path", "/conf/engConf.conf", "引擎配置文件")
	engCmd.Flags().StringVar(&engArgs.Addr, "addr", ":60001", "grpc监听地址")
}
